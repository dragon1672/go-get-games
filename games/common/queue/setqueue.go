package queue

import "github.com/dragon162/go-get-games/games/common/vector"

type SetVecQueue struct {
	data map[vector.IntVec2]bool
}

func (s *SetVecQueue) Add(val vector.IntVec2) {
	if s.data == nil {
		s.data = make(map[vector.IntVec2]bool)
	}
	s.data[val] = true
}
func (s *SetVecQueue) Pop() (vector.IntVec2, bool) {
	if s.Size() == 0 {
		var ret vector.IntVec2
		return ret, false
	}
	for k := range s.data {
		delete(s.data, k)
		return k, true
	}
	var ret vector.IntVec2
	return ret, false
}

func (s *SetVecQueue) Size() int {
	return len(s.data)
}

func FromSlice(s []vector.IntVec2) *SetVecQueue {
	ret := &SetVecQueue{}
	for _, v := range s {
		ret.Add(v)
	}
	return ret
}

func FromMap[T any](m map[vector.IntVec2]T) *SetVecQueue {
	ret := &SetVecQueue{}
	for k := range m {
		ret.Add(k)
	}
	return ret
}
