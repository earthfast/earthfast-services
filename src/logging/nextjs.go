package logging

import (
	"bufio"
	"io"
	"strings"

	"go.uber.org/zap"
)

// NextJSLogProcessor processes Next.js logs from stdout/stderr
func NextJSLogProcessor(r io.Reader, source string, projectID string, logger *zap.Logger) {
	// Use a larger buffer for the scanner to handle longer lines
	scanner := bufio.NewScanner(r)
	buffer := make([]byte, 256*1024)
	scanner.Buffer(buffer, len(buffer))

	for scanner.Scan() {
		output := scanner.Text()

		// Skip empty lines or just whitespace
		if len(strings.TrimSpace(output)) == 0 {
			continue
		}

		// Log to zap logger
		logger.Info("Next.js "+source,
			zap.String("projectId", projectID),
			zap.String("output", output))

		// Store in log buffer
		AddLog(LogTypeNextJS, source, output)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		logger.Error("Error scanning Next.js output",
			zap.String("projectId", projectID),
			zap.String("source", source),
			zap.Error(err))
	}
}
