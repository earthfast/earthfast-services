package site

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"armada-node/model"
	"armada-node/model/modeltest"

	"go.uber.org/zap/zaptest"
)

func TestDefaultVersionProvider(t *testing.T) {
	logger := zaptest.NewLogger(t)
	vp := DefaultVersionProvider(logger, t.TempDir())

	cases := []struct {
		name     string
		project  *model.Project
		wantType interface{}
	}{
		{
			name:     "nil Project",
			project:  nil,
			wantType: reflect.TypeOf(notFoundVersion{}),
		},
		{
			name:     "Empty Content",
			project:  &model.Project{Name: "hello-world"},
			wantType: reflect.TypeOf(notFoundVersion{}),
		},
		{
			name: "Tarball URL in Content",
			project: &model.Project{
				Name:     "hello-world",
				Content:  "http://example.com/data.tar.gz",
				Checksum: "111",
			},
			wantType: reflect.TypeOf(&tarballVersion{}),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ver, err := vp.VersionForProject(tc.project)
			if err != nil {
				t.Fatal(err)
			}
			if gotType := reflect.TypeOf(ver); gotType != tc.wantType {
				t.Errorf("Unexpected site version type: got %s, want %s", gotType, tc.wantType)
			}
		})
	}
}

func TestDefaultVersionProvider_TarballDataDir(t *testing.T) {
	logger := zaptest.NewLogger(t)
	scratch := t.TempDir()

	project := &model.Project{
		ID:       modeltest.RandomID(t),
		Name:     "hello-world",
		Content:  "http://example.com/data.tar.gz",
		Checksum: "111",
	}
	vp := DefaultVersionProvider(logger, scratch)
	if _, err := vp.VersionForProject(project); err != nil {
		t.Fatal(err)
	}

	expectDir := filepath.Join(scratch, project.ID.Hex())
	stat, err := os.Stat(expectDir)
	if err != nil {
		t.Fatalf("Expected %s to exist, got error: %s", expectDir, err)
	}
	if !stat.IsDir() {
		t.Errorf("Expected %s to be a directory, got %s", expectDir, stat)
	}
}
