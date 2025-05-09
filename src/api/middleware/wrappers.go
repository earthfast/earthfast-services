package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// WithLogger returns a middleware function that adds logging
func WithLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return Logger(logger, next)
	}
}

// WithCORS returns a middleware function that adds CORS headers
func WithCORS() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return CORS(next)
	}
}

// WithOpenCensus returns a middleware function that adds OpenCensus metrics
func WithOpenCensus() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return OpenCensus(next)
	}
}

// WithRealIP returns a middleware function that adds RealIP processing
func WithRealIP() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return RealIP(next)
	}
}

// Chain combines multiple middleware functions into a single middleware
func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
