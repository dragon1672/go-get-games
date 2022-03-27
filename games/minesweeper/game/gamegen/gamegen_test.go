package gamegen

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/pellared/fluentassert/f"
	"math/rand"
	"testing"
)

func TestGenerateBombs(t *testing.T) {
	type inputData struct {
		width, height        int
		bombs                int
		discouragedPositions []vector.IntVec2
	}
	tcs := []struct {
		name     string
		input    *inputData
		randSeed int64
		want     map[vector.IntVec2]bool
	}{
		{
			name: "basic",
			input: &inputData{
				width: 10, height: 10, bombs: 3,
				discouragedPositions: []vector.IntVec2{},
			},
			randSeed: 42,
			want: map[vector.IntVec2]bool{
				vector.Of(0, 0): true,
				vector.Of(0, 5): true,
				vector.Of(5, 6): true,
			},
		},
		{
			name: "deny all",
			input: &inputData{
				width: 2, height: 2, bombs: 3,
				discouragedPositions: []vector.IntVec2{
					vector.Of(0, 0), vector.Of(1, 0),
					vector.Of(0, 1), vector.Of(1, 1),
				},
			},
			want: map[vector.IntVec2]bool{},
		},
		{
			name: "bomb all",
			input: &inputData{
				width: 2, height: 2, bombs: 4,
				discouragedPositions: []vector.IntVec2{},
			},
			want: map[vector.IntVec2]bool{
				vector.Of(0, 0): true, vector.Of(1, 0): true,
				vector.Of(0, 1): true, vector.Of(1, 1): true,
			},
		},
		{
			name: "too many bombs",
			input: &inputData{
				width: 2, height: 2, bombs: 9001,
				discouragedPositions: []vector.IntVec2{},
			},
			want: map[vector.IntVec2]bool{
				vector.Of(0, 0): true, vector.Of(1, 0): true,
				vector.Of(0, 1): true, vector.Of(1, 1): true,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rand.Seed(tc.randSeed)
			got := generateNumBombs(
				tc.input.width, tc.input.height,
				tc.input.bombs,
				tc.input.discouragedPositions...,
			)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
