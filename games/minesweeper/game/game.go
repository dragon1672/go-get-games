package game

import (
	"github.com/dragon162/go-get-games/games/common/events"
	"github.com/dragon162/go-get-games/games/common/sliceutls"
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/games/minesweeper/gamegen"
	"github.com/golang/glog"
)

type Game struct {
	width, height int
	bombs         map[vector.IntVec2]bool
	flags         map[vector.IntVec2]bool
	gen           gamegen.GameStateGenerator
	revealed      map[vector.IntVec2]CellState
	ChangeEvent   events.Feed[ChangeEventData] // notably isn't a pointer so each event gets an independent copy
}

func (g *Game) Width() int  { return g.width }
func (g *Game) Height() int { return g.height }
func (g *Game) ValidPos(pos vector.IntVec2) bool {
	return 0 <= pos.X && pos.X < g.Width() &&
		0 <= pos.Y && pos.Y < g.Height()
}
func (g *Game) NumBombs() int { return len(g.bombs) }

func (g *Game) Get(pos vector.IntVec2) CellState {
	if val, ok := g.revealed[pos]; ok {
		return val
	}
	if flagged := g.flags[pos]; flagged {
		return CellFlag
	}
	return CellEmpty
}

func (g *Game) Reveal(pos vector.IntVec2) []ChangeEventData {
	g.ensureGen(pos)
	ret := g.silentReveal(pos)
	// Emit after full updated
	for _, val := range ret {
		g.ChangeEvent.Send(val)
	}

	return ret
}

// ToggleFlag will toggle the square as flagged, returns true if newly flagged.
func (g *Game) ToggleFlag(pos vector.IntVec2) bool {
	if current := g.Get(pos); current != CellEmpty && current != CellFlag {
		glog.Warningf("Cannot flag already revealed location at %v. Already a %v.", pos, current)
		return false
	}
	if flagged := g.flags[pos]; flagged {
		delete(g.flags, pos)
	} else {
		g.flags[pos] = true
	}
	g.ChangeEvent.Send(ChangeEventData{Pos: pos, Val: g.Get(pos)})
	return g.flags[pos]
}

func (g *Game) GetAllRevealed() map[vector.IntVec2]CellState {
	// make a copy
	var ret map[vector.IntVec2]CellState
	for key, val := range g.revealed {
		ret[key] = val
	}
	return ret
}

// silentlySet will not emit an event
func (g *Game) silentlySet(pos vector.IntVec2, s CellState) ChangeEventData {
	if g.revealed == nil {
		g.revealed = make(map[vector.IntVec2]CellState)
	}
	g.revealed[pos] = s
	if flagged := g.flags[pos]; flagged {
		delete(g.flags, pos)
	}
	return ChangeEventData{pos, s}
}

// silentReveal will appropriately reveal all adjacent squares
func (g *Game) silentReveal(toCheck ...vector.IntVec2) []ChangeEventData {
	var ret []ChangeEventData
	for len(toCheck) > 0 {
		var ok bool
		var pos vector.IntVec2
		toCheck, pos, ok = sliceutls.PopLast(toCheck)
		if !ok {
			glog.Warning("Unexpected error when popping from list, as pop should always exceed ")
			continue // unexpected
		}
		if !g.ValidPos(pos) {
			continue // invalid pos
		}
		if _, revealed := g.revealed[pos]; revealed {
			continue // already revealed or invalid
		}
		state := g.calculateAsState(pos)
		ret = append(ret, g.silentlySet(pos, state))
		if state == CellN0 {
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					// add everything including self and let loop handle
					newPos := pos.Add(vector.Of(x, y))
					if g.ValidPos(newPos) {
						toCheck = append(toCheck, newPos)
					}
				}
			}
		}
	}
	return ret
}

// ensureGen makes sure there are bombs. Otherwise noop.
func (g *Game) ensureGen(discouragedPositions ...vector.IntVec2) {
	if g.bombs == nil {
		g.bombs = g.gen.GenerateBombs(g.width, g.height, discouragedPositions...)
	}
}

// calculateAsState calculates the CellState (readonly)
func (g *Game) calculateAsState(pos vector.IntVec2) CellState {
	if isBomb := g.bombs[pos]; isBomb {
		return CellBomb
	}
	return CellStateFromBombCount(g.calcNum(pos))
}

// calcNum calculates adjacent bombs  (readonly)
func (g *Game) calcNum(pos vector.IntVec2) int {
	g.ensureGen(pos)
	bombCount := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			// skip current pos
			if x == 0 && y == 0 {
				continue
			}
			offset := vector.Of(x, y)
			toCheck := pos.Add(offset)
			if b := g.bombs[toCheck]; b {
				bombCount++
			}
		}
	}
	return bombCount
}

func MakeFromGenerator(gen *gamegen.GameGenerator) *Game {
	return &Game{
		width:  gen.Width,
		height: gen.Height,
		gen:    gen.Gen,
		flags:  make(map[vector.IntVec2]bool),
	}
}
