package site

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"armada-node/model"
)

func TestNotFoundVersion(t *testing.T) {
	ver := notFoundVersion{}

	rw := httptest.NewRecorder()
	ver.ServeHTTP(rw, httptest.NewRequest(http.MethodGet, "/", nil))
	if rw.Code != http.StatusNotFound {
		t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, http.StatusNotFound)
	}
}

func TestNotFoundVersion_IsProject(t *testing.T) {
	ver := notFoundVersion{}

	cases := []struct {
		name    string
		project *model.Project
		wantIs  bool
	}{
		{
			name:    "nil Project",
			project: nil,
			wantIs:  true,
		},
		{
			name:    "Empty Content",
			project: &model.Project{Name: "hello-world"},
			wantIs:  true,
		},
		{
			name: "Valid Project",
			project: &model.Project{
				Name:     "hello-world",
				Content:  "http://example.com/data.tar.gz",
				Checksum: "111",
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
