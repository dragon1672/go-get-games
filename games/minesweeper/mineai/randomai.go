package mineai

import (
	minesweeper2 "github.com/dragon162/go-get-games/games/minesweeper"
	"github.com/dragon162/go-get-games/games/minesweeper/common/sliceutls"
	"github.com/dragon1672/go-collections/vector"
)

type RandomAI struct{}

func (r *RandomAI) GetMove(g minesweeper2.ReadOnlyGame) (vector.IntVec2, bool) {
	var possibleMoves []vector.IntVec2
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			pos := vector.Of(x, y)
			if g.Get(pos) == minesweeper2.CellEmpty {
				possibleMoves = append(possibleMoves, pos)
			}
		}
	}
	return sliceutls.RandValue(possibleMoves)
}
