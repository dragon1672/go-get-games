package buttongrid

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"github.com/dragon1672/go-collections/vector"
	"github.com/dragon1672/go-get-games/games/minesweeper/common/events"
	"github.com/dragon1672/go-get-games/ui/uibuttons"
)

type Board interface {
	Width() int
	Height() int
}

type ClickEventData struct {
	Pos    vector.IntVec2
	Tapped bool
}

func (c ClickEventData) String() string {
	return fmt.Sprintf("{Tapped: %v @ %v}", c.Tapped, c.Pos)
}

type RenderableBoard struct {
	grid    *fyne.Container
	buttons map[vector.IntVec2]*uibuttons.Button

	ClickEvent events.Feed[ClickEventData] // notably isn't a pointer so each event gets an independent copy
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
					Pos:    pos,
					Tapped: tapped,
				})
			})
			renderer.buttons[pos] = b
			buttons = append(buttons, b)
		}
	}

	renderer.grid = container.NewGridWithColumns(width, buttons...)
	return renderer
}
