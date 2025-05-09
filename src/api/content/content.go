package content

import (
	"armada-node/api"
	"armada-node/model"
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"hash"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

func (h *Handler) getClientIP(r *http.Request) net.IP {
	// Add debug logging
	h.logger.Debug("Getting client IP",
		zap.String("cf_connecting_ip", r.Header.Get("CF-Connecting-IP")),
		zap.String("cf_ip_country", r.Header.Get("CF-IPCountry")),
		zap.String("x_forwarded_for", r.Header.Get("X-Forwarded-For")),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("true_client_ip", r.Header.Get("True-Client-IP")),
		zap.String("x_real_ip", r.Header.Get("X-Real-IP")),
	)

	// Check Cloudflare-specific headers
	if cfIP := r.Header.Get("CF-Connecting-IP"); cfIP != "" {
		if ip := net.ParseIP(strings.TrimSpace(cfIP)); ip != nil {
			h.logger.Debug("Using CF-Connecting-IP", zap.String("ip", ip.String()))
			return ip
		}
	}

	// Check True-Client-IP header first (if present)
	if trueClientIP := r.Header.Get("True-Client-IP"); trueClientIP != "" {
		if ip := net.ParseIP(strings.TrimSpace(trueClientIP)); ip != nil {
			h.logger.Debug("Using True-Client-IP", zap.String("ip", ip.String()))
			return ip
		}
	}

	// Check X-Forwarded-For header
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ips := strings.Split(forwardedFor, ",")
		if len(ips) > 0 {
			// Get the first (original) IP in the chain
			if ip := net.ParseIP(strings.TrimSpace(ips[0])); ip != nil {
				h.logger.Debug("Using X-Forwarded-For", zap.String("ip", ip.String()))
				return ip
			}
		}
	}

	// Check X-Real-IP
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		if ip := net.ParseIP(strings.TrimSpace(realIP)); ip != nil {
			h.logger.Debug("Using X-Real-IP", zap.String("ip", ip.String()))
			return ip
		}
	}

	// Fall back to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// Try RemoteAddr directly in case it's just an IP
		if ip := net.ParseIP(r.RemoteAddr); ip != nil {
			h.logger.Debug("Using RemoteAddr (direct)", zap.String("ip", ip.String()))
			return ip
		}
		h.logger.Warn("Failed to parse RemoteAddr, using default",
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
		)
		return net.IPv4(0, 0, 0, 0)
	}

	if ip := net.ParseIP(host); ip != nil {
		h.logger.Debug("Using RemoteAddr (parsed)", zap.String("ip", ip.String()))
		return ip
	}

	h.logger.Warn("No valid IP found, using default")
	return net.IPv4(0, 0, 0, 0)
}

type responseBuffer struct {
	headers    http.Header
	body       bytes.Buffer
	statusCode int
}

func (r *responseBuffer) Header() http.Header {
	return r.headers
}

func (r *responseBuffer) Write(b []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = http.StatusOK
	}
	return r.body.Write(b)
}

func (r *responseBuffer) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func (h *Handler) serveContentWithChecksum(w http.ResponseWriter, r *http.Request) string {
	// First, create a buffer to capture the response without sending it
	buffer := &responseBuffer{
		headers: make(http.Header),
	}

	// Serve content to the buffer first to calculate what it would be
	h.serveContent(buffer, r)

	// Default empty checksum
	checksum := ""

	// Check if the response was successful
	if buffer.statusCode == http.StatusOK {
		// Calculate checksum of the buffered content
		checksum = fmt.Sprintf("%x", sha1.Sum(buffer.body.Bytes()))

		// Set the checksum in the original response headers
		w.Header().Set("X-Content-Checksum", checksum)
		w.Header().Set("ETag", fmt.Sprintf(`"%s"`, checksum))

		// Copy all other headers from buffer to the real response
		for k, vv := range buffer.headers {
			if k != "X-Content-Checksum" && k != "ETag" {
				for _, v := range vv {
					w.Header().Set(k, v)
				}
			}
		}

		// Write the status code
		w.WriteHeader(buffer.statusCode)

		// Write the content
		w.Write(buffer.body.Bytes())
	} else {
		// For non-200 responses, just copy the response as-is
		for k, vv := range buffer.headers {
			for _, v := range vv {
				w.Header().Set(k, v)
			}
		}
		w.WriteHeader(buffer.statusCode)
		if buffer.body.Len() > 0 {
			w.Write(buffer.body.Bytes())
		}
	}

	// Return the checksum (will be empty for non-200 responses)
	return checksum
}

