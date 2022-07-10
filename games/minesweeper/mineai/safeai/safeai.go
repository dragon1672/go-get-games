package safeai

import (
	"github.com/dragon162/go-get-games/games/common/vectorutil"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai"
	"github.com/dragon162/go-get-games/games/minesweeper/minesweeper"
	"github.com/dragon1672/go-collections/setqueue"
	"github.com/dragon1672/go-collections/vector"
	"github.com/golang/glog"
)

type SafeAI struct {
	FlagBombs bool
	FlagSafe  bool
}

type BombEval int64

//goland:noinspection GoUnusedConst
const (
	EvalUnset BombEval = iota
	EvalUnknown
	EvalBomb
	EvalSafe
)

// ScoreAndFlagDaBoard maps the unknown spaces to an eval
func (s *SafeAI) ScoreAndFlagDaBoard(g *minesweeper.Game) map[vector.IntVec2]BombEval {
	ret := make(map[vector.IntVec2]BombEval)
	q := setqueue.FromMapKeys(g.GetAllRevealed())
	for pos, ok := q.Pop(); ok; pos, ok = q.Pop() {
		bombCount := g.Get(pos).BombCount()
		if bombCount <= 0 {
			continue // Not a number
		}

		mutatedVals := make(map[vector.IntVec2]bool) // surrounding values will be re-checked

		touchingMoves := getTouchingMoves(g, pos)

		// init to unknowns
		for pos := range touchingMoves {
			if _, ok := ret[pos]; !ok {
				if g.Get(pos) == minesweeper.CellBomb {
					ret[pos] = EvalBomb
				}
				if !g.Get(pos).Revealed() {
					ret[pos] = EvalUnknown
				}
			}
		}

		// if all possible moves == bomb count, they must all be bombs
		// Note this "breaks" if a bomb has already been revealed
		if len(touchingMoves) == bombCount {
			for touching := range touchingMoves {
				if ret[touching] != EvalBomb { // check to only re-enqueue if actually changing
					if s.FlagBombs {
						g.SetFlagged(touching)
					}
					ret[touching] = EvalBomb
					mutatedVals[touching] = true
				}
			}
		}

		// If we have flagged # of bombs as number, then the remaining must be safe
		touchingBombs := make(map[vector.IntVec2]bool)
		for touching := range touchingMoves {
			if ret[touching] == EvalBomb {
				touchingBombs[touching] = true
			}
		}
		if len(touchingBombs) == bombCount {
			// Mark all the non bombs as safe since this number is "satisfied"
			for touching := range touchingMoves {
				if ret[touching] != EvalBomb {
					if ret[touching] != EvalSafe { // check to only re-enqueue if actually changing
						if s.FlagSafe {
							g.SetSafe(touching)
						}
						ret[touching] = EvalSafe
						mutatedVals[touching] = true
					}
				}
			}
		}

		// Add back any surrounding values to be re-checked
		for pos := range mutatedVals {
			vectorutil.IterateSurroundingInclusive(pos, func(pos vector.IntVec2) {
				if g.ValidPos(pos) && g.Get(pos).BombCount() > 0 {
					q.Add(pos)
				}
			})
		}
	}
	return ret
}

func getTouchingMoves(g *minesweeper.Game, pos vector.IntVec2) map[vector.IntVec2]bool {
	ret := make(map[vector.IntVec2]bool)
	vectorutil.IterateSurroundingInclusive(pos, func(pos vector.IntVec2) {
		if g.ValidPos(pos) && !g.Get(pos).Revealed() {
			ret[pos] = true
		}
	})
	return ret
}

func selectMove(moves map[vector.IntVec2]BombEval) (vector.IntVec2, bool) {
	for pos, danger := range moves {
		if danger == EvalSafe {
			return pos, true
		}
	}
	var ret vector.IntVec2
	return ret, false
}

func (s *SafeAI) GetMove(g *minesweeper.Game) (vector.IntVec2, bool) {
	if pos, ok := selectMove(s.ScoreAndFlagDaBoard(g)); ok {
		return pos, true
	}
	glog.Warningf("No fully safe moves found, going full random")
	return (&mineai.RandomAI{}).GetMove(g)
}
