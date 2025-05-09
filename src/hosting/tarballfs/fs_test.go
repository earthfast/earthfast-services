package tarballfs

import (
	"archive/tar"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"armada-node/hosting/tarballfs/testutil"

	"go.uber.org/zap/zaptest"
)

func TestFS(t *testing.T) {
	contents := []testutil.TarFile{
		{
			FType: tar.TypeDir,
			Name:  "example.com",
		},
		{
			FType: tar.TypeReg,
			Name:  "example.com/foo.txt",
			Data:  "Hello world!",
		},
	}
	ts, checksum := testutil.StartTarballServer(t, contents)
	defer ts.Close()

	scratch := t.TempDir()
	mountDir := filepath.Join(scratch, "example.com")
	fs, err := New(Options{
		MountDir:  mountDir,
		SourceURL: ts.URL,
		Checksum:  checksum,
		Logger:    zaptest.NewLogger(t),
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := fs.load(context.Background()); err != nil {
		t.Fatal(err)
	}

	wantContents := []testutil.TarFile{{
		FType: tar.TypeReg,
		Name:  "foo.txt",
		Data:  "Hello world!",
	}}
	testutil.VerifyOutput(t, mountDir, wantContents)
}

func TestFS_AlreadyExistsNoop(t *testing.T) {
	scratch := t.TempDir()
	fs, err := New(Options{
		MountDir:  scratch,
		SourceURL: "http://localhost/dummy",
		Checksum:  "0",
		Logger:    zaptest.NewLogger(t),
	})
	if err != nil {
		t.Fatal(err)
	}
	// The mount directory already exists so load() should just be a noop. If
	// it attempts to load the tarball it will fail since we gave it a dummy URL
	// and checksum, so we simply look for a non-error response as noop proof.
	if err := fs.load(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestFS_BadHTTPStatus(t *testing.T) {
	for _, code := range []int{301, 302, 404, 500} {
		t.Run(fmt.Sprintf("HTTP %d", code), func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, http.StatusText(code), code)
			}))
			defer ts.Close()

			fs, err := New(Options{
				MountDir:  filepath.Join(t.TempDir(), "example.com"),
				SourceURL: ts.URL,
				Checksum:  "0",
				Logger:    zaptest.NewLogger(t),
			})
			if err != nil {
				t.Fatal(err)
			}
			if err := fs.load(context.Background()); err == nil {
				t.Error("load() succeeded, expected error")
			}
		})
	}
}

func TestFS_IncorrectChecksum(t *testing.T) {
	contents := []testutil.TarFile{{
		FType: tar.TypeReg,
		Name:  "foo.txt",
		Data:  "Hello world!",
	}}
	ts, _ := testutil.StartTarballServer(t, contents)
	defer ts.Close()

	fs, err := New(Options{
		MountDir:  filepath.Join(t.TempDir(), "example.com"),
		SourceURL: ts.URL,
		Checksum:  "0",
		Logger:    zaptest.NewLogger(t),
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := fs.load(context.Background()); err == nil {
		t.Error("load() succeeded, expected error")
	}
}

func TestFS_Cancels(t *testing.T) {
	waitCh := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-waitCh
		fmt.Fprint(w, "ok")
	}))
	defer ts.Close()
	defer close(waitCh)

	fs, err := New(Options{
		MountDir:  filepath.Join(t.TempDir(), "example.com"),
		SourceURL: ts.URL,
		Checksum:  "0",
		Logger:    zaptest.NewLogger(t),
	})
	if err != nil {
		t.Fatal(err)
	}
	fs.Start()

	stoppedCh := make(chan struct{})
	go func() {
		fs.Stop()
		close(stoppedCh)
	}()

	select {
	case <-stoppedCh:
	case <-time.After(5 * time.Second):
		t.Fatal("FS did not shutdown quickly when stopped")
	}
}
