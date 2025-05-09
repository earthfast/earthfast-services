package cache

import (
	"context"
	"testing"

	"armada-node/model"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetNode(t *testing.T) {
	ctx := context.Background()

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	project1 := &model.Project{ID: modeltest.RandomID(t), Name: "goodbye-world"}
	node0_v1 := &model.Node{ID: modeltest.RandomID(t), Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node0_v2 := &model.Node{ID: node0_v1.ID, Host: node0_v1.Host, ProjectID: project1.ID}
	m := modeltest.NewClient().WithContentNodes(node0_v1)

	// Passing ttl=0 instructs the cache to never auto-evict
	client := NewClient(m, Options{TTL: 0})

	// Node not found
	got, err := client.GetNode(ctx, modeltest.RandomID(t))
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("Unexpected node: got %+v, want nil", got)
	}

	// Ok, load from source
	got, err = client.GetNode(ctx, node0_v1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(node0_v1, got); diff != "" {
		t.Errorf("GetNode mismatch (-want +got):\n%s", diff)
	}

	// Change node0 in the model
	m.WithContentNodes(node0_v2)

	// Ok, load from cache
	got, err = client.GetNode(ctx, node0_v2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(node0_v1, got); diff != "" {
		t.Errorf("GetNode mismatch (-want +got):\n%s", diff)
	}

	// Force eviction
	client.cache.Delete(nodeKey(node0_v1.ID))

	// Ok, load from source
	got, err = client.GetNode(ctx, node0_v2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(node0_v2, got); diff != "" {
		t.Errorf("GetNode mismatch (-want +got):\n%s", diff)
	}
}

func TestClient_GetProject(t *testing.T) {
	ctx := context.Background()

	project0_v1 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	project0_v2 := &model.Project{ID: project0_v1.ID, Name: "goodbye-world"}
	m := modeltest.NewClient().WithProjects(project0_v1)

	// Passing ttl=0 instructs the cache to never auto-evict
	client := NewClient(m, Options{TTL: 0})

	// Project not found
	got, err := client.GetProject(ctx, modeltest.RandomID(t))
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("Unexpected project: got %+v, want nil", got)
	}

	// Ok, load from source
	got, err = client.GetProject(ctx, project0_v1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(project0_v1, got); diff != "" {
		t.Errorf("GetProject mismatch (-want +got):\n%s", diff)
	}

	// Change project0 in the model
	m.WithProjects(project0_v2)

	// Ok, load from cache
	got, err = client.GetProject(ctx, project0_v2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(project0_v1, got); diff != "" {
		t.Errorf("GetProject mismatch (-want +got):\n%s", diff)
	}

	// Force eviction
	client.cache.Delete(projectKey(project0_v1.ID))

	// Ok, load from source
	got, err = client.GetProject(ctx, project0_v2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(project0_v2, got); diff != "" {
		t.Errorf("GetProject mismatch (-want +got):\n%s", diff)
	}
}

func TestClient_ContentNodes(t *testing.T) {
	ctx := context.Background()

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node0 := &model.Node{ID: modeltest.RandomID(t), Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node1 := &model.Node{ID: modeltest.RandomID(t), Host: "node1.armadanetwork.com", ProjectID: project0.ID}
	node2 := &model.Node{ID: modeltest.RandomID(t), Host: "node2.armadanetwork.com", ProjectID: project0.ID}
	m := modeltest.NewClient()

	// Passing ttl=0 instructs the cache to never auto-evict
	client := NewClient(m, Options{TTL: 0})

	// No content nodes for project (doesn't cache)
	got, err := client.ContentNodes(ctx, project0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(got); n > 0 {
		t.Errorf("Unexpected ContentNode count: got %d, want %d", n, 0)
	}

	// Populate content nodes in model
	m.WithContentNodes(node0, node1)

	// Ok, load from source & cache
	got, err = client.ContentNodes(ctx, project0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff([]*model.Node{node0, node1}, got); diff != "" {
		t.Errorf("ContentNodes mismatch (-want +got):\n%s", diff)
	}

	// Change node0 in the model
	m.WithContentNodes(node2)

	// Ok, load from cache
	got, err = client.ContentNodes(ctx, project0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff([]*model.Node{node0, node1}, got); diff != "" {
		t.Errorf("ContentNodes mismatch (-want +got):\n%s", diff)
	}

	// Force eviction
	client.cache.Delete(contentNodesKey(project0.ID))

	// Ok, load from source
	got, err = client.ContentNodes(ctx, project0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff([]*model.Node{node2}, got); diff != "" {
		t.Errorf("ContentNodes mismatch (-want +got):\n%s", diff)
	}
}
