package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateCaptureLogger creates a logger that captures logs to our buffer
func CreateCaptureLogger(config zap.Config) (*zap.Logger, error) {
	// Build the standard logger first
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	// Create a hook that sends logs to our buffer
	logger = logger.WithOptions(zap.Hooks(func(entry zapcore.Entry) error {
		// Capture log to our buffer
		AddLog(LogTypeServer, entry.Level.String(), entry.Message)
		return nil
	}))

	return logger, nil
}
