package minesweeper

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"github.com/dragon162/go-get-games/games/minesweeper/gamegen"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai/safeai"
	"github.com/dragon162/go-get-games/games/minesweeper/ui"
	"time"
)

// Drive starts a new bugs game
func Drive() {

	a := app.New()
	w := a.NewWindow("Mines!")

	/*
		g := game.MakeFromGenerator(gamegen.MakeGameGenFromString("" +
			"*1 \n" +
			"22 \n" +
			"*1 "))
		//*/
	//g := game.MakeFromGenerator(gamegen.ExpertGame)
	g := game.MakeFromGenerator(gamegen.IntermediateGame)

	w.SetContent(container.NewVBox(
		ui.MakeAndSyncRenderableBoard(g).CanvasObj(),
	))

	// Play with AutoFlagger
	go func() {
		ai := &safeai.SafeAI{}
		for {
			ai.ScoreAndFlagDaBoard(g)
			time.Sleep(time.Millisecond * 1500)
		}
	}()

	/*
		go func() {
			//ai := &mineai.RandomAI{}
			ai := &safeai.SafeAI{}
			mineai.AutoPlay(ai, g, time.Millisecond*1500)
		}()
		//*/

	w.ShowAndRun()
}
