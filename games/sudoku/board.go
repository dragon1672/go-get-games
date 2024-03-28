package sudoku

import (
	"github.com/golang/glog"
	"slices"
)

type ReadOnlyBoard interface {
	Get(x, y int) Move
	GetI(index int) Move
}

type Board struct {
	data []Move
}

// Confirm we meet the interface
var _ ReadOnlyBoard = &Board{}

func BlankBoard() *Board {
	return &Board{
		data: []Move{
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
			UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET, UNSET,
		},
	}
}

func (b *Board) Get(x, y int) Move {
	return b.GetI(x + y*Height)
}

func (b *Board) GetI(index int) Move {
	if 0 > index || index > Width*Height {
		x := index % Width
		y := index / Height
		glog.Errorf("input [%d] {%d, %d} is invalid on a {%d, %d} board", index, x, y, Width, Height)
	}
	return b.data[index]
}

func (b *Board) Set(x, y int, num Move) {
	b.SetI(x+y*Height, num)
}

func (b *Board) SetI(index int, num Move) {
	if 0 > index || index > Width*Height {
		x := index % Width
		y := index / Height
		glog.Errorf("input [%d] {%d, %d} is invalid on a {%d, %d} board", index, x, y, Width, Height)
	}
	b.data[index] = num
}

func (b *Board) Clone() *Board {
	return &Board{data: slices.Clone(b.data)}
}
