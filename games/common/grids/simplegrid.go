package grids

import (
	"fmt"
	"go-get-games/games/common/events"
	"go-get-games/games/common/vector"
	"strings"
)

type gridI interface {
	Width() int
	Height() int
}

type SubscribableBoard interface {
	SubscribeToChanges(handler func(vec2 vector.IntVec2)) *events.Subscription[vector.IntVec2]
}

type baseGrid struct {
	gridI
	SubscribableBoard
	width, height int
	changeEvent   *events.Feed[vector.IntVec2]
}

func (b *baseGrid) getChangeEvent() *events.Feed[vector.IntVec2] {
	if b.changeEvent == nil {
		b.changeEvent = events.Make[vector.IntVec2]()
	}
	return b.changeEvent
}
func (b *baseGrid) SubscribeToChanges(handler func(vec2 vector.IntVec2)) *events.Subscription[vector.IntVec2] {
	return b.getChangeEvent().Subscribe(handler)
}
func (b *baseGrid) Width() int  { return b.width }
func (b *baseGrid) Height() int { return b.height }
func (b *baseGrid) ValidPos(pos vector.IntVec2) error {
	if pos.X() < 0 {
		return fmt.Errorf("X (%d) must not be less than 0", pos.X())
	}
	if pos.Y() < 0 {
		return fmt.Errorf("X (%d) must not be less than 0", pos.Y())
	}
	if pos.X() >= b.Width() {
		return fmt.Errorf("X (%d) must less than width (%d)", pos.X(), b.Width())
	}
	if pos.Y() >= b.Height() {
		return fmt.Errorf("Y (%d) must less than height (%d)", pos.Y(), b.Height())
	}
	return nil
}

type SimpleGrid[T any] struct {
	data map[vector.IntVec2]T
	baseGrid
}

func (s *SimpleGrid[T]) Get(pos vector.IntVec2) (T, bool) {
	if err := s.ValidPos(pos); err != nil {
		return s.zeroVal(), false
	}
	if val, ok := s.data[pos]; ok {
		return val, true
	}
	return s.zeroVal(), true
}

func (s *SimpleGrid[T]) Set(pos vector.IntVec2, val T) error {
	if err := s.ValidPos(pos); err != nil {
		return err
	}
	s.data[pos] = val
	s.getChangeEvent().Send(pos)
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
	ret := &SimpleGrid[T]{
		data: make(map[vector.IntVec2]T),
	}
	ret.width = width
	ret.height = height
	return ret, nil
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
