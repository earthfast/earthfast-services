package site

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"armada-node/model"
	"armada-node/model/modeltest"

	"go.uber.org/zap/zaptest"
)

type testSite struct {
	http.Handler
}

func (s testSite) Stop() {}

type staticSiteProvider map[model.ID]http.Handler

func (sp staticSiteProvider) HandlerForProject(id model.ID) Site {
	return testSite{sp[id]}
}

func TestServeMux(t *testing.T) {
	logger := zaptest.NewLogger(t)
	project := &model.Project{
		ID:   model.ID{1, 2, 3},
		Name: "hello-world",
	}
	m := modeltest.NewClient().WithProjects(project)
	siteProvider := staticSiteProvider{
		project.ID: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "ok")
		}),
	}
	h := NewServeMux(logger, m, siteProvider)

	cases := []struct {
		name      string
		projectID model.ID
		wantCode  int
		wantBody  string
	}{
		{
			name:      "Known project",
			projectID: project.ID,
			wantCode:  http.StatusOK,
			wantBody:  "ok",
		},
		{
			name:      "Project not found",
			projectID: model.ID{9, 9, 9},
			wantCode:  http.StatusNotFound,
			wantBody:  http.StatusText(http.StatusNotFound),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			h.ServeProjectHTTP(tc.projectID, rw, r)
			if rw.Code != tc.wantCode {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, tc.wantCode)
			}
			if got := strings.TrimSpace(rw.Body.String()); got != tc.wantBody {
				t.Errorf("Unexpected response body: got %q, want %q", got, tc.wantBody)
			}
		})
	}
}

func TestServeMux_HTTPMethods(t *testing.T) {
	logger := zaptest.NewLogger(t)
	project := &model.Project{
		ID:   model.ID{1, 2, 3},
		Name: "hello-world",
	}
	m := modeltest.NewClient().WithProjects(project)
	siteProvider := staticSiteProvider{
		project.ID: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "ok")
		}),
	}
	h := NewServeMux(logger, m, siteProvider)

	allowed := []string{
		http.MethodGet,
		http.MethodHead,
	}
	for _, method := range allowed {
		t.Run(method, func(t *testing.T) {
			rw := httptest.NewRecorder()
			r := httptest.NewRequest(method, "/", nil)
			h.ServeProjectHTTP(project.ID, rw, r)
			if rw.Code == http.StatusMethodNotAllowed {
				t.Error("Got http.StatusMethodNotAllowed for allowed method")
			}
		})
	}

	disallowed := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
	for _, method := range disallowed {
		t.Run(method, func(t *testing.T) {
			rw := httptest.NewRecorder()
			r := httptest.NewRequest(method, "/", nil)
			h.ServeProjectHTTP(project.ID, rw, r)
			if rw.Code != http.StatusMethodNotAllowed {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, http.StatusMethodNotAllowed)
			}
		})
	}
}
