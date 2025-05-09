package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"armada-node/api"
	"armada-node/model"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
)

func TestContentNodeProvider(t *testing.T) {
	ctx := context.Background()

	project0 := &model.Project{ID: model.ID{0}, Name: "hello-world"}
	node0 := &model.Node{Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node1 := &model.Node{Host: "node1.armadanetwork.com", ProjectID: project0.ID}
	m := modeltest.NewClient().WithContentNodes(node0, node1)

	np := NewContentNodeProvider(m, http.DefaultClient)
	nodes, err := np.AllNodes(ctx, project0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 2 {
		t.Errorf("Unexpected node count: got %d, want %d", len(nodes), 2)
	}
}

func TestContentNode(t *testing.T) {
	ctx := context.Background()

	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	fakeResponse := &api.UptimeResponse{
		StartTime:    1,
		EndTime:      2,
		RequestCount: 3,
		ProbeResults: map[string]api.UptimeProbeCounts{
			"node1.armadanetwork.com": {Failure: 4},
		},
	}

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantStart := fmt.Sprintf("%d", startTime.Unix())
		if start := r.URL.Query().Get("start"); start != wantStart {
			t.Errorf("Unexpected start query parameter: got %s, want %s", start, wantStart)
		}
		wantEnd := fmt.Sprintf("%d", endTime.Unix())
		if end := r.URL.Query().Get("end"); end != wantEnd {
			t.Errorf("Unexpected end query parameter: got %s, want %s", end, wantEnd)
		}
		if err := json.NewEncoder(w).Encode(fakeResponse); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	tsURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	node0 := &model.Node{Host: tsURL.Host}
	node := newContentNode(node0, ts.Client())
	if node.Host() != node0.Host {
		t.Errorf("Unexpected Host() response: got %s, want %s", node.Host(), node0.Host)
	}

	resCh := make(chan fetchResult)
	go node.Uptime(ctx, resCh, startTime, endTime)

	var got fetchResult
	select {
	case got = <-resCh:
	case <-time.After(5 * time.Second):
		t.Fatal("Timeout while waiting for Uptime() result")
	}

	if got.err != nil {
		t.Fatal(got.err)
	}
	if diff := cmp.Diff(fakeResponse, got.resp); diff != "" {
		t.Errorf("UptimeResponse mismatch (-want +got):\n%s", diff)
	}
}
