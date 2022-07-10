package probabilityai

import (
	"github.com/dragon162/go-get-games/games/minesweeper/minesweeper"
	"github.com/dragon1672/go-collections/vector"
	"github.com/pellared/fluentassert/f"
	"testing"
)

func TestEvaluation(t *testing.T) {
	tcs := []struct {
		name string
		g    minesweeper.ReadOnlyGame
		want map[vector.IntVec2]float64
	}{
		{
			name: "solvable",
			g: minesweeper.MakeReadonlyRevealedString("" +
				".1 \n" +
				".2 \n" +
				".1 "),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): 1,
				vector.Of(0, 1): 0,
				vector.Of(0, 2): 1,
			},
		},
		{ // TODO snag more examples from https://minesweeper.online/help/patterns
			name: "B1",
			g: minesweeper.MakeReadonlyRevealedString("" +
				"  1..\n" +
				"  2..\n" +
				"  3..\n" +
				"112..\n" +
				"....."),
			want: map[vector.IntVec2]float64{
				vector.Of(3, 0): 0,
				vector.Of(3, 1): 1,
				vector.Of(3, 2): 1,
				vector.Of(3, 3): 1,
				vector.Of(3, 4): 0,
				vector.Of(2, 4): 0,
				vector.Of(1, 4): 0,
				vector.Of(0, 4): 1,
			},
		},
		{
			name: "1–1+",
			g: minesweeper.MakeReadonlyRevealedString("" +
				"......\n" +
				".2....\n" +
				"1111..\n" +
				"   1..\n" +
				"   1.."),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): 0.3333333333333333,
				vector.Of(0, 1): 1,
				vector.Of(1, 0): 0.3333333333333333,
				vector.Of(2, 0): 0.3333333333333333,
				vector.Of(2, 1): 0,
				vector.Of(3, 1): 1,
				vector.Of(4, 1): 0,
				vector.Of(4, 2): 0,
				vector.Of(4, 3): 0,
				vector.Of(4, 4): 1,
			},
		},
		{
			name: "1–2",
			g: minesweeper.MakeReadonlyRevealedString("" +
				"......\n" +
				"......\n" +
				"121111\n" +
				"      "),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 1): .5,
				vector.Of(1, 1): .5, // this should be one of the valid locations
				vector.Of(2, 1): 1,
				vector.Of(3, 1): 0,
				vector.Of(4, 1): 0,
				vector.Of(5, 1): 1,
			},
		},
		{
			name: "1–2+",
			g: minesweeper.MakeReadonlyRevealedString("" +
				"......\n" +
				".2....\n" +
				"1114..\n" +
				"   2..\n" +
				"   1.."),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): 0.3333333333333333,
				vector.Of(0, 1): 1,
				vector.Of(1, 0): 0.3333333333333333,
				vector.Of(2, 0): 0.3333333333333333,
				vector.Of(2, 1): .5,
				vector.Of(3, 1): .5,
				vector.Of(4, 1): 1,
				vector.Of(4, 2): 1,
				vector.Of(4, 3): 1,
				vector.Of(4, 4): 0,
			},
		},
		{
			name: "11111",
			g: minesweeper.MakeReadonlyRevealedString("" +
				".....\n" +
				"11111\n" +
				"     "),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): .5,
				vector.Of(1, 0): .5,
				vector.Of(2, 0): 0,
				vector.Of(3, 0): .5,
				vector.Of(4, 0): .5,
			},
		},
		{
			name: "1–3–1 corner",
			g: minesweeper.MakeReadonlyRevealedString("" +
				"......\n" +
				"......\n" +
				"2113..\n" +
				"   1..\n" +
				"   1..\n" +
				"   2.."),
			want: map[vector.IntVec2]float64{
				vector.Of(4, 1): 1,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ai := &ProbabilityAI{}
			got := ai.ScoreDaBoard(tc.g)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
