package minesweeper

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dragon162/go-get-games/games/minesweeper/game"
	"github.com/dragon162/go-get-games/games/minesweeper/gamegen"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai/probabilityai"
	"github.com/dragon162/go-get-games/games/minesweeper/mineai/safeai"
	"github.com/dragon162/go-get-games/games/minesweeper/ui"
	"sync"
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
	g := game.MakeFromGenerator(gamegen.ExpertGame)
	//g := game.MakeFromGenerator(gamegen.InsaneGame)
	//g := game.MakeFromGenerator(&gamegen.GameGenerator{Width: 50, Height: 30, Gen: gamegen.IntermediateDifficulty})

	w.SetContent(container.NewVBox(
		ui.MakeAndSyncRenderableBoard(g).CanvasObj(),
	))

	//* Play with AutoFlagger
	go func() {
		ai := &safeai.SafeAI{
			FlagBombs: true,
			FlagSafe:  true,
		}
		processing := sync.Mutex{}
		g.ChangeEvent.Subscribe(func(data game.ChangeEventData) {
			if processing.TryLock() {
				ai.ScoreAndFlagDaBoard(g)
				processing.Unlock()
			}
		})
	}()
	//*/

	//*
	go func() {
		//ai := &mineai.RandomAI{}
		//ai := &safeai.SafeAI{}
		ai := &probabilityai.ProbabilityAI{}
		time.Sleep(time.Second * 10)
		mineai.AutoPlay(ai, g, time.Millisecond*500)
	}()
	//*/

	w.ShowAndRun()
}
