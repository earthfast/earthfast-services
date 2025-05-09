package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseRecorder struct {
	http.ResponseWriter

	size        int
	statusCode  int
	wroteHeader bool
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	if rr.wroteHeader {
		return
	}
	rr.wroteHeader = true
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func (rr *responseRecorder) Write(data []byte) (int, error) {
	rr.WriteHeader(http.StatusOK)
	n, err := rr.ResponseWriter.Write(data)
	if err == nil {
		rr.size += n
	}
	return n, err
}

func Logger(l *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrec := &responseRecorder{ResponseWriter: w}

		defer func() {
			if l == nil {
				return
			}

			l.Info("handled request",
				zap.String("host", r.Host),
				zap.String("method", r.Method),
				zap.String("uri", r.URL.Path),
				zap.String("proto", r.Proto),
				zap.Int("status", wrec.statusCode),
				zap.Int("size", wrec.size),
				zap.Duration("duration", time.Since(start)),
				zap.String("user_agent", r.UserAgent()),
				zap.String("ip", r.RemoteAddr),
				zap.Any("query", r.URL.Query()),
			)
		}()

		next.ServeHTTP(wrec, r)
	})
}
