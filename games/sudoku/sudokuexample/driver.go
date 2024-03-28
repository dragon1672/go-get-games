package sudokuexample

import (
	"fmt"
	"github.com/dragon1672/go-get-games/games/sudoku"
	"github.com/dragon1672/go-get-games/games/sudoku/boardloader"
)

func Drive() {
	fmt.Println("Yo")
	game := boardloader.FromStr("" +
		"4 9 . | 8 . . | 5 1 .\n" +
		". 1 8 | . 5 . | . . 6\n" +
		". . . | . 6 9 | . . 4\n" +
		"----- + ----- + -----\n" +
		". . 5 | . . . | 6 . .\n" +
		". 7 4 | 5 . 6 | 2 9 .\n" +
		"9 . . | 3 . . | 1 4 5 \n" +
		"----- + ----- + -----\n" +
		"5 . 9 | 9 4 . | . 6 .\n" +
		". . 9 | 2 7 5 | . . .\n" +
		"8 2 7 | . 3 1 | . . ." +
		"")
	fmt.Printf("board: \n%v\nvalidation errors = %v\n", sudoku.TextPrintBoard(game), sudoku.ValidateAll(game))
}
