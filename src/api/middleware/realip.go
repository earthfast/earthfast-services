package middleware

import (
	"net"
	"net/http"
	"strings"
)

// RealIP is an http middleware that attempts to set the RemoteAddr field of an
// http.Request to the true IP address of the client. It looks for well-known
// headers such as X-Real-IP and X-Forwarded-For, which are typically set by
// upstream load balancers.
func RealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ip := realIP(r); ip != "" {
			r.RemoteAddr = ip
		}
		next.ServeHTTP(w, r)
	})
}

func realIP(r *http.Request) string {
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return xrip
	}

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		i := strings.Index(xff, ",")
		if i == -1 {
			i = len(xff)
		}
		return xff[:i]
	}

	if raddr, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return raddr
	}

	return r.RemoteAddr
}
