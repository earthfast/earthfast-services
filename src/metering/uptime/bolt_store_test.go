package uptime

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

func TestBoltDBStore(t *testing.T) {
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	dbPath := filepath.Join(t.TempDir(), "test.db")
	store, err := NewBoltDBStore(logger, dbPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	// Populate the store with test data.
	y2k := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	int0 := IntervalData{
		StartTime:    y2k,
		EndTime:      y2k.Add(DefaultIntervalDuration),
		ProbeResults: map[string]ProbeCounts{"node0": {Success: 1, Failure: 2}},
	}
	int1 := IntervalData{
		StartTime:    int0.EndTime,
		EndTime:      int0.EndTime.Add(DefaultIntervalDuration),
		ProbeResults: map[string]ProbeCounts{"node0": {Success: 1, Failure: 2}},
	}
	int2 := IntervalData{
		StartTime:    int1.EndTime,
		EndTime:      int1.EndTime.Add(DefaultIntervalDuration),
		ProbeResults: map[string]ProbeCounts{"node0": {Success: 1, Failure: 2}},
	}
	for _, data := range []IntervalData{int0, int1, int2} {
		if err := store.Put(ctx, data); err != nil {
			t.Fatal(err)
		}
	}

	cases := []struct {
		name  string
		start time.Time
		end   time.Time
		want  []IntervalData
	}{
		{
			name:  "No matching intervals",
			start: time.Time{},
			end:   time.Time{},
			want:  nil,
		},
		{
			name:  "Exact interval match",
			start: int0.StartTime,
			end:   int0.EndTime,
			want:  []IntervalData{int0},
		},
		{
			name:  "Multiple exact intervals",
			start: int0.StartTime,
			end:   int1.EndTime,
			want:  []IntervalData{int0, int1},
		},
		{
			name:  "Ignore partial intervals",
			start: int0.StartTime.Add(time.Second),
			end:   int2.EndTime.Add(-time.Second),
			want:  []IntervalData{int1},
		},
		{
			name:  "Multiple contained intervals",
			start: int0.StartTime.Add(-time.Second),
			end:   int2.EndTime.Add(time.Second),
			want:  []IntervalData{int0, int1, int2},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []IntervalData
			err = store.Range(ctx, tc.start, tc.end, func(data IntervalData) error {
				got = append(got, data)
				return nil
			})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Unexpected result (-want +got): %s", diff)
			}
		})
	}
}
