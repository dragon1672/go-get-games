package vector

import "fmt"

type IntVec2 struct {
	X, Y int
}

func Of(x, y int) IntVec2 {
	return IntVec2{X: x, Y: y}
}

func (v *IntVec2) Add(that IntVec2) IntVec2 {
	return Of(v.X+that.X, v.Y+that.Y)
}

func (v *IntVec2) Sub(that IntVec2) IntVec2 {
	return Of(v.X-that.X, v.Y-that.Y)
}

func (v *IntVec2) Mul(scale float64) IntVec2 {
	return Of(int(float64(v.X)*scale), int(float64(v.Y)*scale))
}

func (v *IntVec2) LengthSquared() int {
	return v.X*v.X + v.Y*v.Y
}

func (v *IntVec2) String() string {
	return fmt.Sprintf("{%d,%d}", v.X, v.Y)
}

func IterateSurroundingInclusive(pos IntVec2, callback func(vec2 IntVec2)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			offset := Of(x, y)
			callback(pos.Add(offset))
		}
	}
}

func IterateSurroundingExclusive(pos IntVec2, callback func(vec2 IntVec2)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			offset := Of(x, y)
			callback(pos.Add(offset))
		}
	}
}
