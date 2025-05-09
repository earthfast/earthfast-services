package domain

import (
	"context"
	"sync"
	"time"

	"armada-node/model"

	"go.uber.org/zap"
)

type DynamicResolver struct {
	data    map[string]model.ID
	mu      sync.RWMutex // Add a mutex to protect the data map
	dataUrl string
	logger  *zap.Logger
}

func NewDynamicResolver(dataUrl string, logger *zap.Logger) *DynamicResolver {
	if logger == nil {
		logger = zap.NewNop()
	}

	data, err := DownloadAndParseJSON(dataUrl)
	if err != nil {
		logger.Error("Error downloading and parsing JSON",
			zap.String("url", dataUrl),
			zap.Error(err))
		return nil
	}

	resolver := &DynamicResolver{
		dataUrl: dataUrl,
		data:    data,
		logger:  logger,
	}

	go resolver.startDataRefresh()

	return resolver
}

func (r *DynamicResolver) startDataRefresh() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			updatedData, err := DownloadAndParseJSON(r.dataUrl)
			if err != nil {
				// Continue so future updates can still happen
				r.logger.Error("Error refreshing domain mapping data",
					zap.String("url", r.dataUrl),
					zap.Error(err))
				continue
			} else {
				r.logger.Info("Updated domain to project mapping data",
					zap.String("url", r.dataUrl),
					zap.Int("mappings", len(updatedData)))
				r.updateData(updatedData)
			}
		}
	}
}

func (r *DynamicResolver) updateData(newData map[string]model.ID) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data = newData
}

func (r *DynamicResolver) ProjectForDomain(_ context.Context, domain string) (model.ID, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if id, ok := r.data[domain]; ok {
		return id, nil
	}
	if id, ok := r.data["*"]; ok {
		return id, nil
	}
	return model.ZeroID, nil
}
