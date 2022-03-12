package buttongrid

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/dragon162/go-get-games/games/common/events"
	"github.com/dragon162/go-get-games/games/common/vector"
	"github.com/dragon162/go-get-games/ui/uibuttons"
	"image/color"
)

type Board interface {
	Width() int
	Height() int
}

type ClickEventData struct {
	pos    vector.IntVec2
	tapped bool
}

func (c ClickEventData) String() string {
	return fmt.Sprintf("{tapped: %v @ %v}", c.tapped, c.pos)
}

type RenderableBoard struct {
	fyne.WidgetRenderer
	grid    *fyne.Container
	buttons map[vector.IntVec2]*uibuttons.Button

	ClickEvent events.Feed[ClickEventData] // notably isn't a pointer so each event gets an independent copy
}

func (g *RenderableBoard) MinSize() fyne.Size {
	return g.grid.MinSize()
}

func (g *RenderableBoard) Layout(size fyne.Size) {
	g.grid.Layout.Layout(g.grid.Objects, size)
}

func (g *RenderableBoard) ApplyTheme() {
}

func (g *RenderableBoard) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (g *RenderableBoard) Refresh() {
	canvas.Refresh(g.grid)
}

func (g *RenderableBoard) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{g.grid}
}

func (g *RenderableBoard) Destroy() {
}

func (g *RenderableBoard) GetButton(pos vector.IntVec2) (*uibuttons.Button, bool) {
	ret, ok := g.buttons[pos]
	return ret, ok
}

func (g *RenderableBoard) CanvasObj() fyne.CanvasObject {
	return g.grid
}

func MakeRenderableBoard(width, height int, defaultCode *theme.ThemedResource) *RenderableBoard {
	renderer := &RenderableBoard{
		buttons: make(map[vector.IntVec2]*uibuttons.Button),
	}
	var buttons []fyne.CanvasObject
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := vector.Of(x, y) // create copy of x/y for lambda
			b := uibuttons.NewButton("", defaultCode, func(tapped bool) {
				renderer.ClickEvent.Send(ClickEventData{
					pos:    pos,
					tapped: tapped,
				})
			})
			renderer.buttons[pos] = b
			buttons = append(buttons, b)
		}
	}

	renderer.grid = container.NewGridWithColumns(width, buttons...)
	return renderer
}
