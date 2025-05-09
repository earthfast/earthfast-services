package modeltest

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"armada-node/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomID(t *testing.T) model.ID {
	t.Helper()

	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		t.Fatal(err)
	}
	var id model.ID
	copy(id[:], data[:])
	return id
}

type Client struct {
	projects     []*model.Project
	contentNodes []*model.Node
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) WithContentNodes(nodes ...*model.Node) *Client {
	c.contentNodes = nodes
	return c
}

func (c *Client) WithProjects(projects ...*model.Project) *Client {
	c.projects = projects
	return c
}

func (c *Client) GetNode(ctx context.Context, id model.ID) (*model.Node, error) {
	for _, n := range c.contentNodes {
		if n.ID == id {
			return n, nil
		}
	}
	return nil, nil
}

func (c *Client) GetProject(ctx context.Context, id model.ID) (*model.Project, error) {
	for _, p := range c.projects {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, nil
}

func (c *Client) ContentNodes(ctx context.Context, projectID model.ID) ([]*model.Node, error) {
	var nodes []*model.Node
	for _, n := range c.contentNodes {
		if n.ProjectID == projectID {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}

type RONode struct {
	*model.Node
}

func ReadOnlyNode(n *model.Node) RONode {
	return RONode{Node: n}
}

func (n RONode) ID() model.ID {
	return n.Node.ID
}

func (n RONode) ProjectID() model.ID {
	return n.Node.ProjectID
}

func (n RONode) Host() string {
	return n.Node.Host
}
