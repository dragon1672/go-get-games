package minesweeper

import (
	"fmt"
	"github.com/dragon162/go-get-games/games/common/vector"
)

type ChangeEventData struct {
	Pos vector.IntVec2
	Val CellState
}

func (c ChangeEventData) String() string {
	return fmt.Sprintf("{%v == %v}", c.Pos, c.Val)
}
