package minesweeper

import (
	"github.com/dragon1672/go-collections/vector"
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
	return toString(g)
}

func toString(g ReadOnlyGame) string {
	sb := strings.Builder{}
	sb.Grow(g.Height() * (g.Width() + 1))
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			sb.WriteRune(g.Get(vector.Of(x, y)).Rune())
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
	return g.revealed
}

func MakeReadOnly(width, height, bombCount int, revealed map[vector.IntVec2]CellState) ReadOnlyGame {
	return &readonlyCopy{
		width:     width,
		height:    height,
		bombCount: bombCount,
		revealed:  revealed,
	}
}

func MakeReadonlyRevealedString(s string) ReadOnlyGame {
	lines := strings.Split(s, "\n")
	height := len(lines)
	var width int
	revealed := make(map[vector.IntVec2]CellState)
	for y, line := range lines {
		if len(line) > width {
			width = len(line)
		}
		for x, val := range line {
			state := CellStateFromRune(val)
			if state.Revealed() {
				pos := vector.Of(x, y)
				revealed[pos] = state
			}
		}
	}
	return MakeReadOnly(width, height, 99, revealed)
}
