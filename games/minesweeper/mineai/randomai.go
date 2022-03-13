package mineai

import (
	"github.com/dragon162/go-get-games/games/common/sliceutls"
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"time"
)

type RandomAI struct{}

func (r *RandomAI) getMoves(g *game.Game) []vector.IntVec2 {
	var possibleMoves []vector.IntVec2
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			pos := vector.Of(x, y)
			if g.Get(pos) == game.CellEmpty {
				possibleMoves = append(possibleMoves, pos)
			}
		}
	}
	sliceutls.Shuffle(possibleMoves)
	return possibleMoves
}

func (r *RandomAI) Play(g *game.Game, delay int) {

	possibleMoves := r.getMoves(g)
	for len(possibleMoves) > 0 {
		_, pos, _ := sliceutls.PopLast(possibleMoves)
		g.Reveal(pos)
		time.Sleep(time.Millisecond * time.Duration(delay))
		possibleMoves = r.getMoves(g)
	}

}
