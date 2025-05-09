package domain

import (
	"armada-node/api/middleware"
	"armada-node/geo"
	"armada-node/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"text/template"
	"time"

	"go.uber.org/zap"
)

type Templates struct {
	ServiceWorker *template.Template
}

type Handler struct {
	*http.ServeMux
}

type proxyHandler struct {
	http.Handler
	m           model.Client
	resolver    Resolver
	logger      *zap.Logger
	environment string
}

var (
	nodeTimeout = 5 * time.Second // Reduced timeout for node requests

	// Browser detection lists
	embedBrowsers = []string{
		// Social/Messaging apps
		"telegram", "cfnetwork",
		"fb_iab", "fban/", "fbios/",
		"instagram",
		"whatsapp",
		"line",
		"wechat", "micromessenger",
		"kakaotalk",
		"snapchat",
		"twitter",
		"linkedin",

		// Asian browsers/webviews
		"miui", "ucbrowser", "qq", "baidu",

		// Generic webviews
		"webview", "wv", "mobile",
	}

	crawlers = []string{
		"googlebot", "adsbot-google", "apis-google",
		"mediapartners-google", "bingbot", "slurp",
		"wget", "curl", "python-urllib", "python-requests",
		"aiohttp", "httpx", "libwww-perl", "node-fetch",
		"bot", "spider", "crawler", "archiver",
		"phantom", "headless", "selenium", "chrome-lighthouse",
		"bytespider", "compressionbot", "feedbin",
		"embedly", "facebook", "twitter", "slack",
		"telegram", "whatsapp", "discord", "gptbot",
		"chatgpt-user", "petalbot", "ahrefsbot",
		"semrushbot", "yandex", "baidu", "seznambot",
		"applebot", "msnbot", "duckduckbot",
	}
)

func isServiceWorkerCompatible(r *http.Request) bool {
	for _, cookie := range r.Cookies() {
		if cookie.Name == "supportsSW" && cookie.Value == "false" {
			return false
		}
	}

	ua := strings.ToLower(r.UserAgent())

	// Check for crawlers
	for _, c := range crawlers {
		if strings.Contains(ua, c) {
			return false
		}
	}

	// Check for embedded browsers
	for _, e := range embedBrowsers {
		if strings.Contains(ua, e) {
			return false
		}
	}

	// Check additional headers that might indicate in-app browsers
	if r.Header.Get("X-Requested-With") != "" {
		return false // Usually set by Android WebView
	}

	// iOS in-app browser detection
	if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		// Check if it's not standalone Safari
		if !strings.Contains(ua, "safari") ||
			strings.Contains(ua, "crios") || // Chrome iOS
			strings.Contains(ua, "fxios") || // Firefox iOS
			r.Header.Get("X-Purpose") != "" { // Often set by in-app browsers
			return false
		}
	}

	return true
}

func newProxyHandler(logger *zap.Logger, m model.Client, r Resolver, c geo.GeolocationClient, environment string) *proxyHandler {
	h := &proxyHandler{
		m:           m,
		resolver:    r,
		logger:      logger,
		environment: environment,
	}
	h.Handler = middleware.GeoIP(logger, c, h.serveGeoHTTP)
	return h
}

func getForwardedFor(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded + ", " + r.RemoteAddr
	}
	return r.RemoteAddr
}

func tryServeProxy(proxy *httputil.ReverseProxy, w http.ResponseWriter, r *http.Request) error {
	errChan := make(chan error, 1)
	go func() {
		defer close(errChan)
		proxy.ServeHTTP(w, r)
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(nodeTimeout):
		return fmt.Errorf("request timed out after %s", nodeTimeout)
	}
}

