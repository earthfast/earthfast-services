package site

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"armada-node/model"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

type fakeVersion struct {
	project *model.Project
	started bool
	stopped bool
	deleted bool
}

func newFakeVersion(p *model.Project) *fakeVersion {
	pCopy := *p
	return &fakeVersion{
		project: &pCopy,
	}
}

func (v *fakeVersion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, v.project.Checksum)
}

func (v *fakeVersion) Start() {
	v.started = true
}

func (v *fakeVersion) Stop() {
	v.stopped = true
}

func (v *fakeVersion) Delete() error {
	v.deleted = true
	return nil
}

func (v *fakeVersion) IsProject(p *model.Project) bool {
	return cmp.Equal(v, p)
}

func (v *fakeVersion) String() string {
	return fmt.Sprintf("fakeVersion[name=%s, chksum=%s]", v.project.Name, v.project.Checksum)
}

type testVersionProvider struct {
	latest *fakeVersion
}

func (vp *testVersionProvider) VersionForProject(p *model.Project) (Version, error) {
	if vp.latest != nil && vp.latest.IsProject(p) {
		return vp.latest, nil
	}
	vp.latest = newFakeVersion(p)
	return vp.latest, nil
}

func TestHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)
	proj := &model.Project{
		ID:       modeltest.RandomID(t),
		Name:     "hello-world",
		Content:  "http://host.com/data.tar.gz",
		Checksum: "111",
	}
	m := modeltest.NewClient().WithProjects(proj)
	h := NewHandler(logger, m, &testVersionProvider{}, proj.ID)

	doRequest := func(wantBody string) {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		h.ServeHTTP(rw, req)
		if rw.Code != http.StatusOK {
			t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, http.StatusOK)
		}
		if got := rw.Body.String(); got != wantBody {
			t.Errorf("Unexpected response body: got %q, want %q", got, wantBody)
		}
	}

	// Make an initial request expecting v1 to serve.
	doRequest("111")

	// Simulate a new version becoming available.
	proj.Checksum = "222"

	// Make another request expecting the cached version (v1) to serve.
	doRequest("111")

	// Force the live version to become stale.
	h.liveStaleTime = time.Now().Add(-time.Second)

	// Make a final request expecting the new version (v2) to serve.
	doRequest("222")
}

func TestHandler_VersionLifecycle(t *testing.T) {
	logger := zaptest.NewLogger(t)
	proj := &model.Project{
		ID:       modeltest.RandomID(t),
		Name:     "hello-world",
		Content:  "http://host.com/data.tar.gz",
		Checksum: "111",
	}
	m := modeltest.NewClient().WithProjects(proj)
	vp := &testVersionProvider{}
	h := NewHandler(logger, m, vp, proj.ID)

	doRequest := func() {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		h.ServeHTTP(rw, req)
	}

	checkLifecycle := func(v *fakeVersion, started, stopped, deleted bool) error {
		if v == nil {
			return errors.New("missing version")
		}
		if v.started != started {
			return fmt.Errorf("version.started = %t, want %t", v.started, started)
		}
		if v.stopped != stopped {
			return fmt.Errorf("version.stopped = %t, want %t", v.stopped, stopped)
		}
		if v.deleted != deleted {
			return fmt.Errorf("version.deleted = %t, want %t", v.deleted, deleted)
		}
		return nil
	}

	// Make an initial request expecting v1 to serve.
	doRequest()

	// Expect v1 to have been started.
	v1 := vp.latest
	if err := checkLifecycle(v1, true, false, false); err != nil {
		t.Errorf("Incorrect version lifecycle (v1): %v", err)
	}

	// Simulate a new version becoming available.
	proj.Checksum = "222"

	// Force the live version to become stale.
	h.liveStaleTime = time.Now().Add(-time.Second)

	// Make another request expecting the new version (v2) to serve.
	doRequest()

	// Expect v2 to have been started.
	v2 := vp.latest
	if err := checkLifecycle(v2, true, false, false); err != nil {
		t.Errorf("Incorrect version lifecycle (v2): %v", err)
	}

	// Expect v1 to have been stopped and deleted.
	if err := checkLifecycle(v1, true, true, true); err != nil {
		t.Errorf("Incorrect version lifecycle (v1): %v", err)
	}

	// Stop the handler.
	h.Stop()

	// Expect v2 to have been stopped but NOT deleted.
	if err := checkLifecycle(v2, true, true, false); err != nil {
		t.Errorf("Incorrect version lifecycle (v2): %v", err)
	}
}
