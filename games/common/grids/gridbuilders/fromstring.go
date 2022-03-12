package gridbuilders

import (
	"fyne.io/fyne/v2/theme"
	"github.com/dragon162/go-get-games/games/common/grids"
	"github.com/dragon162/go-get-games/games/common/vector"
	"strings"
)

func MakeSimpleGridFromString(input string) (*grids.SimpleGrid[rune], error) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	var width int
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	grid, err := grids.MakeSimpleGrid[rune](width, height)
	if err != nil {
		return nil, err
	}
	for y, line := range lines {
		for x, val := range line {
			pos := vector.Of(x, y)
			if err := grid.Set(pos, val); err != nil {
				return nil, err
			}
		}
	}
	return grid, nil
}

func MakeGuiGridFromString(input string, defaultCode *theme.ThemedResource, translator func(rune) *theme.ThemedResource) (*grids.GuiGrid[rune], error) {
	tmp, err := MakeSimpleGridFromString(input)
	if err != nil {
		return nil, err
	}

	ret, err := grids.MakeGuiBoard[rune](tmp.Width(), tmp.Height(), defaultCode, translator)
	if err != nil {
		return nil, err
	}
	for y := 0; y < ret.Height(); y++ {
		for x := 0; x < ret.Width(); x++ {
			pos := vector.Of(x, y)
			if val, ok := tmp.Get(pos); ok {
				if err := ret.Set(pos, val); err != nil {
					return nil, err
				}
			}
		}
	}

	return ret, nil
}
