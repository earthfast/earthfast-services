package domain

import (
	"context"

	"armada-node/model"
)

type Resolver interface {
	ProjectForDomain(ctx context.Context, domain string) (model.ID, error)
}

// StaticResolver is a Resolver that has a fixed mapping provided at initialization
// time. It supports a wildcard domain entry ("*") that, if included, will be returned
// for any lookup that does not have an exact match.
type StaticResolver struct {
	data map[string]model.ID
}

func NewStaticResolver(data map[string]model.ID) *StaticResolver {
	return &StaticResolver{data: data}
}

func (r *StaticResolver) ProjectForDomain(_ context.Context, domain string) (model.ID, error) {
	if id, ok := r.data[domain]; ok {
		return id, nil
	}
	if id, ok := r.data["*"]; ok {
		return id, nil
	}
	return model.ZeroID, nil
}
