package content

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"armada-node/metering/uptime"
	"armada-node/model"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type probeRequest struct {
	URL      string
	Checksum string
}

type fakeUptimeMeter struct {
	probes       []probeRequest
	store        *uptime.InMemoryStore
	requestCount uint32
}

func newFakeUptimeMeter(logger *zap.Logger) *fakeUptimeMeter {
	return &fakeUptimeMeter{
		store: uptime.NewInMemoryStore(logger, 100),
	}
}

func (m *fakeUptimeMeter) Probe(url *url.URL, checksum string, resultCh chan<- bool) bool {
	m.probes = append(m.probes, probeRequest{
		URL:      url.String(),
		Checksum: checksum,
	})
	return true
}

func (m *fakeUptimeMeter) Results(ctx context.Context, start, end time.Time, fn func(uptime.IntervalData) error) error {
	return m.store.Range(ctx, start, end, fn)
}

func (m *fakeUptimeMeter) IncrementRequestCount(delta uint32) {
	m.requestCount += delta
}

type fakeSiteMux struct {
	projectID model.ID
	http.Handler
}

func newFakeSiteMux(projectID model.ID, h http.Handler) *fakeSiteMux {
	return &fakeSiteMux{
		projectID: projectID,
		Handler:   h,
	}
}

func (m *fakeSiteMux) ServeProjectHTTP(projectID model.ID, w http.ResponseWriter, r *http.Request) {
	if projectID != m.projectID {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	m.ServeHTTP(w, r)
}

func TestHandler_Content(t *testing.T) {
	project := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node := &model.Node{
		ID:        modeltest.RandomID(t),
		Host:      "node0.armadanetwork.com",
		ProjectID: project.ID,
	}

	cases := []struct {
		name                   string
		query                  url.Values
		responseBody           []byte
		wantCode               int
		wantUptimeProbes       []probeRequest
		wantUptimeRequestCount uint32
	}{
		{
			name: "Missing project_id",
			query: url.Values{
				"resource": {"/index.html"},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Missing resource",
			query: url.Values{
				"project_id": {project.ID.Hex()},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Project not registered",
			query: url.Values{
				"project_id": {modeltest.RandomID(t).Hex()},
				"resource":   {"/index.html"},
			},
			wantCode: http.StatusGone,
		},
		{
			name: "Resource not found",
			query: url.Values{
				"project_id": {project.ID.Hex()},
				"resource":   {"/foo.html"},
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "Successful request",
			query: url.Values{
				"project_id": {project.ID.Hex()},
				"resource":   {"/index.html"},
			},
			responseBody:           []byte("ok"),
			wantCode:               http.StatusOK,
			wantUptimeRequestCount: 1,
		},
		{
			name: "Retry probe",
			query: url.Values{
				"project_id": {project.ID.Hex()},
				"resource":   {"/index.html"},
				"retry":      {"node1.armadanetwork.com"},
			},
			responseBody: []byte("ok"),
			wantCode:     http.StatusOK,
			wantUptimeProbes: []probeRequest{
				{
					URL:      fmt.Sprintf("https://node1.armadanetwork.com/v1/content?project_id=%s&resource=%%2Findex.html", project.ID.Hex()),
					Checksum: "7a85f4764bbd6daf1c3545efbbf0f279a6dc0beb",
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			m := modeltest.NewClient().WithProjects(project)
			self := modeltest.ReadOnlyNode(node)
			uptime := newFakeUptimeMeter(logger)

			hostingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				wantURLStr := tc.query["resource"][0]
				if gotURLStr := r.URL.String(); gotURLStr != wantURLStr {
					t.Errorf("Unexpected URL: got %s, want %s", gotURLStr, wantURLStr)
				}
				if tc.responseBody == nil {
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				}
				w.Write(tc.responseBody)
			})
			siteMux := newFakeSiteMux(project.ID, hostingHandler)
			h := NewHandler(logger, m, uptime, siteMux, self)

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/v1/content", nil)
			r.URL.RawQuery = tc.query.Encode()
			h.ServeHTTP(rw, r)
			if rw.Code != tc.wantCode {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, tc.wantCode)
			}
			if diff := cmp.Diff(tc.wantUptimeProbes, uptime.probes); diff != "" {
				t.Errorf("Retry probe mismatch (-want +got): %s", diff)
			}
			if tc.wantUptimeRequestCount != uptime.requestCount {
				t.Errorf("Unexpected uptime request count: got %d, want %d", tc.wantUptimeRequestCount, uptime.requestCount)
			}
		})
	}
}

func TestChecksumResponseWriter(t *testing.T) {
	cases := []struct {
		name         string
		code         int
		body         string
		wantChecksum string
	}{
		{
			name:         "200",
			code:         200,
			body:         "Hello, world!",
			wantChecksum: "943a702d06f34599aee1f8da8ef9f7296031d699",
		},
		{
			name:         "404",
			code:         404,
			body:         http.StatusText(http.StatusNotFound),
			wantChecksum: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			rw := newSHA1ResponseWriter(w)

			rw.WriteHeader(tc.code)
			if _, err := rw.Write([]byte(tc.body)); err != nil {
				t.Fatal(err)
			}

			if gotCode := w.Result().StatusCode; gotCode != tc.code {
				t.Errorf("Underlying http.ResponseWriter has wrong status code: got %d, want %d", gotCode, tc.code)
			}
			if gotBody := w.Body.String(); gotBody != tc.body {
				t.Errorf("Underlying http.ResponseWriter has wrong body: got %q, want %q", gotBody, tc.body)
			}
			if gotChecksum := rw.Checksum(); gotChecksum != tc.wantChecksum {
				t.Errorf("Incorrect checksum: got %s, want %s", gotChecksum, tc.wantChecksum)
			}
		})
	}
}
