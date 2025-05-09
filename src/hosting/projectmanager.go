package hosting

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"armada-node/hosting/site"
	"armada-node/model"

	"go.uber.org/zap"
)

// ProjectManager handles lifecycle events for projects
type ProjectManager struct {
	logger       *zap.Logger
	client       model.Client
	hostingDir   string
	siteProvider site.SiteProvider
	projects     map[model.ID]bool
	mutex        sync.Mutex
}

// NewProjectManager creates a new project manager
func NewProjectManager(
	logger *zap.Logger,
	client model.Client,
	hostingDir string,
	siteProvider site.SiteProvider,
) *ProjectManager {
	return &ProjectManager{
		logger:       logger,
		client:       client,
		hostingDir:   hostingDir,
		siteProvider: siteProvider,
		projects:     make(map[model.ID]bool),
	}
}

// InitializeProject initializes a project, particularly Next.js projects
func (pm *ProjectManager) InitializeProject(ctx context.Context, projectID model.ID) error {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	// Check if already initialized
	if initialized, exists := pm.projects[projectID]; exists && initialized {
		pm.logger.Debug("Project already initialized",
			zap.String("projectId", projectID.Hex()))
		return nil
	}

	pm.logger.Info("Initializing project", zap.String("projectId", projectID.Hex()))

	// Fetch project details
	project, err := pm.client.GetProject(ctx, projectID)
	if err != nil {
		return err
	}
	if project == nil {
		pm.logger.Warn("Project not found", zap.String("projectId", projectID.Hex()))
		return nil
	}

	// Parse metadata
	metadata, err := project.ParseMetadata()
	if err != nil {
		pm.logger.Error("Failed to parse project metadata",
			zap.Error(err),
			zap.String("projectId", projectID.Hex()))
		return err
	}

	// Initialize Next.js projects proactively and wait for them to start
	if metadata.Type == model.ProjectTypeNextJS {
		pm.logger.Info("Initializing Next.js project",
			zap.String("projectId", projectID.Hex()),
			zap.String("name", project.Name))

		// Force immediate initialization
		_ = pm.siteProvider.HandlerForProject(projectID)

		// Allow a little time for the server to start
		time.Sleep(1 * time.Second)

		// Try a test connection to verify the server is running
		if metadata.Port > 0 {
			hostname, _ := os.Hostname()
			if hostname == "" {
				hostname = "127.0.0.1"
			}

			client := http.Client{Timeout: 500 * time.Millisecond}
			_, err := client.Get(fmt.Sprintf("http://%s:%d/", hostname, metadata.Port))
			if err != nil {
				pm.logger.Warn("Next.js server may not be fully started yet",
					zap.Error(err),
					zap.String("projectId", projectID.Hex()))
				// Continue anyway - it might just be starting up
			}
		}
	}

	// Mark as initialized
	pm.projects[projectID] = true
	return nil
}

// InitializeFromNode initializes projects for a node at startup
func (pm *ProjectManager) InitializeFromNode(ctx context.Context, node *model.Node) error {
	if node == nil || node.ProjectID.IsZero() {
		pm.logger.Info("No project assigned to node")
		return nil
	}
	return pm.InitializeProject(ctx, node.ProjectID)
}

// Stop stops all managed projects
func (pm *ProjectManager) Stop() {
	pm.logger.Info("Stopping all projects")
	// The site provider's SiteMux will handle stopping all sites
}
