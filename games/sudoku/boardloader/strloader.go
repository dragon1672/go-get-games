package boardloader

import "github.com/dragon1672/go-get-games/games/sudoku"

func runeToMove(r rune) sudoku.Move {
	if '0' < r && r <= '9' {
		return sudoku.Move(r - '0')
	}
	return sudoku.UNSET
}

func FromStr(s string) *sudoku.Board {
	b := sudoku.BlankBoard()
	index := 0
	for _, r := range s {
		if r == '.' {
			index++
		}
		m := runeToMove(r)
		if m != sudoku.UNSET {
			b.SetI(index, m)
			index++
		}
	}
	return b
}
