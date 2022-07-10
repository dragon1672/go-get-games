package minesweeper

import (
	"fmt"
	"github.com/dragon1672/go-collections/vector"
)

type ChangeEventData struct {
	Pos vector.IntVec2
	Val CellState
}

func (c ChangeEventData) String() string {
	return fmt.Sprintf("{%v == %v}", c.Pos, c.Val)
}
