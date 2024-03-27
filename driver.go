package main

import (
	"math/rand"
	"time"

	"github.com/dragon162/go-get-games/games/minesweeper"
)

func sweeperTime() {
	rand.Seed(time.Now().UTC().UnixNano())
	minesweeper.Drive()
}

func main() {
	sweeperTime()
}
