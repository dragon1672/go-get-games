package queue

import (
	"github.com/pellared/fluentassert/f"
	"sort"
	"testing"
)

func TestFromSlice(t *testing.T) {
	tcs := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  nil,
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "1 val",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "Duplicate",
			input: []int{42, 42},
			want:  []int{42},
		},
		{
			name:  "multiple",
			input: []int{42, 420},
			want:  []int{42, 420},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			q := FromSlice(tc.input...)
			var got []int
			for v, ok := q.Pop(); ok; v, ok = q.Pop() {
				got = append(got, v)
			}
			sort.Ints(got)
			sort.Ints(tc.want)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestFromMapKey(t *testing.T) {
	tcs := []struct {
		name  string
		input map[int]bool
		want  []int
	}{
		{
			name:  "empty",
			input: map[int]bool{},
			want:  nil,
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "1 val",
			input: map[int]bool{42: true},
			want:  []int{42},
		},
		{
			name:  "multiple",
			input: map[int]bool{42: true, 420: true},
			want:  []int{42, 420},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			q := FromMapKeys(tc.input)
			var got []int
			for v, ok := q.Pop(); ok; v, ok = q.Pop() {
				got = append(got, v)
			}
			sort.Ints(got)
			sort.Ints(tc.want)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
