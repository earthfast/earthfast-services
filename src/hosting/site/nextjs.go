package site

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"armada-node/logging"
	"armada-node/model"

	"go.uber.org/zap"
)

// ProcessRegistry tracks all Next.js processes to avoid duplication
var (
	processRegistry     = make(map[string]*ProcessInfo)
	processRegistryLock sync.RWMutex
)

// ProcessInfo tracks information about a running Next.js process
type ProcessInfo struct {
	ProjectID model.ID
	Process   *exec.Cmd
	Port      int
	StartLock sync.Mutex // Controls process lifecycle operations
}

// NextJSVersion represents a Next.js server instance
type NextJSVersion struct {
	logger     *zap.Logger
	project    *model.Project
	projectDir string
	metadata   *model.ProjectMetadata
	port       int
}

// NewNextJSVersion creates a new Next.js server version
func NewNextJSVersion(logger *zap.Logger, dataDir string, p *model.Project) (Version, error) {
	metadata, err := p.ParseMetadata()
	if err != nil {
		return nil, fmt.Errorf("parsing metadata: %w", err)
	}

	// Default port if not specified
	if metadata.Port == 0 {
		metadata.Port = 3000
	}

	// Create project directory
	projectDir := filepath.Join(dataDir, p.ID.Hex())
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return nil, fmt.Errorf("creating project directory: %w", err)
	}

	nextJS := &NextJSVersion{
		logger:     logger,
		project:    p,
		projectDir: projectDir,
		metadata:   metadata,
		port:       metadata.Port,
	}

	// Start the Next.js server asynchronously
	go nextJS.Start()

	return nextJS, nil
}

// PrepareFromBundle downloads and extracts the Next.js bundle
func (v *NextJSVersion) PrepareFromBundle() error {
	// Clear logs when preparing a new bundle
	logging.ClearLogs(logging.LogTypeNextJS)

	// Use bundleURL from metadata if it exists, otherwise fall back to project.Content
	bundleURL := v.metadata.BundleURL
	if bundleURL == "" {
		bundleURL = v.project.Content
	}

	if bundleURL == "" {
		return fmt.Errorf("no bundle URL specified for project")
	}

	v.logger.Info("Downloading Next.js bundle",
		zap.String("url", bundleURL),
		zap.String("projectId", v.project.ID.Hex()))

	// Clean existing directory
	if err := os.RemoveAll(v.projectDir); err != nil {
		return fmt.Errorf("failed to clean project directory: %w", err)
	}
	if err := os.MkdirAll(v.projectDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Download bundle
	resp, err := http.Get(bundleURL)
	if err != nil {
		return fmt.Errorf("failed to download bundle: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bundle download failed with status: %d", resp.StatusCode)
	}

	// Extract tar.gz
	v.logger.Info("Extracting bundle", zap.String("directory", v.projectDir))
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar: %w", err)
		}

		// Skip if not a file or directory
		if header.Typeflag != tar.TypeReg && header.Typeflag != tar.TypeDir {
			continue
		}

		target := filepath.Join(v.projectDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", target, err)
			}
		case tar.TypeReg:
			dir := filepath.Dir(target)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", target, err)
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return fmt.Errorf("failed to write file %s: %w", target, err)
			}
			f.Close()
		}
	}

	v.logger.Info("Bundle extraction complete")
	return nil
}

