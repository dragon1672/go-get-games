package game

import "github.com/golang/glog"

type CellState int64

const (
	CellUnset CellState = iota
	CellEmpty
	CellFlag
	CellSafe
	CellMaybeBomb
	CellBomb
	CellN0
	CellN1
	CellN2
	CellN3
	CellN4
	CellN5
	CellN6
	CellN7
	CellN8
	CellN9
)

func (e CellState) String() string {
	switch e {
	case CellEmpty:
		return "Empty"
	case CellFlag:
		return "Flag"
	case CellSafe:
		return "CellSafe"
	case CellMaybeBomb:
		return "CellMaybeBomb"
	case CellBomb:
		return "Bomb"
	case CellN0:
		return "N0"
	case CellN1:
		return "N1"
	case CellN2:
		return "N2"
	case CellN3:
		return "N3"
	case CellN4:
		return "N4"
	case CellN5:
		return "N5"
	case CellN6:
		return "N6"
	case CellN7:
		return "N7"
	case CellN8:
		return "N8"
	case CellN9:
		return "N9"
	default:
		return "??"
	}
}

func (e CellState) Char() rune {
	switch e {
	case CellEmpty:
		return '.'
	case CellFlag:
		return 'F'
	case CellSafe:
		return 'S'
	case CellMaybeBomb:
		return 'M'
	case CellBomb:
		return 'B'
	case CellN0:
		return '0'
	case CellN1:
		return '1'
	case CellN2:
		return '2'
	case CellN3:
		return '3'
	case CellN4:
		return '4'
	case CellN5:
		return '5'
	case CellN6:
		return '6'
	case CellN7:
		return '7'
	case CellN8:
		return '8'
	case CellN9:
		return '9'
	default:
		return '?'
	}
}

func (e CellState) BombCount() int {
	switch e {
	case CellN0:
		return 0
	case CellN1:
		return 1
	case CellN2:
		return 2
	case CellN3:
		return 3
	case CellN4:
		return 4
	case CellN5:
		return 5
	case CellN6:
		return 6
	case CellN7:
		return 7
	case CellN8:
		return 8
	case CellN9:
		return 9
	default:
		return -1
	}
}
func (e CellState) Revealed() bool {
	if e == CellBomb {
		return true
	}
	return e.BombCount() >= 0
}

func CellStateFromBombCount(n int) CellState {
	switch n {
	case 0:
		return CellN0
	case 1:
		return CellN1
	case 2:
		return CellN2
	case 3:
		return CellN3
	case 4:
		return CellN4
	case 5:
		return CellN5
	case 6:
		return CellN6
	case 7:
		return CellN7
	case 8:
		return CellN8
	case 9:
		return CellN9
	default:
		glog.Warningf("Unsupported bomb count %d", n)
		return CellUnset
	}
}
