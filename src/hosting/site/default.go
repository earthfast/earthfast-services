package site

import (
	"fmt"
	"os"
	"path/filepath"

	"armada-node/model"

	"go.uber.org/zap"
)

type defaultVersionProvider struct {
	dataDir string
	logger  *zap.Logger
}

func DefaultVersionProvider(logger *zap.Logger, dataDir string) VersionProvider {
	return &defaultVersionProvider{
		dataDir: dataDir,
		logger:  logger,
	}
}

func (vp *defaultVersionProvider) VersionForProject(p *model.Project) (Version, error) {
	if p == nil {
		return notFoundVersion{}, nil
	}

	vp.logger.Info("Creating version for project",
		zap.String("projectId", p.ID.Hex()),
		zap.String("name", p.Name))

	// Prepare project directory
	siteDataDir := filepath.Join(vp.dataDir, p.ID.Hex())
	if err := os.MkdirAll(siteDataDir, 0755); err != nil {
		return nil, fmt.Errorf("preparing site directory: %w", err)
	}

	// Check for empty content - maintain original behavior
	if p.Content == "" {
		vp.logger.Info("Creating static site version with empty content",
			zap.String("projectId", p.ID.Hex()),
			zap.String("name", p.Name))
		return notFoundVersion{}, nil
	}

	// Parse metadata to check project type
	metadata, err := p.ParseMetadata()
	if err != nil {
		// If we can't parse metadata, fall back to tarball version
		vp.logger.Warn("Failed to parse project metadata, falling back to tarball",
			zap.Error(err),
			zap.String("projectId", p.ID.Hex()))
	} else if metadata.Type == model.ProjectTypeNextJS {
		// Create Next.js version for Next.js projects
		vp.logger.Info("Creating Next.js version",
			zap.String("projectId", p.ID.Hex()),
			zap.String("name", p.Name))
		return NewNextJSVersion(vp.logger, siteDataDir, p)
	}

	// Default: create tarball version
	vp.logger.Info("Creating static site version",
		zap.String("projectId", p.ID.Hex()),
		zap.String("name", p.Name))
	return newTarballVersion(vp.logger, siteDataDir, p)
}
