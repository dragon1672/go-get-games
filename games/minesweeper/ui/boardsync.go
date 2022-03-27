package ui

import (
	"fyne.io/fyne/v2/theme"
	"github.com/dragon162/go-get-games/games/minesweeper/minesweeper"
	"github.com/dragon162/go-get-games/games/minesweeper/ui/assets"
	"github.com/dragon162/go-get-games/ui/buttongrid"
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

var ch2resource = map[minesweeper.CellState]*theme.ThemedResource{
	minesweeper.CellBomb:      bombResource,
	minesweeper.CellFlag:      signalResource,
	minesweeper.CellEmpty:     emptyResource,
	minesweeper.CellMaybeBomb: tankResource,
	minesweeper.CellSafe:      targetResource,
	minesweeper.CellN0:        n0Resource,
	minesweeper.CellN1:        n1Resource,
	minesweeper.CellN2:        n2Resource,
	minesweeper.CellN3:        n3Resource,
	minesweeper.CellN4:        n4Resource,
	minesweeper.CellN5:        n5Resource,
	minesweeper.CellN6:        n6Resource,
	minesweeper.CellN7:        n7Resource,
	minesweeper.CellN8:        n8Resource,
	minesweeper.CellN9:        n9Resource,
}

func MakeAndSyncRenderableBoard(g *minesweeper.Game) *buttongrid.RenderableBoard {
	bg := buttongrid.MakeRenderableBoard(g.Width(), g.Height(), emptyResource)
	g.ChangeEvent.Subscribe(func(data minesweeper.ChangeEventData) {
		if b, ok := bg.GetButton(data.Pos); ok {
			if i, ok := ch2resource[data.Val]; ok {
				b.SetIcon(i)
			}
		}
	})

	bg.ClickEvent.Subscribe(func(data buttongrid.ClickEventData) {
		if data.Tapped {
			g.Reveal(data.Pos)
		} else {
			g.ToggleFlag(data.Pos)
		}
	})

	return bg
}
