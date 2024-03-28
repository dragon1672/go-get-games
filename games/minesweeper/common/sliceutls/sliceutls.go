package sliceutls

import (
	"github.com/dragon1672/go-get-games/games/minesweeper/common/mmath"
	"math/rand"
)

func List2Map[T comparable](vals ...T) map[T]bool {
	ret := make(map[T]bool)
	for _, val := range vals {
		ret[val] = true
	}
	return ret
}

func Truncate[T any](s []T, l int) []T {
	i := mmath.Max(0, mmath.Min(len(s), l))
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

func Shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

func RandValue[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var ret T
		return ret, false
	}
	return s[rand.Int()%len(s)], true
}
