package mineai

import (
	"github.com/dragon162/go-get-games/games/common/sliceutls"
	"github.com/dragon162/go-get-games/games/minesweeper/minesweeper"
	"github.com/dragon1672/go-collections/vector"
)

type RandomAI struct{}

func (r *RandomAI) GetMove(g minesweeper.ReadOnlyGame) (vector.IntVec2, bool) {
	var possibleMoves []vector.IntVec2
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			pos := vector.Of(x, y)
			if g.Get(pos) == minesweeper.CellEmpty {
				possibleMoves = append(possibleMoves, pos)
			}
		}
	}
	return sliceutls.RandValue(possibleMoves)
}
