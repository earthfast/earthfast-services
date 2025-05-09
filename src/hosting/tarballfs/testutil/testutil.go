package testutil

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type TarFile struct {
	FType byte
	Name  string
	Data  string
}

func MakeTarball(t *testing.T, contents []TarFile) *bytes.Buffer {
	t.Helper()

	var tBuf bytes.Buffer
	tw := tar.NewWriter(&tBuf)
	for _, f := range contents {
		hdr := &tar.Header{
			Typeflag: f.FType,
			Name:     f.Name,
			Size:     int64(len(f.Data)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(f.Data)); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	var gzBuf bytes.Buffer
	gzw := gzip.NewWriter(&gzBuf)
	if _, err := io.Copy(gzw, &tBuf); err != nil {
		t.Fatal(err)
	}
	if err := gzw.Close(); err != nil {
		t.Fatal(err)
	}

	return &gzBuf
}

func VerifyOutput(t *testing.T, dir string, contents []TarFile) {
	t.Helper()

	numFiles := 0
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != dir {
			numFiles++
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Error walking %s: %v", dir, err)
	}
	if numFiles != len(contents) {
		t.Errorf("File count mismatch: got %d, want %d", numFiles, len(contents))
	}

	for _, f := range contents {
		fPath := filepath.Join(dir, f.Name)
		switch f.FType {
		case tar.TypeReg:
			data, err := os.ReadFile(fPath)
			if err != nil {
				t.Errorf("Error reading %s: %v", f.Name, err)
				continue
			}
			if got := string(data); got != f.Data {
				t.Errorf("Unexpected file contents for %s: got %s, want %s", f.Name, got, f.Data)
				continue
			}
		case tar.TypeDir:
			_, err := os.Stat(fPath)
			if err != nil {
				t.Errorf("%s not found: %v", f.Name, err)
				continue
			}
		default:
			_, err := os.Stat(fPath)
			if err == nil {
				t.Errorf("Found file %q with unsupported type %x, expected it to be missing", f.Name, f.FType)
			}
		}
	}
}

func StartTarballServer(t *testing.T, contents []TarFile) (*httptest.Server, string) {
	t.Helper()

	tarball := MakeTarball(t, contents)
	checksum := fmt.Sprintf("%x", sha256.Sum256(tarball.Bytes()))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(tarball.Bytes())
	}))

	return ts, checksum
}