func (h *Handler) contentHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Get project ID first since we need it for metrics
	projectID, err := model.ParseID(r.URL.Query().Get("project_id"))
	if err != nil {
		h.logAndRecordError("Invalid query parameter: project_id", http.StatusBadRequest, r, err)
		http.Error(w, "Invalid query parameter: project_id", http.StatusBadRequest)
		return
	}

	// Extract client IP
	clientIP := h.getClientIP(r)

	// Track status code
	statusCode := http.StatusOK

	// This is not a retry, so serve the request normally.
	retry := r.URL.Query().Get("retry")
	if retry == "" {
		// Use a custom response writer to track status code
		customWriter := &statusResponseWriter{ResponseWriter: w}

		// Serve with checksum in headers
		h.serveContentWithChecksum(customWriter, r)

		statusCode = customWriter.statusCode

		if statusCode == http.StatusOK {
			h.uptime.IncrementRequestCount(1)
		}
	} else {
		// For retry requests, track the status and get the checksum for uptime probe
		customWriter := &statusResponseWriter{ResponseWriter: w}

		// Serve content with checksum and capture the returned checksum
		checksum := h.serveContentWithChecksum(customWriter, r)

		// If we got a valid checksum (which means status was 200), send the uptime probe
		if checksum != "" {
			h.sendUptimeProbe(retry, checksum, r)
		}

		statusCode = customWriter.statusCode
	}

	api.RecordContentRequest(
		projectID.Hex(),
		r.Host,
		clientIP,
		r.URL.Path,
		statusCode,
		time.Since(start),
	)
}

