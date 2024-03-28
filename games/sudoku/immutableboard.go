package sudoku

type ImmutableBoard struct {
	board *Board
}

// Confirm we meet the interface
var _ ReadOnlyBoard = &ImmutableBoard{}

func (i *ImmutableBoard) Get(x, y int) Move {
	return i.board.Get(x, y)
}

func (i *ImmutableBoard) GetI(index int) Move {
	return i.board.GetI(index)
}

func (i *ImmutableBoard) Set(x, y int, num Move) *ImmutableBoard {
	ret := &ImmutableBoard{board: i.board.Clone()}
	ret.board.Set(x, y, num)
	return ret
}

func (i *ImmutableBoard) SetI(index int, num Move) *ImmutableBoard {
	ret := &ImmutableBoard{board: i.board.Clone()}
	ret.board.SetI(index, num)
	return ret
}
