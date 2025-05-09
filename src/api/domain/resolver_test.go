package domain

import (
	"context"
	"testing"

	"armada-node/model"
)

func TestStaticResolver(t *testing.T) {
	cases := []struct {
		name  string
		data  map[string]model.ID
		input string
		want  model.ID
	}{
		{
			name:  "No match",
			data:  map[string]model.ID{},
			input: "example.com",
			want:  model.ZeroID,
		},
		{
			name: "Exact match",
			data: map[string]model.ID{
				"foo.com":     model.ID{1, 2, 3},
				"example.com": model.ID{4, 5, 6},
			},
			input: "example.com",
			want:  model.ID{4, 5, 6},
		},
		{
			name: "Wildcard match",
			data: map[string]model.ID{
				"foo.com":     model.ID{1, 2, 3},
				"example.com": model.ID{4, 5, 6},
				"*":           model.ID{7, 8, 9},
			},
			input: "bar.com",
			want:  model.ID{7, 8, 9},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := NewStaticResolver(tc.data)

			got, err := r.ProjectForDomain(ctx, tc.input)
			if got != tc.want {
				t.Errorf("Unexpected project ID: got=%x, want=%x, err=%v", got, tc.want, err)
			}
		})
	}
}
