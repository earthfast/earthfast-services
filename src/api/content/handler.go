package content

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"armada-node/metering/uptime"
	"armada-node/model"

	"go.uber.org/zap"
)

type ReadOnlyNode interface {
	ID() model.ID
	Host() string
	ProjectID() model.ID
}

type SiteMux interface {
	ServeProjectHTTP(model.ID, http.ResponseWriter, *http.Request)
}

type UptimeMeter interface {
	Probe(url *url.URL, checksum string, resultCh chan<- bool) bool
	Results(ctx context.Context, start, end time.Time, fn func(uptime.IntervalData) error) error
	IncrementRequestCount(delta uint32)
}

type Handler struct {
	*http.ServeMux

	m       model.Client
	self    ReadOnlyNode
	uptime  UptimeMeter
	siteMux SiteMux

	logger *zap.Logger
}

func NewHandler(logger *zap.Logger, m model.Client, uptime UptimeMeter, siteMux SiteMux, self ReadOnlyNode) *Handler {
	h := &Handler{
		ServeMux: http.NewServeMux(),
		m:        m,
		self:     self,
		uptime:   uptime,
		siteMux:  siteMux,
		logger:   logger,
	}

	h.ServeMux.Handle("/v1/content", http.HandlerFunc(h.contentHandler))
	h.ServeMux.Handle("/v1/uptime", http.HandlerFunc(h.uptimeHandler))
	h.ServeMux.Handle("/nextjs/", http.HandlerFunc(h.nextJSHandler))

	return h
}

// nextJSHandler handles requests for Next.js applications
func (h *Handler) nextJSHandler(w http.ResponseWriter, r *http.Request) {
	// Get the project ID from the node
	projectID := h.self.ProjectID()
	if projectID.IsZero() {
		h.logger.Warn("No project assigned to this node for Next.js request")
		http.NotFound(w, r)
		return
	}

	// Strip /nextjs prefix from path
	originalPath := r.URL.Path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/nextjs")
	if r.URL.Path == "" {
		r.URL.Path = "/"
	}

	h.logger.Debug("Next.js request",
		zap.String("originalPath", originalPath),
		zap.String("newPath", r.URL.Path),
		zap.String("projectId", projectID.Hex()))

	// Forward to the site mux
	h.siteMux.ServeProjectHTTP(projectID, w, r)
}

func (h *Handler) sendJSON(data interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Error writing JSON response", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
