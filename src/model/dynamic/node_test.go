package dynamic

import (
	"context"
	"testing"

	"armada-node/model"
	"armada-node/model/modeltest"

	"go.uber.org/zap/zaptest"
)

func TestNode_Bootstrap(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node0 := &model.Node{
		ID:        modeltest.RandomID(t),
		Host:      "node0.armadanetwork.com",
		ProjectID: project0.ID,
	}
	m := modeltest.NewClient().WithContentNodes(node0)

	node := NewNode(logger, m, node0.ID, Options{})
	if err := node.Bootstrap(ctx); err != nil {
		t.Fatal(err)
	}

	if node.ID() != node0.ID {
		t.Errorf("Unexpected ID: got %s, want %s", node.ID(), node0.ID)
	}
	if node.Host() != node0.Host {
		t.Errorf("Unexpected Host: got %s, want %s", node.Host(), node0.Host)
	}
	if node.ProjectID() != node0.ProjectID {
		t.Errorf("Unexpected ProjectID: got %s, want %s", node.ProjectID(), node0.ProjectID)
	}
}

func TestNode_Refresh(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	project1 := &model.Project{ID: modeltest.RandomID(t), Name: "goodbye-world"}
	node0_v1 := &model.Node{
		ID:        modeltest.RandomID(t),
		Host:      "node0.armadanetwork.com",
		ProjectID: project0.ID,
	}
	m := modeltest.NewClient().WithContentNodes(node0_v1)

	node := NewNode(logger, m, node0_v1.ID, Options{})
	if err := node.Bootstrap(ctx); err != nil {
		t.Fatal(err)
	}
	if node.ProjectID() != node0_v1.ProjectID {
		t.Errorf("Unexpected ProjectID: got %s, want %s", node.ProjectID(), node0_v1.ProjectID)
	}

	// Update node0's ProjectID in the model
	node0_v2 := &model.Node{
		ID:        node0_v1.ID,
		Host:      node0_v1.Host,
		ProjectID: project1.ID,
	}
	m.WithContentNodes(node0_v2)

	if err := node.refresh(ctx); err != nil {
		t.Fatal(err)
	}
	if node.ProjectID() != node0_v2.ProjectID {
		t.Errorf("Unexpected ProjectID: got %s, want %s", node.ProjectID(), node0_v2.ProjectID)
	}
}
