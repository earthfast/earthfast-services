package fileserver

import (
	"fmt"
	"net/http"
	"os"
)

type fileOnlyFileSystem struct {
	fs http.FileSystem
}

// FileOnlyFS returns an http.FileSystem that won't list directories.
// Suggested by: https://groups.google.com/g/golang-nuts/c/bStLPdIVM6w/m/hidTJgDZpHcJ
func FileOnlyFS(fs http.FileSystem) http.FileSystem {
	return fileOnlyFileSystem{fs}
}

func (fs fileOnlyFileSystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", name, err)
	}
	return fileOnlyFile{f}, nil
}

type fileOnlyFile struct {
	http.File
}

func (f fileOnlyFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
