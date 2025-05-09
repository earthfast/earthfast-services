package domain

import (
	"armada-node/geo"
	"armada-node/geo/geotest"
	"armada-node/model"
	"armada-node/model/modeltest"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"go.uber.org/zap/zaptest"
)

func TestIsServiceWorkerCompatible(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
		headers   http.Header
		want      bool
	}{
		{
			name:      "Regular Chrome Browser",
			userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			headers:   make(http.Header),
			want:      true,
		},
		{
			name:      "Telegram Browser",
			userAgent: "TelegramBot (like TwitterBot)",
			headers:   make(http.Header),
			want:      false,
		},
		{
			name:      "Googlebot",
			userAgent: "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			headers:   make(http.Header),
			want:      false,
		},
		{
			name:      "Mobile WebView",
			userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.124 Mobile Safari/537.36",
			headers: http.Header{
				"X-Requested-With": []string{"com.example.webview"},
			},
			want: false,
		},
		{
			name:      "iOS Standalone Safari",
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Safari/604.1",
			headers:   make(http.Header),
			want:      true,
		},
		{
			name:      "iOS Chrome",
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/92.0.4515.90 Mobile/15E148 Safari/604.1",
			headers:   make(http.Header),
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", nil)
			r.Header = tt.headers
			r.Header.Set("User-Agent", tt.userAgent)

			got := isServiceWorkerCompatible(r)
			if got != tt.want {
				t.Errorf("isServiceWorkerCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupTestData(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("<html><body>Test</body></html>"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestNewHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)
	m := modeltest.NewClient()
	r := NewStaticResolver(map[string]model.ID{})
	templates := Templates{
		ServiceWorker: template.Must(template.New("sw").Parse("test")),
	}

	testDir := setupTestData(t)
	root := http.Dir(testDir)
	c := &geotest.AbstractClientMock{}

	handler, err := NewHandler(logger, m, r, templates, root, c, "production")
	if err != nil {
		t.Fatalf("NewHandler() error = %v", err)
	}

	if handler.ServeMux == nil {
		t.Error("NewHandler() returned nil ServeMux")
	}
}

type mockResolver struct {
	projectID model.ID
	err       error
}

func (m *mockResolver) ProjectForDomain(ctx context.Context, domain string) (model.ID, error) {
	return m.projectID, m.err
}

func TestProxyHandler_ServeGeoHTTP(t *testing.T) {
	logger := zaptest.NewLogger(t)
	projectID := model.ID{1, 2, 3}

	tests := []struct {
		name       string
		resolver   Resolver
		nodes      []*model.Node
		wantStatus int
	}{
		{
			name:       "No nodes available",
			resolver:   &mockResolver{projectID: projectID},
			nodes:      []*model.Node{},
			wantStatus: http.StatusServiceUnavailable,
		},
		{
			name:     "With available nodes",
			resolver: &mockResolver{projectID: projectID},
			nodes: []*model.Node{
				{
					Host:      "test1.example.com",
					Region:    geo.Europe.ID,
					ProjectID: projectID,
				},
				{
					Host:      "test2.example.com",
					Region:    geo.Europe.ID,
					ProjectID: projectID,
				},
			},
			wantStatus: http.StatusBadGateway,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := modeltest.NewClient().WithContentNodes(tt.nodes...)
			h := newProxyHandler(logger, m, tt.resolver, &geotest.AbstractClientMock{}, "production")

			req := httptest.NewRequest("GET", "http://example.com/", nil)
			w := httptest.NewRecorder()

			coord := geo.Coordinate{Latitude: 51.5074, Longitude: -0.1278} // London
			h.serveGeoHTTP(coord, w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("serveGeoHTTP() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestHandler_ServeHTTP(t *testing.T) {
	logger := zaptest.NewLogger(t)
	m := modeltest.NewClient()
	r := NewStaticResolver(map[string]model.ID{})
	templates := Templates{
		ServiceWorker: template.Must(template.New("sw").Parse("test")),
	}

	testDir := setupTestData(t)
	root := http.Dir(testDir)
	c := &geotest.AbstractClientMock{}

	handler, err := NewHandler(logger, m, r, templates, root, c, "production")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		path      string
		userAgent string
		headers   http.Header
	}{
		{
			name:      "Service Worker request",
			path:      "/earthfast-sw.js",
			userAgent: "Mozilla/5.0 Chrome/91.0.4472.124",
			headers:   make(http.Header),
		},
		{
			name:      "Regular request with compatible browser",
			path:      "/",
			userAgent: "Mozilla/5.0 Chrome/91.0.4472.124",
			headers:   make(http.Header),
		},
		{
			name:      "Regular request with incompatible browser",
			path:      "/",
			userAgent: "TelegramBot",
			headers:   make(http.Header),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			req.Header = tt.headers
			req.Header.Set("User-Agent", tt.userAgent)
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			if w.Code == http.StatusInternalServerError {
				t.Errorf("Handler returned internal server error for path: %s", tt.path)
			}
		})
	}
}
