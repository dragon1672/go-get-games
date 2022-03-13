package sliceutls

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/pellared/fluentassert/f"
	"testing"
)

func TestList2Map(t *testing.T) {
	tcs := []struct {
		name  string
		input []vector.IntVec2
		want  map[vector.IntVec2]bool
	}{
		{
			name:  "empty",
			input: []vector.IntVec2{},
			want:  map[vector.IntVec2]bool{},
		},
		{
			name:  "nil",
			input: nil,
			want:  map[vector.IntVec2]bool{},
		},
		{
			name:  "1 val",
			input: []vector.IntVec2{vector.Of(1, 1)},
			want:  map[vector.IntVec2]bool{vector.Of(1, 1): true},
		},
		{
			name:  "duplicate",
			input: []vector.IntVec2{vector.Of(1, 1), vector.Of(1, 1)},
			want:  map[vector.IntVec2]bool{vector.Of(1, 1): true},
		},
		{
			name:  "multiple",
			input: []vector.IntVec2{vector.Of(0, 0), vector.Of(1, 1)},
			want: map[vector.IntVec2]bool{
				vector.Of(0, 0): true,
				vector.Of(1, 1): true,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := VecList2Map(tc.input...)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestTruncate(t *testing.T) {
	tcs := []struct {
		name      string
		input     []rune
		inputSize int
		want      []rune
	}{
		{
			name:      "empty",
			input:     []rune{},
			inputSize: 0,
			want:      []rune{},
		},
		{
			name:      "nil",
			input:     nil,
			inputSize: 0,
			want:      nil,
		},
		{
			name:      "1 val -> empty",
			input:     []rune{'a'},
			inputSize: 0,
			want:      []rune{},
		},
		{
			name:      "1 val -> 1",
			input:     []rune{'a'},
			inputSize: 1,
			want:      []rune{'a'},
		},
		{
			name:      "basic",
			input:     []rune{'a', 'b', 'c', 'd', 'e'},
			inputSize: 3,
			want:      []rune{'a', 'b', 'c'},
		},
		{
			name:      "larger than list",
			input:     []rune{'a', 'b', 'c', 'd', 'e'},
			inputSize: 500,
			want:      []rune{'a', 'b', 'c', 'd', 'e'},
		},
		{
			name:      "same as list",
			input:     []rune{'a', 'b', 'c', 'd', 'e'},
			inputSize: 5,
			want:      []rune{'a', 'b', 'c', 'd', 'e'},
		},
		{
			name:      "zero",
			input:     []rune{'a', 'b', 'c', 'd', 'e'},
			inputSize: 0,
			want:      []rune{},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Truncate(tc.input, tc.inputSize)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