func (h *Handler) serveContent(w http.ResponseWriter, r *http.Request) {
	projectID, err := model.ParseID(r.URL.Query().Get("project_id"))
	if err != nil {
		h.logAndRecordError("Invalid query parameter: project_id", http.StatusBadRequest, r, err)
		http.Error(w, "Invalid query parameter: project_id", http.StatusBadRequest)
		return
	}
	if projectID != h.self.ProjectID() {
		h.logAndRecordError("Node not reserved by requested project", http.StatusGone, r, errors.New("node-project mismatch"))
		http.Error(w, "This node is not reserved by the requested project", http.StatusGone)
		return
	}

	// Check if this is a Next.js project
	isNextJS := false
	project, err := h.m.GetProject(r.Context(), projectID)
	if err == nil && project != nil {
		metadata, err := project.ParseMetadata()
		if err == nil && metadata.Type == model.ProjectTypeNextJS {
			isNextJS = true
		}
	}

	// Special handling for Next.js paths
	resource := r.URL.Query().Get("resource")
	isNextJSPath := strings.HasPrefix(resource, "/_next/") || strings.HasPrefix(resource, "_next/") ||
		resource == "/" || resource == "" || resource == "index.html"

	// If this is a Next.js project and path, handle it specially
	if isNextJS && isNextJSPath {
		h.logger.Debug("Next.js resource request detected",
			zap.String("resource", resource),
			zap.String("projectId", projectID.Hex()))

		// Create a new request with the appropriate path
		resourcePath := resource
		if !strings.HasPrefix(resourcePath, "/") {
			resourcePath = "/" + resourcePath
		}

		// If root is requested, make it explicitly "/"
		if resourcePath == "" {
			resourcePath = "/"
		}

		// Parse into URL and create new request
		resourceURL, err := url.Parse(resourcePath)
		if err != nil {
			h.logAndRecordError("Invalid Next.js resource path", http.StatusBadRequest, r, err)
			http.Error(w, "Invalid Next.js resource path", http.StatusBadRequest)
			return
		}

		// Create a new request for the Next.js path
		nextjsReq := r.Clone(r.Context())
		nextjsReq.URL = resourceURL

		// Serve the Next.js request
		h.siteMux.ServeProjectHTTP(projectID, w, nextjsReq)
		return
	}

	// Standard resource handling for non-Next.js paths
	if resource == "" {
		h.logAndRecordError("Missing resource parameter", http.StatusBadRequest, r, errors.New("missing resource"))
		http.Error(w, "Missing required query parameter: resource", http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(resource, "/") {
		resource = "/" + resource
	}

	// Convert the resource parameter into a *url.URL so that:
	//   1. We know it's well-formed.
	//   2. We can use the new *url.URL to invoke siteMux.
	resourceURL, err := url.Parse(resource)
	if err != nil {
		h.logAndRecordError("Invalid resource URL", http.StatusBadRequest, r, err)
		http.Error(w, "Invalid query parameter: resource", http.StatusBadRequest)
		return
	}

	// Create a new HTTP Request that simulates the desired resource being requested
	// directly, which can be used to invoke siteMux.
	contentReq := r.Clone(r.Context())
	contentReq.Host = ""
	contentReq.URL = resourceURL
	h.siteMux.ServeProjectHTTP(projectID, w, contentReq)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Debug all requests
	h.logger.Debug("Request received",
		zap.String("method", r.Method),
		zap.String("path", path),
		zap.String("query", r.URL.RawQuery),
		zap.String("host", r.Host))

	if strings.HasPrefix(path, "/v1/content") {
		h.contentHandler(w, r)
		return
	} else if strings.HasPrefix(path, "/v1/uptime") {
		h.uptimeHandler(w, r)
		return
	}

	projectID := h.self.ProjectID()
	if projectID.IsZero() {
		http.NotFound(w, r)
		return
	}

	if strings.HasPrefix(path, "/_next/image") || strings.HasPrefix(path, "/nextjs/_next/image") {
		// Extract the image URL from the query
		imageURL := r.URL.Query().Get("url")
		if imageURL == "" {
			h.logger.Error("Missing url parameter",
				zap.String("path", path),
				zap.String("query", r.URL.RawQuery))
			http.Error(w, "url parameter is required", http.StatusBadRequest)
			return
		}

		// URL-decode the image path if needed
		decodedURL, err := url.QueryUnescape(imageURL)
		if err != nil {
			h.logger.Error("Failed to decode URL", zap.Error(err))
			http.Error(w, "Invalid url parameter", http.StatusBadRequest)
			return
		}

		h.logger.Debug("Next.js image request intercepted",
			zap.String("path", path),
			zap.String("decodedURL", decodedURL),
			zap.String("width", r.URL.Query().Get("w")),
			zap.String("quality", r.URL.Query().Get("q")))

		// Instead of redirecting, handle the content request directly
		contentURL := url.URL{
			Path: "/v1/content",
			RawQuery: fmt.Sprintf("project_id=%s&resource=%s",
				url.QueryEscape(projectID.Hex()),
				url.QueryEscape(decodedURL)),
		}

		contentReq, err := http.NewRequestWithContext(
			r.Context(),
			"GET",
			contentURL.String(),
			nil,
		)
		if err != nil {
			h.logger.Error("Failed to create content request", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Copy relevant headers
		for k, v := range r.Header {
			if k != "Host" {
				contentReq.Header[k] = v
			}
		}

		h.logger.Debug("Serving image through content handler",
			zap.String("contentURL", contentURL.String()))

		// Handle the request directly instead of redirecting
		h.contentHandler(w, contentReq)
		return
	}

	// Handle the /nextjs prefix case
	if strings.HasPrefix(path, "/nextjs") {
		h.logger.Debug("Next.js prefixed request detected",
			zap.String("path", path),
			zap.String("projectId", projectID.Hex()))

		// Create a modified path without the /nextjs prefix
		modifiedPath := strings.TrimPrefix(path, "/nextjs")
		if modifiedPath == "" {
			modifiedPath = "/"
		}

		resourceURL, _ := url.Parse(modifiedPath)

		// Clone the request with the new path
		nextjsReq := r.Clone(r.Context())
		nextjsReq.URL = resourceURL

		// Forward to the project handler
		h.siteMux.ServeProjectHTTP(projectID, w, nextjsReq)
		return
	} else if strings.HasPrefix(path, "/_next/") {
		// Direct Next.js asset requests
		h.logger.Debug("Next.js asset request detected",
			zap.String("path", path),
			zap.String("projectId", projectID.Hex()))

		resourceURL, _ := url.Parse(path)
		nextjsReq := r.Clone(r.Context())
		nextjsReq.URL = resourceURL
		h.siteMux.ServeProjectHTTP(projectID, w, nextjsReq)
		return
	} else if !strings.HasPrefix(path, "/v1/") && !strings.HasPrefix(path, "/metrics") &&
		!strings.HasPrefix(path, "/statusz") && !strings.HasPrefix(path, "/debug/") {
		// All other non-API paths should be treated as Next.js routes for simplicity
		h.logger.Debug("Redirecting general path to /nextjs prefix",
			zap.String("path", path),
			zap.String("projectId", projectID.Hex()))

		// Redirect to the /nextjs/ prefix
		http.Redirect(w, r, "/nextjs"+path, http.StatusTemporaryRedirect)
		return
	}

	// Default to 404 for unknown paths
	http.NotFound(w, r)
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

func (w *statusResponseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *statusResponseWriter) Write(data []byte) (int, error) {
	w.WriteHeader(http.StatusOK)
	return w.ResponseWriter.Write(data)
}

type checksumResponseWriter struct {
	*statusResponseWriter
	sum hash.Hash
}

func newSHA1ResponseWriter(w http.ResponseWriter) *checksumResponseWriter {
	return &checksumResponseWriter{
		statusResponseWriter: &statusResponseWriter{ResponseWriter: w},
		sum:                  sha1.New(),
	}
}

func (w *checksumResponseWriter) Write(data []byte) (n int, err error) {
	if n, err = w.statusResponseWriter.Write(data); err != nil {
		return
	}
	if w.statusResponseWriter.statusCode != http.StatusOK {
		return
	}
	_, err = w.sum.Write(data)
	return
}

func (w *checksumResponseWriter) Checksum() string {
	if w.statusResponseWriter.statusCode != http.StatusOK {
		return ""
	}
	return fmt.Sprintf("%x", w.sum.Sum(nil))
}

func (h *Handler) logAndRecordError(message string, statusCode int, r *http.Request, err error) {
	h.logger.Error(message,
		zap.Error(err),
		zap.String("path", r.URL.Path),
		zap.Int("status", statusCode),
	)
}

func (h *Handler) sendUptimeProbe(host, checksum string, r *http.Request) {
	probeURL := *r.URL
	probeURL.Scheme = "https"
	probeURL.Host = host
	q := probeURL.Query()
	q.Del("retry")
	probeURL.RawQuery = q.Encode()
	h.uptime.Probe(&probeURL, checksum, nil)
}
