package probabilityai

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"github.com/dragon162/go-get-games/games/minesweeper/gamegen"
	"github.com/pellared/fluentassert/f"
	"testing"
)

func MakeGameAndReveal(pos vector.IntVec2, gen *gamegen.GameGenerator) *game.Game {
	g := game.MakeFromGenerator(gen)
	g.Reveal(pos)
	return g
}

func TestEvaluation(t *testing.T) {
	tcs := []struct {
		name string
		g    *game.Game
		want map[vector.IntVec2]float64
	}{
		{
			name: "11111",
			g: MakeGameAndReveal(vector.Of(0, 2), gamegen.MakeGameGenFromString(""+
				"-*--*\n"+
				"11111\n"+
				"     ")),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): .5,
				vector.Of(1, 0): .5,
				vector.Of(2, 0): 0,
				vector.Of(3, 0): .5,
				vector.Of(4, 0): .5,
			},
		},
		{
			name: "solvable",
			g: MakeGameAndReveal(vector.Of(2, 2), gamegen.MakeGameGenFromString(""+
				"*1 \n"+
				"22 \n"+
				"*1 ")),
			want: map[vector.IntVec2]float64{
				vector.Of(0, 0): 1,
				vector.Of(0, 1): 0,
				vector.Of(0, 2): 1,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ai := &ProbabilityAI{}
			got := ai.ScoreAndFlagDaBoard(tc.g)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
