package hosting

import (
	"armada-node/hosting/site"
	"armada-node/model"

	"go.uber.org/zap"
)

type defaultSiteProvider struct {
	m               model.Client
	versionProvider site.VersionProvider
	logger          *zap.Logger
}

func DefaultSiteProvider(logger *zap.Logger, m model.Client, dataDir string) site.SiteProvider {
	return &defaultSiteProvider{
		m:               m,
		versionProvider: site.DefaultVersionProvider(logger, dataDir),
		logger:          logger,
	}
}

func (sp *defaultSiteProvider) HandlerForProject(id model.ID) site.Site {
	return site.NewHandler(sp.logger, sp.m, sp.versionProvider, id)
}
