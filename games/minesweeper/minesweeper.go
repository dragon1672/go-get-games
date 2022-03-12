package minesweeper

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"go-get-games/ui/buttongrid"
	"log"
)

// Show starts a new bugs game
func Drive() {

	a := app.New()
	w := a.NewWindow("Hello")

	ms, err := MakeMineSweeperBoard()
	if err != nil {
		log.Fatal(err)
	}
	ms.SubscribeToClicks(func(data buttongrid.ClickEventData) {
		fmt.Printf("Click Event: %v\n", data)
	})
	// https://github.com/fyne-io/examples/ for more examples to poke with
	w.SetContent(container.NewVBox(
		ms.CanvasObj(),
	))

	w.ShowAndRun()
}
