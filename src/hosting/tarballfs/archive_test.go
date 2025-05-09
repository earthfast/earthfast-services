package tarballfs

import (
	"archive/tar"
	"testing"

	"armada-node/hosting/tarballfs/testutil"

	"go.uber.org/zap/zaptest"
)

func TestExtract(t *testing.T) {
	logger := zaptest.NewLogger(t)
	contents := []testutil.TarFile{
		{
			FType: tar.TypeReg,
			Name:  "foo.txt",
			Data:  "Hello world!",
		},
		{
			FType: tar.TypeReg,
			Name:  "bar.txt",
			Data:  "Hello bar!",
		},
		{
			FType: tar.TypeDir,
			Name:  "a",
		},
		{
			FType: tar.TypeReg,
			Name:  "a/b",
			Data:  "Hi from a/b",
		},
	}
	tarball := testutil.MakeTarball(t, contents)

	scratch := t.TempDir()
	if err := extract(logger, scratch, tarball, 0); err != nil {
		t.Fatal(err)
	}
	testutil.VerifyOutput(t, scratch, contents)
}

func TestExtract_SkippedFiles(t *testing.T) {
	logger := zaptest.NewLogger(t)
	contents := []testutil.TarFile{
		{
			FType: tar.TypeReg,
			Name:  "/foo.txt",
		},
		{
			FType: tar.TypeReg,
			Name:  "a/../../bar.txt",
		},
		{
			FType: tar.TypeSymlink,
			Name:  "a/symlink",
		},
	}
	tarball := testutil.MakeTarball(t, contents)

	scratch := t.TempDir()
	if err := extract(logger, scratch, tarball, 0); err != nil {
		t.Fatal(err)
	}
	testutil.VerifyOutput(t, scratch, nil)
}

func TestExtract_StripComponents(t *testing.T) {
	logger := zaptest.NewLogger(t)
	contents := []testutil.TarFile{
		{
			FType: tar.TypeReg,
			Name:  "ignore_me.txt",
			Data:  "Too few components",
		},
		{
			FType: tar.TypeDir,
			Name:  "example.com",
		},
		{
			FType: tar.TypeReg,
			Name:  "example.com/foo.txt",
			Data:  "Hello world!",
		},
		{
			FType: tar.TypeDir,
			Name:  "example.com/a",
		},
		{
			FType: tar.TypeReg,
			Name:  "example.com/a/b",
			Data:  "Hi from example.com/a/b",
		},
	}
	tarball := testutil.MakeTarball(t, contents)

	scratch := t.TempDir()
	if err := extract(logger, scratch, tarball, 1); err != nil {
		t.Fatal(err)
	}

	wantContents := []testutil.TarFile{
		{
			FType: tar.TypeReg,
			Name:  "foo.txt",
			Data:  "Hello world!",
		},
		{
			FType: tar.TypeDir,
			Name:  "a",
		},
		{
			FType: tar.TypeReg,
			Name:  "a/b",
			Data:  "Hi from example.com/a/b",
		},
	}
	testutil.VerifyOutput(t, scratch, wantContents)
}
