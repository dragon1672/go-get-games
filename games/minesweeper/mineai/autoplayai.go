package mineai

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"time"
)

type AutoPlayableAI interface {
	GetMove(g game.ReadOnlyGame) (vector.IntVec2, bool)
}

func AutoPlay(ai AutoPlayableAI, g *game.Game, delay time.Duration) {
	for move, ok := ai.GetMove(g.SnapshotReadonly()); ok; move, ok = ai.GetMove(g.SnapshotReadonly()) {
		g.Reveal(move)
		time.Sleep(delay)
	}
}
