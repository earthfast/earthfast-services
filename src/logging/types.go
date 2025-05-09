package logging

import (
	"time"
)

// LogType identifies the source of logs
type LogType string

const (
	LogTypeNextJS LogType = "nextjs" // Next.js application logs
	LogTypeServer LogType = "server" // Go server logs
)

// LogEntry represents a single log line with metadata
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"` // stdout, stderr, or level (info, error, etc.)
	Message   string    `json:"message"`
}

// LogOptions provides configuration for log retrieval
type LogOptions struct {
	Type  LogType
	Limit int
}

// ProjectInfo contains information about the current project
type ProjectInfo struct {
	ProjectID string
}
