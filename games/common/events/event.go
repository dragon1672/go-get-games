package events

import "sync"

type Subscription[T any] struct {
	active   bool
	callback func(data T)
}

func (s *Subscription[T]) UnSubscribe() {
	s.active = false
}

type Feed[T any] struct {
	subscribers []*Subscription[T]
	mu          sync.Mutex
}

func (f *Feed[T]) Subscribe(callback func(data T)) *Subscription[T] {
	s := &Subscription[T]{
		active:   true,
		callback: callback,
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	f.subscribers = append([]*Subscription[T]{s}, f.subscribers...)
	return s
}

// pruneSubscribers removed any inactive subscribers and returns a copy
func (f *Feed[T]) pruneSubscribers() []*Subscription[T] {
	f.mu.Lock()
	defer f.mu.Unlock()
	var dup []*Subscription[T]
	n := 0
	for _, s := range f.subscribers {
		if s.active {
			f.subscribers[n] = s
			n++
			dup = append(dup, s)
		}
	}

	f.subscribers = f.subscribers[:n]
	return dup
}

func (f *Feed[T]) Send(data T) int {
	defensiveCopy := f.pruneSubscribers()
	count := 0
	for _, s := range defensiveCopy {
		ss := s
		if s.active {
			count++
			go func() {
				ss.callback(data)
			}()
		}
	}
	return count
}
