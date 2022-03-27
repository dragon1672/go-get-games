package queue

type SetQueue[T comparable] struct {
	data map[T]bool
}

func (s *SetQueue[T]) Add(val T) {
	if s.data == nil {
		s.data = make(map[T]bool)
	}
	s.data[val] = true
}
func (s *SetQueue[T]) Pop() (T, bool) {
	if s.Size() == 0 {
		var ret T
		return ret, false
	}
	for k := range s.data {
		delete(s.data, k)
		return k, true
	}
	var ret T
	return ret, false
}

func (s *SetQueue[T]) Size() int {
	return len(s.data)
}

func FromSlice[T comparable](s ...T) *SetQueue[T] {
	ret := &SetQueue[T]{}
	for _, v := range s {
		ret.Add(v)
	}
	return ret
}

func FromMapKeys[T comparable, V any](m map[T]V) *SetQueue[T] {
	ret := &SetQueue[T]{}
	for k := range m {
		ret.Add(k)
	}
	return ret
}
