package ui

import (
	"fyne.io/fyne/v2/theme"
	minesweeper2 "github.com/dragon1672/go-get-games/games/minesweeper"
	"github.com/dragon1672/go-get-games/games/minesweeper/ui/assets"
	"github.com/dragon1672/go-get-games/ui/buttongrid"
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

var ch2resource = map[minesweeper2.CellState]*theme.ThemedResource{
	minesweeper2.CellBomb:      bombResource,
	minesweeper2.CellFlag:      signalResource,
	minesweeper2.CellEmpty:     emptyResource,
	minesweeper2.CellMaybeBomb: tankResource,
	minesweeper2.CellSafe:      targetResource,
	minesweeper2.CellN0:        n0Resource,
	minesweeper2.CellN1:        n1Resource,
	minesweeper2.CellN2:        n2Resource,
	minesweeper2.CellN3:        n3Resource,
	minesweeper2.CellN4:        n4Resource,
	minesweeper2.CellN5:        n5Resource,
	minesweeper2.CellN6:        n6Resource,
	minesweeper2.CellN7:        n7Resource,
	minesweeper2.CellN8:        n8Resource,
	minesweeper2.CellN9:        n9Resource,
}

func MakeAndSyncRenderableBoard(g *minesweeper2.Game) *buttongrid.RenderableBoard {
	bg := buttongrid.MakeRenderableBoard(g.Width(), g.Height(), emptyResource)
	g.ChangeEvent.Subscribe(func(data minesweeper2.ChangeEventData) {
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
