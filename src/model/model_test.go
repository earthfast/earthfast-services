package model

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomID(t *testing.T) ID {
	t.Helper()

	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		t.Fatal(err)
	}
	var id ID
	copy(id[:], data[:])
	return id
}

func TestParseID(t *testing.T) {
	testID := randomID(t)

	cases := []struct {
		name  string
		input string
		want  ID
	}{
		{
			name:  "empty input",
			input: "",
			want:  ZeroID,
		},
		{
			name:  "input too short",
			input: "abc",
			want:  ZeroID,
		},
		{
			name:  "input too long",
			input: strings.Repeat("a", 100),
			want:  ZeroID,
		},
		{
			name:  "input not hex",
			input: strings.Repeat("z", 64),
			want:  ZeroID,
		},
		{
			name:  "ok",
			input: testID.Hex(),
			want:  testID,
		},
		{
			name:  "ok with 0x prefix",
			input: "0x" + testID.Hex(),
			want:  testID,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseID(tc.input)
			if got != tc.want {
				t.Errorf("Unexpected parsed ID: got=%x, want=%x, err=%v", got, tc.want, err)
			}
		})
	}
}
