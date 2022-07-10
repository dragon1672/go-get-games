package vector

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestCreate(t *testing.T) {
	got := Of(1, 2)
	f.Assert(t, got.X).Eq(1, "should return proper value")
	f.Assert(t, got.Y).Eq(2, "should return proper value")
}

func TestAdd(t *testing.T) {
	tcs := []struct {
		name string
		a, b IntVec2
		want IntVec2
	}{
		{
			name: "Basic",
			a:    Of(1, 2), b: Of(4, 5),
			want: Of(5, 7),
		},
		{
			name: "negative values",
			a:    Of(-1, -2), b: Of(4, 5),
			want: Of(3, 3),
		},
		{
			name: "zeros",
			a:    Of(0, 0), b: Of(4, 5),
			want: Of(4, 5),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.Add(tc.b)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestSub(t *testing.T) {
	tcs := []struct {
		name string
		a, b IntVec2
		want IntVec2
	}{
		{
			name: "Basic",
			a:    Of(1, 2), b: Of(4, 5),
			want: Of(-3, -3),
		},
		{
			name: "negative values",
			a:    Of(-1, -2), b: Of(4, 5),
			want: Of(-5, -7),
		},
		{
			name: "zeros",
			a:    Of(0, 0), b: Of(4, 5),
			want: Of(-4, -5),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.Sub(tc.b)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestMul(t *testing.T) {
	tcs := []struct {
		name string
		a    IntVec2
		s    float64
		want IntVec2
	}{
		{
			name: "Basic",
			a:    Of(1, 2), s: 2,
			want: Of(2, 4),
		},
		{
			name: "negative values",
			a:    Of(1, 2), s: -2,
			want: Of(-2, -4),
		},
		{
			name: "zeros",
			a:    Of(1, 2), s: 0,
			want: Of(0, 0),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.Mul(tc.s)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
