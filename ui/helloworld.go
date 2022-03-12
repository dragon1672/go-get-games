package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dragon162/go-get-games/games/common/grids/gridbuilders"
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/ui/textui"
	"math/rand"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	g, _ := gridbuilders.MakeSimpleGridFromString("" +
		"123\n" +
		"456\n" +
		"789\n" +
		"")
	hello := textui.Construct(g)
	w.SetContent(container.NewVBox(
		hello.Widget(),
		widget.NewButton("Hi!", func() {
			randomLetter := rune(rand.Int()%('z'-'a') + 'a')
			hello.Grid().Set(vector.Of(1, 1), randomLetter)
			hello.Sync()
		}),
	))

	w.ShowAndRun()
}