// Start launches the Next.js server process using the centralized registry
func (v *NextJSVersion) Start() {
	// Get or create process info using a single write lock
	processRegistryLock.Lock()
	procInfo, exists := processRegistry[v.project.ID.Hex()]
	if !exists {
		procInfo = &ProcessInfo{
			ProjectID: v.project.ID,
			Port:      v.port,
		}
		processRegistry[v.project.ID.Hex()] = procInfo
	}
	processRegistryLock.Unlock()

	// Try to acquire the start lock
	if !procInfo.StartLock.TryLock() {
		v.logger.Info("Process already being started by another goroutine",
			zap.String("projectId", v.project.ID.Hex()))
		return
	}
	defer procInfo.StartLock.Unlock()

	// Check if process is running
	if procInfo.Process != nil && procInfo.Process.Process != nil {
		// Check if process is still alive
		if err := procInfo.Process.Process.Signal(syscall.Signal(0)); err == nil {
			// Process is still alive
			v.logger.Info("Process already running",
				zap.String("projectId", v.project.ID.Hex()))
			return
		}
		// Process is dead but we still have a reference - clean it up
		procInfo.Process = nil
	}

	// Prepare bundle if needed
	if err := v.PrepareFromBundle(); err != nil {
		v.logger.Error("Failed to prepare bundle",
			zap.Error(err),
			zap.String("projectId", v.project.ID.Hex()))
		return
	}

	// Find server.js location
	serverDir, err := v.findServerJsPath()
	if err != nil {
		v.logger.Error("Failed to find server.js",
			zap.Error(err),
			zap.String("projectId", v.project.ID.Hex()))
		return
	}

	// Start Node.js server
	v.logger.Info("Starting Next.js server",
		zap.String("serverDir", serverDir),
		zap.Int("port", procInfo.Port))

	cmd := exec.Command("node", "server.js")
	cmd.Dir = serverDir
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("PORT=%d", procInfo.Port),
		"NODE_ENV=production",
		"HOST=0.0.0.0",
		"HOSTNAME=0.0.0.0",
	)

	// Capture output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		v.logger.Error("Failed to create stdout pipe", zap.Error(err))
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		v.logger.Error("Failed to create stderr pipe", zap.Error(err))
		return
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		v.logger.Error("Failed to start Next.js process", zap.Error(err))
		return
	}

	// Save process info
	procInfo.Process = cmd

	// Log stdout/stderr
	go v.logOutput(stdout, "stdout")
	go v.logOutput(stderr, "stderr")

	// Start a separate goroutine to monitor the process
	go v.monitorProcess(procInfo)

	// Give the process a moment to start
	time.Sleep(1 * time.Second)

	v.logger.Info("Next.js process started",
		zap.String("projectId", v.project.ID.Hex()),
		zap.Int("pid", cmd.Process.Pid))
}

// findServerJsPath locates the server.js file
func (v *NextJSVersion) findServerJsPath() (string, error) {
	// Try several likely locations
	possibleLocations := []string{
		filepath.Join(v.projectDir, "server.js"),
		filepath.Join(v.projectDir, ".next", "server.js"),
		filepath.Join(v.projectDir, "output", "server.js"),
	}

	// Also check first-level subdirectories
	files, err := os.ReadDir(v.projectDir)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				possibleLocations = append(possibleLocations,
					filepath.Join(v.projectDir, file.Name(), "server.js"))
			}
		}
	}

	// Log all checked locations for debugging
	v.logger.Debug("Searching for server.js",
		zap.Strings("locations", possibleLocations))

	// Check each location
	for _, path := range possibleLocations {
		if _, err := os.Stat(path); err == nil {
			v.logger.Info("Found server.js", zap.String("path", path))
			return filepath.Dir(path), nil
		}
	}

	return "", fmt.Errorf("server.js not found in any expected location")
}

// logOutput logs process output with project ID context
func (v *NextJSVersion) logOutput(r io.Reader, source string) {
	logging.NextJSLogProcessor(r, source, v.project.ID.Hex(), v.logger)
}

// monitorProcess watches the process and restarts it if it crashes
func (v *NextJSVersion) monitorProcess(procInfo *ProcessInfo) {
	// Wait for process to exit
	err := procInfo.Process.Wait()

	// Process has exited
	if err != nil {
		v.logger.Warn("Next.js process exited with error",
			zap.String("projectId", v.project.ID.Hex()),
			zap.Error(err))
	} else {
		v.logger.Info("Next.js process exited normally",
			zap.String("projectId", v.project.ID.Hex()))
	}

	// Clear the process reference
	procInfo.StartLock.Lock()
	procInfo.Process = nil
	procInfo.StartLock.Unlock()

	// Wait before restarting
	time.Sleep(5 * time.Second)

	// Start again
	v.Start()
}

// Stop terminates the Next.js server process
func (v *NextJSVersion) Stop() {
	processRegistryLock.RLock()
	procInfo, exists := processRegistry[v.project.ID.Hex()]
	processRegistryLock.RUnlock()

	if !exists || procInfo.Process == nil {
		return
	}

	// Lock to safely stop the process
	procInfo.StartLock.Lock()
	defer procInfo.StartLock.Unlock()

	if procInfo.Process != nil && procInfo.Process.Process != nil {
		v.logger.Info("Stopping Next.js server",
			zap.String("projectId", v.project.ID.Hex()))

		// First try SIGTERM for graceful shutdown
		if err := procInfo.Process.Process.Signal(syscall.SIGTERM); err != nil {
			v.logger.Warn("Failed to send SIGTERM to Next.js server, using SIGKILL",
				zap.Error(err))
			// If SIGTERM fails, fall back to Kill
			_ = procInfo.Process.Process.Kill()
		} else {
			// Give the process a moment to shut down gracefully (500ms)
			gracePeriodTimer := time.NewTimer(500 * time.Millisecond)
			processExited := make(chan struct{})

			// Start a goroutine to wait for the process
			go func() {
				procInfo.Process.Wait()
				close(processExited)
			}()

			// Wait for either the process to exit or the grace period to expire
			select {
			case <-processExited:
				v.logger.Info("Next.js server shut down gracefully")
				gracePeriodTimer.Stop()
			case <-gracePeriodTimer.C:
				v.logger.Warn("Grace period expired, force killing Next.js server")
				_ = procInfo.Process.Process.Kill()
			}
		}

		procInfo.Process = nil
	}
}

