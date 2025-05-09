package tarballfs

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

func extract(logger *zap.Logger, dst string, r io.Reader, stripComponents int) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("advancing tar reader: %w", err)
		}
		name := filepath.Clean(header.Name)

		// Ensure the file won't escape dst
		if filepath.IsAbs(name) {
			logger.Warn("Skipping file with absolute path",
				zap.String("path", header.Name),
			)
			continue
		}
		if strings.HasPrefix(filepath.Dir(name), "..") {
			logger.Warn("Skipping file with target outside archive",
				zap.String("path", header.Name),
			)
			continue
		}

		// Strip leading path components if requested. Files with too few components get skipped.
		if stripComponents > 0 {
			parts := strings.Split(name, "/")
			if len(parts) <= stripComponents {
				continue
			}
			name = filepath.Join(parts[stripComponents:]...)
		}

		// Construct the final target location within dst
		target := filepath.Join(dst, name)

		switch header.Typeflag {
		case tar.TypeDir:
			// Only create the directory if it doesn't already exist
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return fmt.Errorf("creating directory %s: %w", target, err)
				}
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
			if err != nil {
				return fmt.Errorf("creating file %s: %w", target, err)
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return fmt.Errorf("writing file contents: %w", err)
			}
			f.Close()
		default:
			logger.Warn("Skipping file with unsupported type",
				zap.String("path", header.Name),
				zap.String("type", fmt.Sprintf("%x", header.Typeflag)),
			)
		}
	}
}
