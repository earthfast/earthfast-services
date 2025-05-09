package uptime

import (
	"context"
	"net/url"
	"sort"
	"testing"
	"time"

	"armada-node/model"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

type fakeClock struct {
	now time.Time
}

func newFakeClock(now time.Time) *fakeClock {
	return &fakeClock{
		now: now,
	}
}

func (c *fakeClock) Now() time.Time {
	return c.now
}

var sortStrings = cmp.Transformer("SortStrings", func(in []string) []string {
	out := append([]string{}, in...)
	sort.Strings(out)
	return out
})

func TestMeter(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	cleanup, probeURL, wantChecksum := newNodeServer(t, "Hello, world!")
	defer cleanup()

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node0 := &model.Node{Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node1 := &model.Node{Host: probeURL.Host, ProjectID: project0.ID}
	m := modeltest.NewClient().WithContentNodes(node0, node1)

	meter, err := NewMeter(
		MeterArgs{
			Model:  m,
			Self:   modeltest.ReadOnlyNode(node0),
			Store:  NewInMemoryStore(logger, 100),
			Logger: logger,
		},
		MeterOptions{},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer meter.Stop()

	// Setup probers.
	if err := meter.manageProbers(ctx); err != nil {
		t.Fatal(err)
	}

	// Probe request for an unknown node gets dropped.
	badURL := &url.URL{Host: "foo.com", Path: "index.html"}
	if ok := meter.Probe(badURL, "checksum", nil); ok {
		t.Errorf("Probe request for unknown node was accepted, expected it to be dropped: url=%s", badURL)
	}

	// Probe request for an unknown node gets enqueued
	doneCh := make(chan bool)
	if ok := meter.Probe(probeURL, wantChecksum, doneCh); !ok {
		t.Fatalf("Probe request was dropped, expected it to be accepted: url=%s", probeURL)
	}
	select {
	case ok := <-doneCh:
		if !ok {
			t.Error("Expected probe to succeed, got failed")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Timed out waiting for probe result")
	}
}

func TestMeter_manageProbers(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node0 := &model.Node{Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node1 := &model.Node{Host: "node1.armadanetwork.com", ProjectID: project0.ID}
	node2 := &model.Node{Host: "node2.armadanetwork.com", ProjectID: project0.ID}

	cases := []struct {
		name        string
		nodes       []*model.Node
		wantProbers []string
	}{
		{
			name: "No nodes",
		},
		{
			name:  "Self only",
			nodes: []*model.Node{node0},
		},
		{
			name:        "Peers only",
			nodes:       []*model.Node{node1, node2},
			wantProbers: []string{node1.Host, node2.Host},
		},
		{
			name:        "Self and peers",
			nodes:       []*model.Node{node0, node1, node2},
			wantProbers: []string{node1.Host, node2.Host},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := modeltest.NewClient().WithContentNodes(tc.nodes...)
			meter, err := NewMeter(
				MeterArgs{
					Model:  m,
					Self:   modeltest.ReadOnlyNode(node0),
					Store:  NewInMemoryStore(logger, 100),
					Logger: logger,
				},
				MeterOptions{},
			)
			if err != nil {
				t.Fatal(err)
			}
			defer meter.Stop()

			if err := meter.manageProbers(ctx); err != nil {
				t.Fatal(err)
			}

			var gotProbers []string
			for host := range meter.probers {
				gotProbers = append(gotProbers, host)
			}
			if diff := cmp.Diff(tc.wantProbers, gotProbers, sortStrings); diff != "" {
				t.Errorf("Prober mismatch (-want +got): %s", diff)
			}
		})
	}
}

func TestMeter_manageIntervals(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	intervalDuration := time.Minute
	clock := newFakeClock(time.Now().Truncate(intervalDuration))
	startInterval := clock.now.Unix() / int64(intervalDuration.Seconds())

	node0 := &model.Node{Host: "node0.armadanetwork.com"}

	meter, err := NewMeter(
		MeterArgs{
			Model:  modeltest.NewClient(),
			Self:   modeltest.ReadOnlyNode(node0),
			Store:  NewInMemoryStore(logger, 100),
			Logger: logger,
		},
		MeterOptions{
			IntervalDuration: intervalDuration,
			clock:            clock,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer meter.Stop()

	cases := []struct {
		name         string
		timeDelta    time.Duration
		wantInterval int64
	}{
		{
			name:         "Now",
			wantInterval: startInterval,
		},
		{
			name:         "Same interval",
			timeDelta:    intervalDuration / 2,
			wantInterval: startInterval,
		},
		{
			name:         "Next interval",
			timeDelta:    intervalDuration,
			wantInterval: startInterval + 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clock.now = clock.now.Add(tc.timeDelta)
			if err := meter.manageInterval(ctx); err != nil {
				t.Fatal(err)
			}
			if gotInterval := meter.interval; gotInterval != tc.wantInterval {
				t.Errorf("Incorrect interval: got %d, want %d", gotInterval, tc.wantInterval)
			}
		})
	}
}

func TestMeter_persistCurrentInterval(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	node0 := &model.Node{Host: "node0.armadanetwork.com", ProjectID: project0.ID}
	node1 := &model.Node{Host: "node1.armadanetwork.com", ProjectID: project0.ID}
	m := modeltest.NewClient().WithContentNodes(node0, node1)

	clock := newFakeClock(time.Now())
	intervalDuration := time.Minute
	startInterval := clock.now.Unix() / int64(intervalDuration.Seconds())

	store := NewInMemoryStore(logger, 100)
	meter, err := NewMeter(
		MeterArgs{
			Model:  m,
			Self:   modeltest.ReadOnlyNode(node0),
			Store:  store,
			Logger: logger,
		},
		MeterOptions{
			IntervalDuration: intervalDuration,
			clock:            clock,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer meter.Stop()

	// Start an interval.
	if err := meter.manageInterval(ctx); err != nil {
		t.Fatal(err)
	}

	// Setup probers.
	if err := meter.manageProbers(ctx); err != nil {
		t.Fatal(err)
	}

	// Manually register some probe results.
	prb, ok := meter.probers[node1.Host]
	if !ok {
		t.Fatalf("Missing prober for host: %s", node1.Host)
	}
	prb.successCount = 4
	prb.failureCount = 1

	// Add some requests.
	meter.IncrementRequestCount(42)

	// Force an interval transition so persisting occurs.
	clock.now = clock.now.Add(intervalDuration)
	if err := meter.manageInterval(ctx); err != nil {
		t.Fatal(err)
	}

	wantData := []IntervalData{
		{
			StartTime:    time.Unix(startInterval*int64(intervalDuration.Seconds()), 0),
			EndTime:      time.Unix((startInterval+1)*int64(intervalDuration.Seconds()), 0),
			RequestCount: 42,
			ProbeResults: map[string]ProbeCounts{
				node1.Host: {
					Success: 4,
					Failure: 1,
				},
			},
		},
	}
	if diff := cmp.Diff(wantData, store.data); diff != "" {
		t.Errorf("IntervalData mismatch (-want +got):\n%s", diff)
	}

	// Force another interval transition so an *empty* interval gets persisted.
	clock.now = clock.now.Add(intervalDuration)
	if err := meter.manageInterval(ctx); err != nil {
		t.Fatal(err)
	}
	if n := len(store.data); n > 1 {
		t.Error("An empty interval was persisted, expected it to be dropped")
	}
}
