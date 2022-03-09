package grid

import (
	"github.com/pellared/fluentassert/f"
	"go-get-games/games/common/vector"
	"testing"
)

func TestMake(t *testing.T) {
	tcs := []struct {
		name          string
		width, height int
		wantErr       bool
	}{
		{
			name:  "Basic",
			width: 42, height: 42,
			wantErr: false,
		},
		{
			name:  "zero width",
			width: 0, height: 42,
			wantErr: true,
		},
		{
			name:  "zero height",
			width: 42, height: 0,
			wantErr: true,
		},
		{
			name:  "negative width",
			width: -2, height: 42,
			wantErr: true,
		},
		{
			name:  "negative height",
			width: 42, height: -2,
			wantErr: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, gotErr := MakeSimpleGrid[int](tc.width, tc.height)
			f.Assert(t, gotErr != nil).Eq(tc.wantErr, "should return proper error")
		})
	}
}

func TestGetAndSet(t *testing.T) {
	tcs := []struct {
		name    string
		grid    func() (*SimpleGrid[string], error)
		pos     vector.IntVec2
		want    string
		wantErr bool
	}{
		{
			name: "Basic",
			grid: func() (*SimpleGrid[string], error) {
				grid, err := MakeSimpleGrid[string](5, 5)
				if err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(1, 1), "da value"); err != nil {
					return nil, err
				}
				return grid, nil
			},
			pos:     vector.Of(1, 1),
			want:    "da value",
			wantErr: false,
		},
		{
			name: "multiple entries",
			grid: func() (*SimpleGrid[string], error) {
				grid, err := MakeSimpleGrid[string](5, 5)
				if err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(1, 1), "da value"); err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(2, 2), "some other value"); err != nil {
					return nil, err
				}
				return grid, nil
			},
			pos:     vector.Of(1, 1),
			want:    "da value",
			wantErr: false,
		},
		{
			name: "overriding value",
			grid: func() (*SimpleGrid[string], error) {
				grid, err := MakeSimpleGrid[string](5, 5)
				if err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(1, 1), "OG value"); err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(1, 1), "da value now overridden"); err != nil {
					return nil, err
				}
				return grid, nil
			},
			pos:     vector.Of(1, 1),
			want:    "da value now overridden",
			wantErr: false,
		},
		{
			name: "default value",
			grid: func() (*SimpleGrid[string], error) {
				grid, err := MakeSimpleGrid[string](5, 5)
				if err != nil {
					return nil, err
				}
				if err := grid.Set(vector.Of(2, 2), "random entry"); err != nil {
					return nil, err
				}
				return grid, nil
			},
			pos:     vector.Of(1, 1),
			want:    "",
			wantErr: false,
		},
		{
			name: "invalidPos",
			grid: func() (*SimpleGrid[string], error) {
				grid, err := MakeSimpleGrid[string](5, 5)
				if err != nil {
					return nil, err
				}
				return grid, nil
			},
			pos:     vector.Of(-1, -1),
			want:    "",
			wantErr: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			grid, err := tc.grid()
			if err != nil {
				t.Fatalf("Set setup failed with error %v", err)
			}
			got, gotErr := grid.Get(tc.pos)
			f.Assert(t, gotErr != nil).Eq(tc.wantErr, "should return proper error")
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}
