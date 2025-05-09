package uptime

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type InMemoryStore struct {
	data       []IntervalData
	maxEntries int
	logger     *zap.Logger
}

func NewInMemoryStore(logger *zap.Logger, maxEntries uint) *InMemoryStore {
	return &InMemoryStore{
		maxEntries: int(maxEntries),
		logger:     logger,
	}
}

func (s *InMemoryStore) Put(_ context.Context, data IntervalData) error {
	s.data = append(s.data, data)
	if len(s.data) > s.maxEntries {
		s.data = s.data[1:]
	}

	s.logger.Info("Uptime interval finished",
		zap.Time("startTime", data.StartTime),
		zap.Time("endTime", data.EndTime),
		zap.Any("probeResults", data.ProbeResults),
	)
	return nil
}

func (s *InMemoryStore) Range(_ context.Context, start, end time.Time, fn func(IntervalData) error) error {
	for _, data := range s.data {
		if data.StartTime.Before(start) || data.EndTime.After(end) {
			continue
		}
		if err := fn(data); err != nil {
			return err
		}
	}
	return nil
}
