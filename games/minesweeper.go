package main

import (
	"github.com/dragon162/go-get-games/games/minesweeper"
)

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	minesweeper.Drive()
}
