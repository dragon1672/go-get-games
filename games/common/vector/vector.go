package vector

import "fmt"

type IntVec2 struct {
	x, y int
}

func (v *IntVec2) X() int { return v.x }
func (v *IntVec2) Y() int { return v.y }

func Of(x, y int) IntVec2 {
	return IntVec2{x: x, y: y}
}

func (v *IntVec2) Add(that IntVec2) IntVec2 {
	return Of(v.x+that.x, v.y+that.y)
}

func (v *IntVec2) Sub(that IntVec2) IntVec2 {
	return Of(v.x-that.x, v.y-that.y)
}

func (v *IntVec2) Mul(scale float64) IntVec2 {
	return Of(int(float64(v.x)*scale), int(float64(v.y)*scale))
}

func (v *IntVec2) LengthSquared() int {
	return v.x*v.x + v.y*v.y
}

func (v *IntVec2) String() string {
	return fmt.Sprintf("{%d,%d}", v.x, v.y)
}
