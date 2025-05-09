package logging

import (
	"strings"
	"sync"
	"time"
)

// Default buffer size (lines)
const defaultBufferSize = 10000

// logBuffers holds all the log buffers
var (
	nextJSBuffer = newLogBuffer(defaultBufferSize)
	serverBuffer = newLogBuffer(defaultBufferSize)
)

// LogBuffer implements a thread-safe circular buffer for logs
type LogBuffer struct {
	logs     []LogEntry
	position int
	size     int
	lock     sync.RWMutex
}

// newLogBuffer creates a new log buffer with specified size
func newLogBuffer(size int) *LogBuffer {
	return &LogBuffer{
		logs: make([]LogEntry, size),
		size: size,
	}
}

// Add appends a log entry to the buffer
func (b *LogBuffer) Add(source, message string) {
	// Skip empty messages
	if len(strings.TrimSpace(message)) == 0 {
		return
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	// Add log entry to circular buffer
	b.logs[b.position] = LogEntry{
		Timestamp: time.Now(),
		Source:    source,
		Message:   message,
	}

	// Move position for next entry
	b.position = (b.position + 1) % b.size
}

// Get retrieves logs from the buffer
func (b *LogBuffer) Get(limit int) []LogEntry {
	b.lock.RLock()
	defer b.lock.RUnlock()

	// If no logs yet, return empty array
	if b.logs[0].Timestamp.IsZero() && b.position == 0 {
		return []LogEntry{}
	}

	// Determine actual log count in buffer
	var logCount int
	if b.logs[b.position].Timestamp.IsZero() {
		// Buffer isn't full yet
		logCount = b.position
	} else {
		// Buffer is full
		logCount = b.size
	}

	if limit > 0 && limit < logCount {
		logCount = limit
	}

	result := make([]LogEntry, logCount)

	// Copy logs in chronological order
	for i := 0; i < logCount; i++ {
		pos := (b.position - logCount + i + b.size) % b.size
		result[i] = b.logs[pos]
	}

	return result
}

// Clear empties the buffer
func (b *LogBuffer) Clear() {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.logs = make([]LogEntry, b.size)
	b.position = 0
}

// Public API functions

// AddLog adds a log to the appropriate buffer
func AddLog(logType LogType, source, message string) {
	switch logType {
	case LogTypeNextJS:
		nextJSBuffer.Add(source, message)
	case LogTypeServer:
		serverBuffer.Add(source, message)
	}
}

// GetLogs retrieves logs from the specified buffer
func GetLogs(options LogOptions) []LogEntry {
	switch options.Type {
	case LogTypeNextJS:
		return nextJSBuffer.Get(options.Limit)
	case LogTypeServer:
		return serverBuffer.Get(options.Limit)
	default:
		return []LogEntry{}
	}
}

// ClearLogs empties the specified buffer
func ClearLogs(logType LogType) {
	switch logType {
	case LogTypeNextJS:
		nextJSBuffer.Clear()
	case LogTypeServer:
		serverBuffer.Clear()
	}
}
