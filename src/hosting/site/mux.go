package site

import (
	"net/http"
	"sync"

	"armada-node/model"

	"go.uber.org/zap"
)

type Site interface {
	http.Handler
	Stop()
}

type SiteProvider interface {
	HandlerForProject(id model.ID) Site
}

type ServeMux struct {
	m            model.Client
	siteProvider SiteProvider
	sites        map[model.ID]Site
	sitesLock    *sync.RWMutex
	stopped      bool
	logger       *zap.Logger
}

func NewServeMux(logger *zap.Logger, m model.Client, sp SiteProvider) *ServeMux {
	return &ServeMux{
		m:            m,
		siteProvider: sp,
		sites:        make(map[model.ID]Site),
		sitesLock:    &sync.RWMutex{},
		logger:       logger,
	}
}

func (h *ServeMux) Stop() {
	h.sitesLock.Lock()
	defer h.sitesLock.Unlock()

	for id, site := range h.sites {
		h.logger.Info("Stopping site", zap.String("projectId", id.Hex()))
		site.Stop()
	}
	h.stopped = true
}

func (h *ServeMux) ServeProjectHTTP(projectID model.ID, w http.ResponseWriter, r *http.Request) {
	// Get project information to determine project type
	proj, err := h.m.GetProject(r.Context(), projectID)
	if err != nil {
		h.logger.Error("Failed to lookup project", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if proj == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Check project type
	metadata, err := proj.ParseMetadata()
	if err != nil {
		h.logger.Error("Failed to parse project metadata",
			zap.Error(err),
			zap.String("projectId", projectID.Hex()))
		http.Error(w, "Invalid project metadata", http.StatusInternalServerError)
		return
	}

	// For static sites, restrict methods to GET/HEAD
	isNextJS := metadata.Type == model.ProjectTypeNextJS
	if !isNextJS && r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Debug log for Next.js requests
	if isNextJS {
		h.logger.Debug("Next.js request in ServeMux",
			zap.String("path", r.URL.Path),
			zap.String("method", r.Method),
			zap.String("projectId", projectID.Hex()))
	}

	// Get or create site handler for this project
	site := h.getSiteForProject(projectID)
	if site == nil {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	// Serve the request
	site.ServeHTTP(w, r)
}

// getSiteForProject gets or creates a site handler for a project
func (h *ServeMux) getSiteForProject(projectID model.ID) Site {
	// Check if we already have a site handler
	h.sitesLock.RLock()
	site, exists := h.sites[projectID]
	h.sitesLock.RUnlock()

	if exists {
		return site
	}

	// Create a new site handler if needed
	h.sitesLock.Lock()
	defer h.sitesLock.Unlock()

	// Double-check after acquiring write lock
	if site, exists = h.sites[projectID]; exists {
		return site
	}

	if h.stopped {
		return nil
	}

	// Create a new site handler
	h.logger.Info("Creating new site handler for project",
		zap.String("projectId", projectID.Hex()))
	site = h.siteProvider.HandlerForProject(projectID)
	if site != nil {
		h.sites[projectID] = site
	}

	return site
}
