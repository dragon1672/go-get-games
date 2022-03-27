package mineai

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/minesweeper"
	"github.com/golang/glog"
	"time"
)

type AutoPlayableAI interface {
	GetMove(g *minesweeper.Game) (vector.IntVec2, bool)
}

func AutoPlay(ai AutoPlayableAI, g *minesweeper.Game, delay time.Duration) {
	for move, ok := ai.GetMove(g); ok; move, ok = ai.GetMove(g) {
		g.Reveal(move)
		if g.GameOver() {
			glog.Infof("Game Over! Win: %v, Lose: %v", g.HasWon(), g.HasLost())
			break
		}
		time.Sleep(delay)
	}
}
