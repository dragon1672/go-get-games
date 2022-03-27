package mmath

import (
	"github.com/pellared/fluentassert/f"
	"testing"
)

func TestMin(t *testing.T) {
	tcs := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "Basic",
			a:    1, b: 2,
			want: 1,
		},
		{
			name: "negative value",
			a:    -1, b: 2,
			want: -1,
		},
		{
			name: "negative values",
			a:    -1, b: -2,
			want: -2,
		},
		{
			name: "zeros",
			a:    0, b: 42,
			want: 0,
		},
		{
			name: "zeros with negatives",
			a:    0, b: -42,
			want: -42,
		},
		{
			name: "equal",
			a:    42, b: 42,
			want: 42,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Min(tc.a, tc.b)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestMax(t *testing.T) {
	tcs := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "Basic",
			a:    1, b: 2,
			want: 2,
		},
		{
			name: "negative value",
			a:    -1, b: 2,
			want: 2,
		},
		{
			name: "negative values",
			a:    -1, b: -2,
			want: -1,
		},
		{
			name: "zeros",
			a:    0, b: 42,
			want: 42,
		},
		{
			name: "zeros with negatives",
			a:    0, b: -42,
			want: 0,
		},
		{
			name: "equal",
			a:    42, b: 42,
			want: 42,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := Max(tc.a, tc.b)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
