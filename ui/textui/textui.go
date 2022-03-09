package textui

import (
	"bytes"
	"fyne.io/fyne/v2/widget"
	"go-get-games/games/common/grid"
	"go-get-games/games/common/vector"
)

type TextWidgetUi struct {
	grid *grid.SimpleGrid[rune]
	widg *widget.TextGrid
}

func (t *TextWidgetUi) Sync() {
	// To prevent blinking, use set text which will bundle all changes within a single refresh
	// This will remove any invalid rows/columns and update any entries without blinking the UI
	var buffer bytes.Buffer
	for y := 0; y < t.grid.Height(); y++ {
		for x := 0; x < t.grid.Width(); x++ {
			if val, err := t.grid.Get(vector.Of(x, y)); err == nil {
				buffer.WriteRune(val)
			}
		}
		buffer.WriteRune('\n')
	}
	t.widg.SetText(buffer.String())
}

func (t *TextWidgetUi) Widget() *widget.TextGrid {
	return t.widg
}
func (t *TextWidgetUi) Grid() *grid.SimpleGrid[rune] {
	return t.grid
}

func Construct(grid *grid.SimpleGrid[rune]) *TextWidgetUi {
	ret := &TextWidgetUi{
		grid: grid,
		widg: widget.NewTextGrid(),
	}
	ret.Sync()
	return ret
}
