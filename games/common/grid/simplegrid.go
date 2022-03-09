package grid

import (
	"fmt"
	"go-get-games/games/common/vector"
	"strings"
)

type SimpleGrid[T any] struct {
	width, height int
	data          map[vector.IntVec2]T
}

func (s SimpleGrid[T]) ValidPos(pos vector.IntVec2) error {
	if pos.X() < 0 {
		return fmt.Errorf("X (%d) must not be less than 0", pos.X())
	}
	if pos.Y() < 0 {
		return fmt.Errorf("X (%d) must not be less than 0", pos.Y())
	}
	if pos.X() >= s.Width() {
		return fmt.Errorf("X (%d) must less than width (%d)", pos.X(), s.Width())
	}
	if pos.Y() >= s.Height() {
		return fmt.Errorf("Y (%d) must less than height (%d)", pos.Y(), s.Height())
	}
	return nil
}

func (s *SimpleGrid[T]) Width() int  { return s.width }
func (s *SimpleGrid[T]) Height() int { return s.height }
func (s *SimpleGrid[T]) Get(pos vector.IntVec2) (T, error) {
	if err := s.ValidPos(pos); err != nil {
		return s.zeroVal(), err
	}
	if val, ok := s.data[pos]; ok {
		return val, nil
	}
	return s.zeroVal(), nil
}
func (s *SimpleGrid[T]) Set(pos vector.IntVec2, val T) error {
	if err := s.ValidPos(pos); err != nil {
		return err
	}
	s.data[pos] = val
	return nil
}

func (s *SimpleGrid[T]) zeroVal() T {
	var ret T
	return ret
}

func MakeSimpleGrid[T any](width, height int) (*SimpleGrid[T], error) {
	if width <= 0 {
		return nil, fmt.Errorf("width (%d) must be positive", width)
	}
	if height <= 0 {
		return nil, fmt.Errorf("height (%d) must be positive", height)
	}
	return &SimpleGrid[T]{
		width:  width,
		height: height,
		data:   make(map[vector.IntVec2]T),
	}, nil
}

func MakeSimpleGridFromString(input string) (*SimpleGrid[rune], error) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	var width int
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	grid, err := MakeSimpleGrid[rune](width, height)
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
