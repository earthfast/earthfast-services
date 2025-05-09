package domain

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.uber.org/zap/zaptest"
)

const indexHTML = `
	<html>
		<body>Hello from the service worker landing page!</body>
	</html>
`

const mainCSS = `
	body {
		background-color: green;
	}
`

func addFile(t *testing.T, dir, name, contents string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
}

func addDir(t *testing.T, dir, name string) {
	t.Helper()
	if err := os.Mkdir(filepath.Join(dir, name), 0755); err != nil {
		t.Fatal(err)
	}
}

func TestLandingPageHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)

	scratch := t.TempDir()
	addFile(t, scratch, "index.html", indexHTML)
	addFile(t, scratch, "foo.txt", "foo")
	addDir(t, scratch, "styles")
	addFile(t, scratch, "styles/main.css", mainCSS)

	cases := []struct {
		name         string
		url          string
		headers      http.Header
		wantCode     int
		wantContains []string
	}{
		{
			name:         "index.html via /",
			url:          "http://example.com/",
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name:         "index.html via /index.html",
			url:          "http://example.com/index.html",
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name:         "index.html via unknown path",
			url:          "http://example.com/foo",
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name:         "index.html via known (but ignored) directory path",
			url:          "http://example.com/styles",
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name:         "index.html via unknown html file",
			url:          "http://example.com/faq.html",
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name: "index.html via unknown file accepting html",
			url:  "http://example.com/some/file.txt",
			headers: map[string][]string{
				"Accept": {"text/html"},
			},
			wantCode:     http.StatusOK,
			wantContains: []string{"service worker landing page"},
		},
		{
			name:         "File in root directory",
			url:          "http://example.com/foo.txt",
			wantCode:     http.StatusOK,
			wantContains: []string{"foo"},
		},
		{
			name:         "File in subdirectory",
			url:          "http://example.com/styles/main.css",
			wantCode:     http.StatusOK,
			wantContains: []string{"background-color: green"},
		},
		{
			name: "Unknown file with extension",
			url:  "http://example.com/favicon.ico",
			headers: map[string][]string{
				"Accept": {"image/jpeg"},
			},
			wantCode: http.StatusNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := newLandingPageHandler(logger, http.Dir(scratch))
			if err != nil {
				t.Fatal(err)
			}

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			r.Header = tc.headers
			h.ServeHTTP(rw, r)
			if rw.Code != tc.wantCode {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, tc.wantCode)
			}
			for _, want := range tc.wantContains {
				if got := rw.Body.String(); !strings.Contains(got, want) {
					t.Errorf("Unexpected response body: got %q, want it to contain %q", got, want)
				}
			}
		})
	}
}

func TestAcceptsHTML(t *testing.T) {
	cases := []struct {
		path   string
		accept string
		want   bool
	}{
		{path: "/", want: true},
		{path: "/foo", want: true},
		{path: "/foo/bar", want: true},
		{path: "/index.html", want: true},
		{path: "/index.htm", want: true},
		{path: "/home/index.html", want: true},
		{path: "/home/index.htm", want: true},

		{path: "/", accept: "text/html", want: true},
		{path: "/foo.txt", accept: "text/html", want: true},
		{path: "/foo.txt", accept: "text/html;q=1.0", want: true},
		{path: "/foo.txt", accept: "text/plain", want: false},
		{path: "/bar.jpg", accept: "text/plain, image/jpeg, text/html, application/json", want: true},
		{path: "/favicon.ico", accept: "application/*, image/png", want: false},
	}

	for _, tc := range cases {
		name := tc.path
		if tc.accept != "" {
			name += ":" + tc.accept
		}
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tc.path, nil)
			r.Header.Set("Accept", tc.accept)
			if got := acceptsHTML(r); got != tc.want {
				t.Errorf("Unexpected return value: got %t, want %t", got, tc.want)
			}
		})
	}
}
