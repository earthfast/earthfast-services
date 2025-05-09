package cache

import (
	"context"
	"fmt"
	"time"

	"armada-node/model"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

var (
	DefaultTTL             = 10 * time.Second
	DefaultCleanupInterval = 5 * time.Minute
)

func nodeKey(id model.ID) string {
	return fmt.Sprintf("node.%s", id.Hex())
}

func projectKey(id model.ID) string {
	return fmt.Sprintf("project.%s", id.Hex())
}

func contentNodesKey(projectID model.ID) string {
	return fmt.Sprintf("content-nodes.%s", projectID.Hex())
}

type Options struct {
	TTL             time.Duration
	CleanupInterval time.Duration
}

type Client struct {
	source model.Client
	cache  *gocache.Cache
	dedupe *singleflight.Group
}

func NewClient(source model.Client, opts Options) *Client {
	if opts.TTL == 0 {
		opts.TTL = DefaultTTL
	}
	if opts.CleanupInterval == 0 {
		opts.CleanupInterval = DefaultCleanupInterval
	}
	return &Client{
		source: source,
		cache:  gocache.New(opts.TTL, opts.CleanupInterval),
		dedupe: &singleflight.Group{},
	}
}

func (c *Client) GetNode(ctx context.Context, id model.ID) (*model.Node, error) {
	key := nodeKey(id)
	got, ok := c.cache.Get(key)
	if ok {
		return got.(*model.Node), nil
	}
	got, err, _ := c.dedupe.Do(key, func() (interface{}, error) {
		return c.source.GetNode(ctx, id)
	})
	if err != nil {
		return nil, err
	}
	node := got.(*model.Node)
	if node != nil {
		c.cache.SetDefault(key, node)
	}
	return node, nil
}

func (c *Client) GetProject(ctx context.Context, id model.ID) (*model.Project, error) {
	key := projectKey(id)
	got, ok := c.cache.Get(key)
	if ok {
		return got.(*model.Project), nil
	}
	got, err, _ := c.dedupe.Do(key, func() (interface{}, error) {
		return c.source.GetProject(ctx, id)
	})
	if err != nil {
		return nil, err
	}
	proj := got.(*model.Project)
	if proj != nil {
		c.cache.SetDefault(key, proj)
	}
	return proj, nil
}

func (c *Client) ContentNodes(ctx context.Context, projectID model.ID) ([]*model.Node, error) {
	key := contentNodesKey(projectID)
	got, ok := c.cache.Get(key)
	if ok {
		return got.([]*model.Node), nil
	}
	got, err, _ := c.dedupe.Do(key, func() (interface{}, error) {
		return c.source.ContentNodes(ctx, projectID)
	})
	if err != nil {
		return nil, err
	}
	nodes := got.([]*model.Node)
	if len(nodes) > 0 {
		c.cache.SetDefault(key, nodes)
	}
	return nodes, nil
}
