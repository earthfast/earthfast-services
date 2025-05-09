package content

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
	"armada-node/metering/uptime"
	"armada-node/model/modeltest"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

func TestHandler_Uptime(t *testing.T) {
	int0 := uptime.IntervalData{
		StartTime:    time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		EndTime:      time.Date(2000, time.January, 1, 0, 10, 0, 0, time.UTC),
		RequestCount: 42,
		ProbeResults: map[string]uptime.ProbeCounts{
			"node0": {Success: 1, Failure: 2},
		},
	}
	int1 := uptime.IntervalData{
		StartTime:    int0.EndTime,
		EndTime:      int0.EndTime.Add(10 * time.Minute),
		RequestCount: 7,
		ProbeResults: map[string]uptime.ProbeCounts{
			"node0": {Success: 3, Failure: 4},
			"node1": {Success: 5, Failure: 6},
		},
	}

	cases := []struct {
		name         string
		query        url.Values
		intervalData []uptime.IntervalData
		wantCode     int
		wantResponse *api.UptimeResponse
	}{
		{
			name: "Missing start",
			query: url.Values{
				"end": {"12345"},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Missing end",
			query: url.Values{
				"start": {"12345"},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Malformed start",
			query: url.Values{
				"start": {"hi"},
				"end":   {"12345"},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Malformed end",
			query: url.Values{
				"start": {"hi"},
				"end":   {"12345"},
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "No matching intervals",
			query: url.Values{
				"start": {"0"},
				"end":   {"1"},
			},
			intervalData: []uptime.IntervalData{int0, int1},
			wantCode:     http.StatusOK,
			wantResponse: &api.UptimeResponse{
				StartTime:    0,
				EndTime:      1,
				ProbeResults: map[string]api.UptimeProbeCounts{},
			},
		},
		{
			name: "Matching intervals",
			query: url.Values{
				"start": {fmt.Sprintf("%d", int0.StartTime.Unix())},
				"end":   {fmt.Sprintf("%d", int1.EndTime.Unix())},
			},
			intervalData: []uptime.IntervalData{int0, int1},
			wantCode:     http.StatusOK,
			wantResponse: &api.UptimeResponse{
				StartTime:    int(int0.StartTime.Unix()),
				EndTime:      int(int1.EndTime.Unix()),
				RequestCount: int(int0.RequestCount + int1.RequestCount),
				ProbeResults: map[string]api.UptimeProbeCounts{
					"node0": {
						Success: int(int0.ProbeResults["node0"].Success + int1.ProbeResults["node0"].Success),
						Failure: int(int0.ProbeResults["node0"].Failure + int1.ProbeResults["node0"].Failure),
					},
					"node1": {
						Success: int(int1.ProbeResults["node1"].Success),
						Failure: int(int1.ProbeResults["node1"].Failure),
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			logger := zaptest.NewLogger(t)
			m := modeltest.NewClient()

			uptime := newFakeUptimeMeter(logger)
			for _, data := range tc.intervalData {
				if err := uptime.store.Put(ctx, data); err != nil {
					t.Fatal(err)
				}
			}

			h := NewHandler(logger, m, uptime, nil, nil)
			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/v1/uptime", nil)
			r.URL.RawQuery = tc.query.Encode()
			h.ServeHTTP(rw, r)
			if rw.Code != tc.wantCode {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, tc.wantCode)
			}
			if tc.wantResponse != nil {
				var got api.UptimeResponse
				if err := json.NewDecoder(rw.Body).Decode(&got); err != nil {
					t.Fatalf("Error decoding body: %v", err)
				}
				if diff := cmp.Diff(*tc.wantResponse, got); diff != "" {
					t.Errorf("Unexpected response (-want +got): %s", diff)
				}
			}
		})
	}
}
