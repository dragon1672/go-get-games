package probabilityai

import (
	"github.com/dragon162/go-get-games/games/common/queue"
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai"
	"github.com/golang/glog"
)

type BombEval int64

const (
	EvalUnset BombEval = iota
	EvalUnknown
	EvalBomb
	EvalSafe
)

type ProbabilityAI struct{}

/*
Desired calculation
11111
?????

11111
*--?? valid-mid
	*--*- valid
	*---* invalid missing middle
11111
-*--? valid-mid
	-*--* valid
11111
--*-- invalid edge 1s unmet

11111
?--*- valid-mid

11111
??--* valid-mid
	*---* invalid missing middle
	-*--* valid

Final results of valid
-*--* valid
*--*- valid
-*--* valid
*--*- valid
22022 out of 4 results
final output: .5 .5 0 .5 .5


Need a method to determine an invalid board state to discard results.
Only count fully resolved boards, as to ensure all future possibilities aren't invalid
*/

// ScoreAndFlagDaBoard maps the unknown spaces to an eval
// 1 == 100% a bomb
// 0 == 0% a bomb
// .5 == 50% a bomb
func (p *ProbabilityAI) ScoreAndFlagDaBoard(g game.ReadOnlyGame) map[vector.IntVec2]float64 {
	m, ok := p.evalBoard(g, make(map[vector.IntVec2]BombEval))
	if !ok {
		glog.Error("Board State unexpectedly has an error...ignoring that entirely and proceeding")
	}
	possibilities := p.resolveUnknowns(g, m)
	bombCount := make(map[vector.IntVec2]int)
	ret := make(map[vector.IntVec2]float64)
	for pi, possibility := range possibilities {
		for k, v := range possibility {
			if v == EvalBomb {
				bombCount[k]++
			} else if v != EvalSafe {
				glog.Infof("Unexpected unknown in result set at possibility index %d pos: %v, of %v", pi, k, v)
			}
			if !g.Get(k).Revealed() {
				ret[k] = 0 // make sure we have keys to capture fully safe moves
			}
		}
	}

	for k, v := range bombCount {
		if !g.Get(k).Revealed() {
			ret[k] = float64(v) / float64(len(possibilities))
		}
	}

	return ret
}

func (p *ProbabilityAI) evalBoard(g game.ReadOnlyGame, ret map[vector.IntVec2]BombEval) (map[vector.IntVec2]BombEval, bool) {
	q := queue.FromMapKeys(g.GetAllRevealed())
	for pos, ok := q.Pop(); ok; pos, ok = q.Pop() {
		bombCount := g.Get(pos).BombCount()
		if bombCount <= 0 {
			continue // Not a number
		}
		mutatedVals := make(map[vector.IntVec2]bool) // surrounding values will be re-checked

		touchingMoves := getSurrounding(g, pos)

		// init to unknowns
		for pos := range touchingMoves {
			if _, ok := ret[pos]; !ok {
				if g.Get(pos) == game.CellBomb {
					ret[pos] = EvalBomb
				}
				if !g.Get(pos).Revealed() {
					ret[pos] = EvalUnknown
				}
			}
		}

		// if all possible moves == bomb count, they must all be bombs
		var touchingBombsOrUnknowns []vector.IntVec2
		touchingBombs := make(map[vector.IntVec2]bool)
		for pos := range touchingMoves {
			if v, ok := ret[pos]; ok {
				if v == EvalBomb || v == EvalUnknown {
					touchingBombsOrUnknowns = append(touchingBombsOrUnknowns, pos)
				}
				if v == EvalBomb {
					touchingBombs[pos] = true
				}
			}
		}
		if len(touchingBombsOrUnknowns) < bombCount {
			// Invalid game state not enough possible bombs!
			return nil, false
		}
		if len(touchingBombs) > bombCount {
			// Invalid state too many bombs!
			return nil, false
		}

		// shortcut if fully evaluated
		allResolved := true
		for pos := range touchingMoves {
			if v, ok := ret[pos]; ok && v == EvalUnknown || !g.Get(pos).Revealed() {
				allResolved = false
				break
			}
		}
		if allResolved {
			return ret, true
		}

		if len(touchingBombsOrUnknowns) == bombCount {
			for _, touching := range touchingBombsOrUnknowns {
				if ret[touching] != EvalBomb { // only mutate if we are actually changing
					ret[touching] = EvalBomb
					mutatedVals[touching] = true
				}
			}
		}

		// If we have flagged # of bombs as number, then the remaining must be safe

		if len(touchingBombs) == bombCount {
			// Mark all the non bombs as safe since this number is "satisfied"
			for touching := range touchingMoves {
				if ret[touching] != EvalBomb && !g.Get(touching).Revealed() {
					if ret[touching] != EvalSafe { // check to only re-enqueue if actually changing
						ret[touching] = EvalSafe
						mutatedVals[touching] = true
					}
				}
			}
		}

		// Mutated values are unrevealed positions that changed, add back any surrounding numbered positions to re-check
		for pos := range mutatedVals {
			vector.IterateSurroundingInclusive(pos, func(pos vector.IntVec2) {
				if g.ValidPos(pos) && g.Get(pos).Revealed() {
					q.Add(pos)
				}
			})
		}
	}
	return ret, true
}

