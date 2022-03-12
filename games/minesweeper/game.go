package minesweeper

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/dragon162/go-get-games/games/common/grids"
	"github.com/dragon162/go-get-games/games/common/grids/gridbuilders"
	"github.com/dragon162/go-get-games/games/minesweeper/ui/assets"
	"github.com/dragon162/go-get-games/ui/buttongrid"
	"log"
)

type CellState int64

const (
	CellUnset CellState = iota
	CellEmpty
	CellFlag
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

var (
	emptyResource  = theme.NewThemedResource(assets.EmptyIcon)
	bombResource   = theme.NewThemedResource(assets.BombIcon)
	signalResource = theme.NewThemedResource(assets.SignalIcon)
	tankResource   = theme.NewThemedResource(assets.TankIcon)
	targetResource = theme.NewThemedResource(assets.TargetIcon)
	n0Resource     = theme.NewThemedResource(assets.N0Icon)
	n1Resource     = theme.NewThemedResource(assets.N1Icon)
	n2Resource     = theme.NewThemedResource(assets.N2Icon)
	n3Resource     = theme.NewThemedResource(assets.N3Icon)
	n4Resource     = theme.NewThemedResource(assets.N4Icon)
	n5Resource     = theme.NewThemedResource(assets.N5Icon)
	n6Resource     = theme.NewThemedResource(assets.N6Icon)
	n7Resource     = theme.NewThemedResource(assets.N7Icon)
	n8Resource     = theme.NewThemedResource(assets.N8Icon)
	n9Resource     = theme.NewThemedResource(assets.N9Icon)
)

var ch2resource = map[rune]*theme.ThemedResource{
	'b': bombResource,
	'f': targetResource,
	't': tankResource,
	's': signalResource,
	' ': emptyResource,
	'0': n0Resource,
	'1': n1Resource,
	'2': n2Resource,
	'3': n3Resource,
	'4': n4Resource,
	'5': n5Resource,
	'6': n6Resource,
	'7': n7Resource,
	'8': n8Resource,
	'9': n9Resource,
}

func MakeMineSweeperBoard() (*grids.GuiGrid[rune], error) {
	return gridbuilders.MakeGuiGridFromString(""+
		"bbbbbbb\n"+
		"f 123 f\n"+
		"b 456 b\n"+
		"t 789 t\n"+
		"bts0stb\n"+
		"",
		emptyResource,
		func(r rune) *theme.ThemedResource {
			if val, ok := ch2resource[r]; ok {
				return val
			}
			return emptyResource
		})
}

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
