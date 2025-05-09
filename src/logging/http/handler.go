package http

import (
	"armada-node/logging"
	"armada-node/model/dynamic"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// LogsHandler handles requests for logs
func LogsHandler(logger *zap.Logger, self *dynamic.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get log type parameter (nextjs or server)
		logTypeStr := r.URL.Query().Get("type")
		if logTypeStr == "" {
			logTypeStr = "nextjs" // Default to Next.js logs
		}

		var logType logging.LogType
		switch logTypeStr {
		case "nextjs":
			logType = logging.LogTypeNextJS
		case "server":
			logType = logging.LogTypeServer
		default:
			http.Error(w, "Invalid log type. Use 'nextjs' or 'server'.", http.StatusBadRequest)
			return
		}

		// Get format parameter (json or text)
		format := r.URL.Query().Get("format")
		if format == "" {
			format = "json" // Default to JSON
		}

		// Get limit parameter
		limitStr := r.URL.Query().Get("limit")
		limit := 100 // Default limit
		if limitStr != "" {
			if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
				limit = parsedLimit
			}
		}

		// Get the current node state for project info
		node := self.Get()
		projectID := "none"
		if node != nil && !node.ProjectID.IsZero() {
			projectID = node.ProjectID.Hex()
		}

		// Get logs
		logs := logging.GetLogs(logging.LogOptions{
			Type:  logType,
			Limit: limit,
		})

		// Return response based on format
		if format == "text" {
			w.Header().Set("Content-Type", "text/plain")

			// Add header with info
			if logType == logging.LogTypeNextJS {
				fmt.Fprintf(w, "Next.js Logs for Project: %s\n", projectID)
			} else {
				fmt.Fprintf(w, "Server Logs\n")
			}
			fmt.Fprintf(w, "-------------------------------------------\n")

			// Format logs as text for easier reading
			for _, entry := range logs {
				fmt.Fprintf(w, "[%s] [%s] %s\n",
					entry.Timestamp.Format(time.RFC3339),
					entry.Source,
					entry.Message)
			}
			return
		}

		// Return JSON response (default)
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			LogType   string             `json:"logType"`
			ProjectID string             `json:"projectId,omitempty"`
			Logs      []logging.LogEntry `json:"logs"`
		}{
			LogType: string(logType),
			Logs:    logs,
		}

		// Only include projectID for Next.js logs
		if logType == logging.LogTypeNextJS {
			response.ProjectID = projectID
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("Failed to encode logs", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