// Delete removes the Next.js project files
func (v *NextJSVersion) Delete() error {
	v.Stop()
	return os.RemoveAll(v.projectDir)
}

// ServeHTTP handles HTTP requests to the Next.js server
func (v *NextJSVersion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get process info
	processRegistryLock.RLock()
	procInfo, exists := processRegistry[v.project.ID.Hex()]
	processRegistryLock.RUnlock()

	// If no process exists or the process is nil, start it
	if !exists || procInfo.Process == nil {
		v.logger.Info("Process not running, requesting start",
			zap.String("projectId", v.project.ID.Hex()),
			zap.String("path", r.URL.Path))

		// Start in background
		go v.Start()

		// Return simple text response
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Retry-After", "5")
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "Service is starting. Please try again in a few moments.")
		return
	}

	// Get hostname for proxy
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "127.0.0.1"
	}

	// Create reverse proxy
	target := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", hostname, procInfo.Port),
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Set up custom director to preserve the path
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = r.URL.Path
		req.URL.RawPath = r.URL.RawPath
		req.URL.RawQuery = r.URL.RawQuery

		v.logger.Debug("Proxying to Next.js server",
			zap.String("path", req.URL.Path),
			zap.String("query", req.URL.RawQuery))
	}

	// Add error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		// Check if this is a client disconnection
		if err != nil && (err.Error() == "context canceled" || err.Error() == "net/http: request canceled") {
			// This is just a client disconnection, log as debug not error
			v.logger.Debug("Client disconnected during request",
				zap.String("path", r.URL.Path),
				zap.String("projectId", v.project.ID.Hex()))
			return // Don't send a response - the client is gone
		}

		// For actual errors, continue with existing logic
		v.logger.Error("Proxy error",
			zap.Error(err),
			zap.String("path", r.URL.Path),
			zap.String("projectId", v.project.ID.Hex()))

		// Check if process is still running
		procInfo.StartLock.Lock()
		processAlive := procInfo.Process != nil && procInfo.Process.Process != nil
		procInfo.StartLock.Unlock()

		if !processAlive {
			// Process might have crashed, try to restart it
			go v.Start()

			// Return simple text response
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Retry-After", "5")
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, "Service error. The server is restarting. Please try again in a few moments.")
			return
		}

		// Otherwise show a generic error
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error")
	}

	// Forward the request
	proxy.ServeHTTP(w, r)
}

// IsProject checks if this version is for the given project
func (v *NextJSVersion) IsProject(p *model.Project) bool {
	if p == nil || p.ID != v.project.ID {
		return false
	}

	// Get current bundle URL from metadata or content
	currentBundleURL := ""
	if metadata, err := p.ParseMetadata(); err == nil && metadata.BundleURL != "" {
		currentBundleURL = metadata.BundleURL
	} else {
		currentBundleURL = p.Content
	}

	// Get the bundle URL we're currently using
	previousBundleURL := v.metadata.BundleURL
	if previousBundleURL == "" {
		previousBundleURL = v.project.Content
	}

	// If bundle URL changed, this is no longer the same version
	if currentBundleURL != previousBundleURL {
		v.logger.Info("Project bundle URL changed",
			zap.String("projectId", v.project.ID.Hex()),
			zap.String("oldURL", previousBundleURL),
			zap.String("newURL", currentBundleURL))
		return false
	}

	return true
}

// String returns a string representation of this version
func (v *NextJSVersion) String() string {
	return fmt.Sprintf("NextJSSite[id=%s]", v.project.ID.Hex())
}

// StopAllNextJSProcesses stops all managed processes (useful for shutdown)
func StopAllNextJSProcesses() {
	processRegistryLock.RLock()
	defer processRegistryLock.RUnlock()

	for _, procInfo := range processRegistry {
		if procInfo.Process != nil && procInfo.Process.Process != nil {
			_ = procInfo.Process.Process.Kill()
		}
	}
}
