package site

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"armada-node/hosting/fileserver"
	"armada-node/hosting/tarballfs"
	"armada-node/model"

	"go.uber.org/zap"
)

type tarballVersion struct {
	http.Handler
	*tarballfs.FS

	TarballURL string
	Checksum   string
}

func newTarballVersion(logger *zap.Logger, dataDir string, p *model.Project) (*tarballVersion, error) {
	if _, err := url.Parse(p.Content); err != nil {
		return nil, fmt.Errorf("parsing project content as URL: %v", err)
	}
	if p.Checksum == "" {
		return nil, fmt.Errorf("missing checksum for project: %s", p.Name)
	}

	fs, err := tarballfs.New(tarballfs.Options{
		MountDir:  filepath.Join(dataDir, p.Checksum),
		SourceURL: p.Content,
		Checksum:  p.Checksum,
		Logger:    logger,
	})
	if err != nil {
		return nil, fmt.Errorf("creating tarballfs: %v", err)
	}

	return &tarballVersion{
		Handler:    http.FileServer(fileserver.FileOnlyFS(http.FS(fs))),
		FS:         fs,
		TarballURL: p.Content,
		Checksum:   p.Checksum,
	}, nil
}

func (v *tarballVersion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")

	// http.FileServer always 301 redirects index.html, but we just always want
	// to return the actual content, so we'll preempt the redirect by trimming
	// "index.html" off the path if it's present.
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "index.html")

	v.Handler.ServeHTTP(w, r)
}

func (v *tarballVersion) IsProject(p *model.Project) bool {
	if p == nil {
		return false
	}
	return p.Content == v.TarballURL && p.Checksum == v.Checksum
}

func (v *tarballVersion) String() string {
	return fmt.Sprintf("TarballSite[chksum=%s]", v.Checksum)
}
