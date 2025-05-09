package site

import (
	"archive/tar"
	"net/http"
	"net/http/httptest"
	"testing"

	"armada-node/hosting/tarballfs/testutil"
	"armada-node/model"

	"go.uber.org/zap/zaptest"
)

func TestTarballVersion(t *testing.T) {
	logger := zaptest.NewLogger(t)
	indexDotHTML := testutil.TarFile{
		FType: tar.TypeReg,
		Name:  "hello-world/index.html",
		Data:  "Hello world!",
	}
	contents := []testutil.TarFile{
		{
			FType: tar.TypeDir,
			Name:  "hello-world",
		},
		indexDotHTML,
	}
	ts, checksum := testutil.StartTarballServer(t, contents)
	defer ts.Close()

	p := &model.Project{
		Name:     "hello-world",
		Content:  ts.URL,
		Checksum: checksum,
	}
	ver, err := newTarballVersion(logger, t.TempDir(), p)
	if err != nil {
		t.Fatal(err)
	}
	ver.Start()
	defer ver.Stop()

	rw := httptest.NewRecorder()
	ver.ServeHTTP(rw, httptest.NewRequest(http.MethodGet, "/", nil))
	if rw.Code != http.StatusOK {
		t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, http.StatusOK)
	}
	if got := rw.Body.String(); got != indexDotHTML.Data {
		t.Errorf("Unexpected response content: got %q, want %q", got, indexDotHTML.Data)
	}
}

func TestTarballVersion_IsProject(t *testing.T) {
	logger := zaptest.NewLogger(t)
	p := &model.Project{
		Name:     "hello-world",
		Content:  "http://host.com/data.tar.gz",
		Checksum: "111",
	}
	ver, err := newTarballVersion(logger, t.TempDir(), p)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name    string
		project *model.Project
		wantIs  bool
	}{
		{
			name:    "Identical Project",
			project: p,
			wantIs:  true,
		},
		{
			name:    "nil Project",
			project: nil,
			wantIs:  false,
		},
		{
			name: "Different Content",
			project: &model.Project{
				Name:     p.Name,
				Content:  "http://other-host.com/data.tar.gz",
				Checksum: p.Checksum,
			},
			wantIs: false,
		},
		{
			name: "Different Checksum",
			project: &model.Project{
				Name:     p.Name,
				Content:  p.Content,
				Checksum: "222",
			},
			wantIs: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ver.IsProject(tc.project); got != tc.wantIs {
				t.Errorf("IsProject(%+v) = %t, want %t", tc.project, got, tc.wantIs)
			}
		})
	}
}
