package minesweeper

import (
	"fyne.io/fyne/v2/theme"
	"go-get-games/games/common/grids"
	"go-get-games/games/common/grids/gridbuilders"
	"go-get-games/games/minesweeper/ui/assets"
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
