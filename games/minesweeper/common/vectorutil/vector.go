package vectorutil

import "github.com/dragon1672/go-collections/vector"

func IterateSurroundingInclusive(pos vector.IntVec2, callback func(vec2 vector.IntVec2)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			offset := vector.Of(x, y)
			callback(pos.Add(offset))
		}
	}
}

func IterateSurroundingExclusive(pos vector.IntVec2, callback func(vec2 vector.IntVec2)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			offset := vector.Of(x, y)
			callback(pos.Add(offset))
		}
	}
}
