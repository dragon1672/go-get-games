package mineexample

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dragon1672/go-get-games/games/minesweeper"
	"github.com/dragon1672/go-get-games/games/minesweeper/mineai"
	"github.com/dragon1672/go-get-games/games/minesweeper/mineai/probabilityai"
	"github.com/dragon1672/go-get-games/games/minesweeper/mineai/safeai"
	"github.com/dragon1672/go-get-games/games/minesweeper/ui"
	"sync"
	"time"
)

// Drive starts a new bugs minesweeper
func Drive() {

	a := app.New()
	w := a.NewWindow("Mines!")

	/*
		g := minesweeper.MakeFromGenerator(gamegen.MakeGameGenFromString("" +
			"*1 \n" +
			"22 \n" +
			"*1 "))
		//*/
	//g := minesweeper.MakeFromGenerator(gamegen.ExpertGame)
	//g := minesweeper.MakeFromGenerator(gamegen.InsaneGame)
	g := minesweeper.MakeFromGenerator(&minesweeper.GameGenerator{Width: 10, Height: 10, BigOpening: true, Gen: minesweeper.InsaneDifficulty})
	//g := minesweeper.MakeFromGenerator(&gamegen.GameGenerator{Width: 50, Height: 30, Gen: gamegen.IntermediateDifficulty})

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
		g.ChangeEvent.Subscribe(func(data minesweeper.ChangeEventData) {
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
		ai := &probabilityai.ProbabilityAI{FlagBombs: true}
		time.Sleep(time.Second * 10)
		mineai.AutoPlay(ai, g, time.Millisecond*500)
	}()
	//*/

	w.ShowAndRun()
}
