package game

import (
	"github.com/dragon162/go-get-games/games/common/vector"
	"strings"
)

type ReadOnlyGame interface {
	Width() int
	Height() int
	ValidPos(pos vector.IntVec2) bool
	NumBombs() int
	Get(pos vector.IntVec2) CellState
	GetAllRevealed() map[vector.IntVec2]CellState
}

type readonlyCopy struct {
	width, height, bombCount int
	annotations              map[vector.IntVec2]annotation
	revealed                 map[vector.IntVec2]CellState
}

func (g *readonlyCopy) Width() int  { return g.width }
func (g *readonlyCopy) Height() int { return g.height }
func (g *readonlyCopy) ValidPos(pos vector.IntVec2) bool {
	return 0 <= pos.X && pos.X < g.Width() &&
		0 <= pos.Y && pos.Y < g.Height()
}
func (g *readonlyCopy) NumBombs() int { return g.bombCount }
func (g *readonlyCopy) String() string {
	sb := strings.Builder{}
	sb.Grow(g.Height() * (g.Width() + 1))
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			sb.WriteRune(g.Get(vector.Of(x, y)).Char())
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (g *readonlyCopy) Get(pos vector.IntVec2) CellState {
	if val, ok := g.revealed[pos]; ok {
		return val
	}
	if f, ok := g.annotations[pos]; ok {
		return f.State()
	}
	return CellEmpty
}

func (g *readonlyCopy) GetAllRevealed() map[vector.IntVec2]CellState {
	// make a copy
	ret := make(map[vector.IntVec2]CellState)
	for key, val := range g.revealed {
		ret[key] = val
	}
	return ret
}
