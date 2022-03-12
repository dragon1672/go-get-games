package grids

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"go-get-games/games/common/events"
	"go-get-games/games/common/vector"
	"go-get-games/ui/buttongrid"
)

type GuiGrid[T any] struct {
	baseGrid
	simpleGrid *SimpleGrid[T]
	buttonGrid *buttongrid.RenderableBoard
	translator func(T) *theme.ThemedResource
}

func (g *GuiGrid[T]) Set(pos vector.IntVec2, val T) error {
	// Simple grid will emit an event
	if err := g.simpleGrid.Set(pos, val); err != nil {
		return err
	}
	if b, ok := g.buttonGrid.GetButton(pos); ok {
		b.SetIcon(g.translator(val))
	}

	return nil
}

func (g *GuiGrid[T]) Get(pos vector.IntVec2) (T, bool) {
	return g.simpleGrid.Get(pos)
}

func (g *GuiGrid[T]) CanvasObj() fyne.CanvasObject {
	return g.buttonGrid.CanvasObj()
}

func (g *GuiGrid[T]) SubscribeToClicks(handler func(data buttongrid.ClickEventData)) *events.Subscription[buttongrid.ClickEventData] {
	return g.buttonGrid.ClickEvent.Subscribe(handler)
}

func MakeGuiBoard[T any](width, height int, defaultCode *theme.ThemedResource, translator func(T) *theme.ThemedResource) (*GuiGrid[T], error) {
	s, err := MakeSimpleGrid[T](width, height)
	if err != nil {
		return nil, err
	}
	ret := &GuiGrid[T]{
		simpleGrid: s,
		buttonGrid: buttongrid.MakeRenderableBoard(width, height, defaultCode),
		translator: translator,
	}
	ret.width = width
	ret.height = height

	return ret, nil
}
