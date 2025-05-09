package content

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

type StatuszHandler struct {
	self     ReadOnlyNode
	initTime time.Time

	logger *zap.Logger
}

func NewStatuszHandler(logger *zap.Logger, self ReadOnlyNode) *StatuszHandler {
	return &StatuszHandler{
		self:     self,
		initTime: time.Now(),

		logger: logger,
	}
}

func (h *StatuszHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Uptime    string
		NodeID    string
		Host      string
		ProjectID string
		GitSHA    string
	}{
		Uptime:    time.Since(h.initTime).Truncate(time.Second).String(),
		NodeID:    h.self.ID().Hex(),
		Host:      h.self.Host(),
		ProjectID: h.self.ProjectID().Hex(),
		GitSHA:    os.Getenv("GIT_SHA"),
	}
	h.sendJSON(data, w, r)
}

func (h *StatuszHandler) sendJSON(data interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Error writing JSON response", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
