package sliceutls

import (
	"github.com/dragon162/go-get-games/games/common/mmath"
	"github.com/dragon162/go-get-games/games/common/vector"
)

func VecList2Map(vals ...vector.IntVec2) map[vector.IntVec2]bool {
	ret := make(map[vector.IntVec2]bool)
	for _, val := range vals {
		ret[val] = true
	}
	return ret
}

func Truncate[T any](s []T, l int) []T {
	i := mmath.MaxInt(0, mmath.MinInt(len(s), l))
	return s[:i]
}

// PopLast  removes the last element and returns it, also true if successful
func PopLast[T any](s []T) ([]T, T, bool) {
	if len(s) == 0 {
		var d T
		return s, d, false
	}
	ret := s[len(s)-1]
	s = s[:len(s)-1]
	return s, ret, true
}
