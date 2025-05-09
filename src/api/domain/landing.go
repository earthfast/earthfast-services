package domain

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

type landingPageHandler struct {
	rootDir http.Dir
	logger  *zap.Logger
}

func newLandingPageHandler(logger *zap.Logger, root http.Dir) (*landingPageHandler, error) {
	// Require that index.html be present in the root directory.
	index, err := root.Open("index.html")
	if err != nil {
		return nil, fmt.Errorf("missing index.html: %v", err)
	}
	index.Close()

	return &landingPageHandler{
		rootDir: root,
		logger:  logger,
	}, nil
}

func (h *landingPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Look for a file (not a directory) with the exact name as the request path,
	// serving it directly if found.
	if f, err := h.findFile(r.URL.Path); err != nil {
		h.logger.Error("Failed to find file", zap.String("path", r.URL.Path), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else if f != nil {
		defer f.Close()
		http.ServeContent(w, r, r.URL.Path, time.Time{}, f)
		return
	}

	// There isn't an exact file match, so now we'll serve index.html as long as we
	// think the client will accept an HTML response.
	if !acceptsHTML(r) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filepath.Join(string(h.rootDir), "index.html"))
}

func (h *landingPageHandler) findFile(path string) (http.File, error) {
	f, err := h.rootDir.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("opening file: %v", err)
	}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("getting file info: %v", err)
	}
	if stat.IsDir() {
		return nil, nil
	}

	return f, nil
}

// acceptsHTML returns true if the request appears to support a text/html MIME type response.
func acceptsHTML(r *http.Request) bool {
	ext := path.Ext(r.URL.Path)

	// Yes if the path contains no file extension.
	if ext == "" {
		return true
	}

	// Yes if the extension indicates an HTML file.
	if ext == ".html" || ext == ".htm" {
		return true
	}

	// Yes if the Accept header contains text/html explicitly.
	if accept := strings.TrimSpace(r.Header.Get("Accept")); accept != "" {
		for _, entry := range strings.Split(accept, ",") {
			if strings.HasPrefix(strings.TrimSpace(entry), "text/html") {
				return true
			}
		}
	}

	return false
}