func (p *ProbabilityAI) resolveUnknowns(g game.ReadOnlyGame, m map[vector.IntVec2]BombEval) []map[vector.IntVec2]BombEval {
	fullyResolved := true
	for _, v := range m {
		if v == EvalUnknown {
			fullyResolved = false
			break
		}
	}
	if fullyResolved {
		return []map[vector.IntVec2]BombEval{dup(m)}
	}

	type funcResponse []map[vector.IntVec2]BombEval
	var futures []chan funcResponse
	for k, v := range m {
		if v != EvalUnknown {
			continue
		}
		mm := dup(m)
		mm[k] = EvalBomb
		c := make(chan funcResponse, 1)
		futures = append(futures, c)
		go func() {
			if mm, ok := p.evalBoard(g, mm); ok {
				c <- p.resolveUnknowns(g, mm)
			} else {
				c <- nil
			}
		}()
	}

	var ret []map[vector.IntVec2]BombEval
	for _, f := range futures {
		ret = append(ret, <-f...)
	}

	return ret
}

func dup(m map[vector.IntVec2]BombEval) map[vector.IntVec2]BombEval {
	r := make(map[vector.IntVec2]BombEval)
	for k, v := range m {
		r[k] = v
	}
	return r
}

func getSurrounding(g game.ReadOnlyGame, p vector.IntVec2) map[vector.IntVec2]bool {
	ret := make(map[vector.IntVec2]bool)
	vector.IterateSurroundingExclusive(p, func(pos vector.IntVec2) {
		if g.ValidPos(pos) {
			ret[pos] = true
		}
	})
	return ret
}

func selectMove(moves map[vector.IntVec2]float64) (vector.IntVec2, float64) {
	var safestMove vector.IntVec2
	safestVal := 2.0 // 200% chance of bomb is higher than any possible value
	for pos, danger := range moves {
		if danger < safestVal {
			glog.Infof("New safest position found %f @ %v", danger, pos)
			safestMove, safestVal = pos, danger
		}
	}
	return safestMove, safestVal
}

func (p *ProbabilityAI) getSimpleMove(g game.ReadOnlyGame) (vector.IntVec2, bool) {
	m, ok := p.evalBoard(g, make(map[vector.IntVec2]BombEval))
	if !ok {
		glog.Error("Board State unexpectedly has an error...ignoring that entirely and proceeding")
	}
	for k, v := range m {
		if v == EvalSafe && !g.Get(k).Revealed() {
			return k, true
		}
	}
	var ret vector.IntVec2
	return ret, false
}

func (p *ProbabilityAI) GetMove(g game.ReadOnlyGame) (vector.IntVec2, bool) {
	if pos, ok := p.getSimpleMove(g); ok {
		return pos, true
	}
	glog.Info("No simple moves found, calculating possible bomb locations")
	moves := p.ScoreAndFlagDaBoard(g)
	if len(moves) > 0 {
		pos, _ := selectMove(moves)
		return pos, true
	}
	glog.Warningf("No fully safe moves found, going full random")
	return (&mineai.RandomAI{}).GetMove(g)
}
