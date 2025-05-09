package uptime

import (
	"context"
	"fmt"
	"testing"
	"time"

	"armada-node/api"
	"armada-node/model"

	"github.com/google/go-cmp/cmp"
)

type staticNodeProvider struct {
	nodes []Node
}

func newStaticNodeProvider(nodes ...Node) *staticNodeProvider {
	return &staticNodeProvider{
		nodes: nodes,
	}
}

func (snp *staticNodeProvider) AllNodes(ctx context.Context, _ model.ID) ([]Node, error) {
	return snp.nodes, nil
}

type fakeNode struct {
	host string
	resp *api.UptimeResponse
	err  error
}

func (n *fakeNode) Host() string {
	return n.host
}

func (n *fakeNode) Uptime(ctx context.Context, resCh chan<- fetchResult, startTime, endTime time.Time) {
	select {
	case <-ctx.Done():
		resCh <- fetchResult{err: ctx.Err()}
	case resCh <- fetchResult{resp: n.resp, err: n.err}:
	}
}

func TestReporter(t *testing.T) {
	makeNodes := func(start, end time.Time, requestCounts []int, failureCounts []map[string]int) []Node {
		var nodes []Node
		for i := 0; i < len(requestCounts); i++ {
			n := &fakeNode{
				host: fmt.Sprintf("node%d", i),
				resp: &api.UptimeResponse{
					StartTime:    int(start.Unix()),
					EndTime:      int(end.Unix()),
					RequestCount: requestCounts[i],
					ProbeResults: make(map[string]api.UptimeProbeCounts),
				},
			}
			for host, count := range failureCounts[i] {
				n.resp.ProbeResults[host] = api.UptimeProbeCounts{Failure: count}
			}
			nodes = append(nodes, n)
		}
		return nodes
	}

	cases := []struct {
		name      string
		projectID model.ID

		start         time.Time
		end           time.Time
		requestCounts []int
		failureCounts []map[string]int

		requestCountPercentile       int
		probeFailureCountPercentile  int
		nominalUptimeRatioPercentile int

		want UptimeReport
	}{
		{
			name:      "No requests",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{0, 0, 0},
			failureCounts: []map[string]int{
				map[string]int{},
				map[string]int{},
				map[string]int{},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 1.0,
					"node1": 1.0,
					"node2": 1.0,
				},
			},
		},
		{
			name:      "No downtime",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{1000, 1000, 1000},
			failureCounts: []map[string]int{
				map[string]int{},
				map[string]int{},
				map[string]int{},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 1.0,
					"node1": 1.0,
					"node2": 1.0,
				},
			},
		},
		{
			name:      "one node down 50% of the time",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{1000, 1000, 1000},
			failureCounts: []map[string]int{
				map[string]int{},
				map[string]int{
					"node0": 250,
				},
				map[string]int{
					"node0": 250,
				},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 0.5,
					"node1": 1.0,
					"node2": 1.0,
				},
			},
		},
		{
			name:      "majority nodes down 50% of the time",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{1000, 1000, 1000},
			failureCounts: []map[string]int{
				map[string]int{
					"node2": 250,
				},
				map[string]int{
					"node0": 250,
					"node2": 250,
				},
				map[string]int{
					"node0": 250,
				},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 1.0,
					"node1": 1.0,
					"node2": 1.0,
				},
			},
		},
		{
			name:      "outlier request counts are ignored",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{1000000, 1000, 1000},
			failureCounts: []map[string]int{
				map[string]int{},
				map[string]int{
					"node0": 250,
				},
				map[string]int{
					"node0": 250,
				},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 0.5,
					"node1": 1.0,
					"node2": 1.0,
				},
			},
		},
		{
			name:      "outlier failure counts are ignored",
			projectID: model.ID{0},

			start:         time.Unix(0, 0),
			end:           time.Unix(10, 0),
			requestCounts: []int{1000, 1000, 1000, 1000, 1000},
			failureCounts: []map[string]int{
				map[string]int{},
				map[string]int{
					"node0": 125,
				},
				map[string]int{
					"node0": 1000,
				},
				map[string]int{
					"node0": 125,
				},
				map[string]int{
					"node0": 125,
				},
			},

			requestCountPercentile:       50,
			probeFailureCountPercentile:  50,
			nominalUptimeRatioPercentile: 50,

			want: UptimeReport{
				StartTime: time.Unix(0, 0),
				EndTime:   time.Unix(10, 0),
				UptimeRatios: map[string]float64{
					"node0": 0.5,
					"node1": 1.0,
					"node2": 1.0,
					"node3": 1.0,
					"node4": 1.0,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			nodes := makeNodes(tc.start, tc.end, tc.requestCounts, tc.failureCounts)

			np := newStaticNodeProvider(nodes...)
			opts := ReportOptions{
				StartTime:                    tc.start,
				EndTime:                      tc.end,
				RequestCountPercentile:       tc.requestCountPercentile,
				ProbeFailureCountPercentile:  tc.probeFailureCountPercentile,
				NominalUptimeRatioPercentile: tc.nominalUptimeRatioPercentile,
			}
			reporter, err := NewReporter(np, opts)
			if err != nil {
				t.Fatal(err)
			}

			got, err := reporter.Run(ctx, tc.projectID)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("UptimeReport mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