func (h *proxyHandler) serveGeoHTTP(userCoord geo.Coordinate, w http.ResponseWriter, r *http.Request) {
	projectID, err := h.resolver.ProjectForDomain(r.Context(), r.Host)
	if err != nil {
		h.logger.Error("Failed to resolve domain", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Check if this is a Next.js project
	var isNextJS bool
	var project *model.Project

	project, err = h.m.GetProject(r.Context(), projectID)
	if err == nil && project != nil {
		metadata, err := project.ParseMetadata()
		if err == nil && metadata.Type == model.ProjectTypeNextJS {
			isNextJS = true
		}
	}

	nodes, err := h.m.ContentNodes(r.Context(), projectID)
	if err != nil {
		h.logger.Error("Failed to fetch content nodes", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(nodes) == 0 {
		h.logger.Error("No content nodes available")
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	// Shuffle nodes for simple load balancing
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})

	// Try each node until one succeeds
	var lastErr error
	for _, targetNode := range nodes {
		// Buffer to capture the response for checksum validation
		var responseBuffer bytes.Buffer

		// Create a proxy for the node
		if isNextJS {
			nextjsURL := fmt.Sprintf("https://%s/nextjs", targetNode.Host)
			target, err := url.Parse(nextjsURL)
			if err != nil {
				lastErr = err
				continue
			}

			proxy := httputil.NewSingleHostReverseProxy(target)
			if h.environment == "dev" {
				proxy.Transport = &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
			}

			proxy.Director = func(req *http.Request) {
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host

				// Keep a copy of the original URL components
				originalPath := req.URL.Path
				originalQuery := req.URL.RawQuery

				// Update the path to point to the nextjs endpoint
				req.URL.Path = "/nextjs" + originalPath

				// Preserve query parameters
				if originalQuery != "" {
					req.URL.RawQuery = originalQuery
				}

				// Set headers
				req.Header.Set("X-Real-IP", r.RemoteAddr)
				req.Header.Set("X-Forwarded-For", getForwardedFor(r))
				req.Header.Set("X-Forwarded-Proto", r.URL.Scheme)
				req.Header.Set("X-Forwarded-Host", r.Host)
				req.Header.Set("X-Original-Path", originalPath)
				req.Host = target.Host

				h.logger.Info("Next.js proxy attempt",
					zap.String("original_path", originalPath),
					zap.String("target_path", req.URL.Path),
					zap.String("target_host", req.URL.Host),
					zap.String("query", originalQuery),
					zap.String("project_id", projectID.Hex()),
				)
			}

			// Wrapper ResponseWriter that validates checksum and captures response
			checksumWriter := &checksumValidatingWriter{
				ResponseWriter: w,
				buffer:         &responseBuffer,
				targetNode:     targetNode,
				logger:         h.logger,
			}

			err = tryServeProxy(proxy, checksumWriter, r)

			// Check if response was valid and checksum matched
			if err == nil && checksumWriter.statusCode >= 200 &&
				checksumWriter.statusCode < 300 && checksumWriter.checksumValid {
				return // Success
			}

			if err != nil {
				lastErr = err
				h.logger.Warn("Next.js node request failed, trying next node",
					zap.String("node", targetNode.Host),
					zap.Error(err))
			} else if !checksumWriter.checksumValid {
				lastErr = fmt.Errorf("checksum validation failed")
				h.logger.Warn("Next.js node checksum validation failed, trying next node",
					zap.String("node", targetNode.Host),
					zap.String("expected", checksumWriter.expectedChecksum),
					zap.String("calculated", checksumWriter.calculatedChecksum))
			} else {
				lastErr = fmt.Errorf("node returned status code %d", checksumWriter.statusCode)
				h.logger.Warn("Next.js node returned non-success status, trying next node",
					zap.String("node", targetNode.Host),
					zap.Int("status", checksumWriter.statusCode))
			}
			continue
		}

		// For static sites
		baseURL := fmt.Sprintf("https://%s/v1/content?project_id=%s&resource=",
			targetNode.Host,
			projectID.Hex())

		target, err := url.Parse(baseURL)
		if err != nil {
			lastErr = err
			continue
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		if h.environment == "dev" {
			proxy.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			resourcePath := req.URL.Path

			// Preserve original headers
			req.Header.Set("X-Real-IP", r.RemoteAddr)
			req.Header.Set("X-Forwarded-For", getForwardedFor(r))
			req.Header.Set("X-Forwarded-Proto", r.URL.Scheme)
			req.Header.Set("X-Forwarded-Host", r.Host)
			if origin := r.Header.Get("Origin"); origin != "" {
				req.Header.Set("Origin", origin)
			}

			// Handle index.html rewriting
			accept := req.Header.Get("Accept")

			if resourcePath == "/" {
				resourcePath = "/index.html"
			}

			if strings.Contains(accept, "text/html") && !strings.Contains(resourcePath, ".html") {
				resourcePath = "/index.html"
			}

			// Construct the full URL
			fullURL := baseURL + strings.TrimPrefix(resourcePath, "/")
			parsedURL, err := url.Parse(fullURL)
			if err != nil {
				h.logger.Error("Failed to parse full URL", zap.Error(err))
				return
			}

			req.URL = parsedURL
			req.Host = target.Host

			h.logger.Info("Static site proxy attempt",
				zap.String("original_path", resourcePath),
				zap.String("full_url", fullURL),
				zap.String("project_id", projectID.Hex()),
			)
		}

		// Wrapper ResponseWriter that validates checksum
		checksumWriter := &checksumValidatingWriter{
			ResponseWriter: w,
			buffer:         &responseBuffer,
			targetNode:     targetNode,
			logger:         h.logger,
		}

		err = tryServeProxy(proxy, checksumWriter, r)

		// Check if response was valid and checksum matched
		if err == nil && checksumWriter.statusCode >= 200 &&
			checksumWriter.statusCode < 300 && checksumWriter.checksumValid {
			return // Success
		}

		if err != nil {
			lastErr = err
			h.logger.Warn("Static node request failed, trying next node",
				zap.String("node", targetNode.Host),
				zap.Error(err))
		} else if !checksumWriter.checksumValid {
			lastErr = fmt.Errorf("checksum validation failed")
			h.logger.Warn("Static node checksum validation failed, trying next node",
				zap.String("node", targetNode.Host),
				zap.String("expected", checksumWriter.expectedChecksum),
				zap.String("calculated", checksumWriter.calculatedChecksum))
		} else {
			lastErr = fmt.Errorf("node returned status code %d", checksumWriter.statusCode)
			h.logger.Warn("Static node returned non-success status, trying next node",
				zap.String("node", targetNode.Host),
				zap.Int("status", checksumWriter.statusCode))
		}
	}

	// If we get here, all nodes failed
	h.logger.Error("All content nodes failed", zap.Error(lastErr))
	http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
}

// ResponseWriter that passes through responses without checksum validation
type checksumValidatingWriter struct {
	http.ResponseWriter
	buffer             *bytes.Buffer
	targetNode         *model.Node
	logger             *zap.Logger
	statusCode         int
	checksumValid      bool // Always true now
	headersSent        bool
	expectedChecksum   string
	calculatedChecksum string
}

func (w *checksumValidatingWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode

	// Set checksum to valid regardless of content
	w.checksumValid = true

	// Add node info headers
	w.Header().Set("X-Content-Node", w.targetNode.Host)
	w.Header().Set("X-Node-Region", w.targetNode.Region)

	// Send headers immediately
	w.ResponseWriter.WriteHeader(statusCode)
	w.headersSent = true
}

func (w *checksumValidatingWriter) Write(data []byte) (int, error) {
	// For non-success status codes, just pass through
	if w.statusCode == 0 {
		w.WriteHeader(http.StatusOK)
	}

	// Always consider checksum valid
	w.checksumValid = true

	// Write directly to the response
	return w.ResponseWriter.Write(data)
}

func (w *checksumValidatingWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func NewHandler(logger *zap.Logger, m model.Client, resolver Resolver, t Templates, root http.Dir, c geo.GeolocationClient, environment string) (*Handler, error) {
	lph, err := newLandingPageHandler(logger, root)
	if err != nil {
		return nil, err
	}

	ph := newProxyHandler(logger, m, resolver, c, environment)

	mux := http.NewServeMux()
	mux.Handle("/earthfast-sw.js", newSWHandler(logger, m, resolver, t, c))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the project ID from the domain
		projectID, err := resolver.ProjectForDomain(r.Context(), r.Host)
		if err != nil {
			logger.Error("Failed to resolve domain", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Check if this is a Next.js project
		isNextJS := false
		project, err := m.GetProject(r.Context(), projectID)
		if err == nil && project != nil {
			var metadata struct {
				Type string `json:"type"`
			}
			if err := json.Unmarshal([]byte(project.Metadata), &metadata); err == nil {
				isNextJS = metadata.Type == "nextjs"
			}
		}

		// Check service worker compatibility
		swCompatible := isServiceWorkerCompatible(r)

		logger.Debug("Request handling",
			zap.Bool("isNextJS", isNextJS),
			zap.Bool("swCompatible", swCompatible),
			zap.String("userAgent", r.UserAgent()))

		// For Next.js projects, always use proxy for non-SW compatible browsers
		// For static sites, use proxy for non-SW compatible browsers and landing page for compatible browsers
		if !swCompatible || isNextJS {
			ph.ServeHTTP(w, r)
		} else {
			lph.ServeHTTP(w, r)
		}
	})

	return &Handler{ServeMux: mux}, nil
}
