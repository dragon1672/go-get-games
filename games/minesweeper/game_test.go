package minesweeper

import (
	"testing"

	"github.com/dragon1672/go-collections/vector"
	"github.com/pellared/fluentassert/f"
)

func list2Map(vals ...ChangeEventData) map[ChangeEventData]bool {
	ret := make(map[ChangeEventData]bool)
	for _, val := range vals {
		ret[val] = true
	}
	return ret
}

func TestGameCalcNum(t *testing.T) {
	tcs := []struct {
		name  string
		game  *Game
		input vector.IntVec2
		want  int
	}{
		{
			name: "empty",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				"   \n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  0,
		},
		{
			name: "out of bounds",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				"   \n" +
				"   ")),
			input: vector.Of(-1, -1),
			want:  0,
		},
		{
			name: "direct on",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				" * \n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  0,
		},
		{
			name: "1 above",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				" * \n" +
				"   \n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  1,
		},
		{
			name: "1 side",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				"  *\n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  1,
		},
		{
			name: "1 corner",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"  *\n" +
				"   \n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  1,
		},
		{
			name: "multiple each direction",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"***\n" +
				"*  \n" +
				"   ")),
			input: vector.Of(1, 1),
			want:  4,
		},
		{
			name: "multiple each direction pt 2",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				"  *\n" +
				"***")),
			input: vector.Of(1, 1),
			want:  4,
		},
		{
			name: "all around",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"***\n" +
				"* *\n" +
				"***")),
			input: vector.Of(1, 1),
			want:  8,
		},
		{
			name: "solid",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"***\n" +
				"***\n" +
				"***")),
			input: vector.Of(1, 1),
			want:  8,
		},
		{
			name: "out of bounds touching",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"***\n" +
				"***\n" +
				"***")),
			input: vector.Of(3, 1),
			want:  3,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.game.calcNum(tc.input)
			f.Assert(t, got).Eq(tc.want, "should return proper value")
		})
	}
}

func TestReveal(t *testing.T) {
	tcs := []struct {
		name  string
		game  *Game
		input vector.IntVec2
		want  []ChangeEventData
	}{
		{
			name: "empty",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"   \n" +
				"   \n" +
				"   ")),
			input: vector.Of(1, 1),
			want: []ChangeEventData{
				{Pos: vector.Of(0, 0), Val: CellN0}, {Pos: vector.Of(1, 0), Val: CellN0}, {Pos: vector.Of(2, 0), Val: CellN0},
				{Pos: vector.Of(0, 1), Val: CellN0}, {Pos: vector.Of(1, 1), Val: CellN0}, {Pos: vector.Of(2, 1), Val: CellN0},
				{Pos: vector.Of(0, 2), Val: CellN0}, {Pos: vector.Of(1, 2), Val: CellN0}, {Pos: vector.Of(2, 2), Val: CellN0},
			},
		},
		{
			name: "bomb reveal",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"111\n" +
				"1*1\n" +
				"111")),
			input: vector.Of(1, 1),
			want: []ChangeEventData{
				{vector.Of(1, 1), CellBomb},
			},
		},
		{
			name: "single reveal",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"*1 \n" +
				"11 \n" +
				"   ")),
			input: vector.Of(1, 1),
			want: []ChangeEventData{
				{vector.Of(1, 1), CellN1},
			},
		},
		{
			name: "multi reveal",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"*1 \n" +
				"11 \n" +
				"   ")),
			input: vector.Of(2, 2),
			want: []ChangeEventData{
				{Pos: vector.Of(1, 0), Val: CellN1}, {Pos: vector.Of(2, 0), Val: CellN0},
				{Pos: vector.Of(0, 1), Val: CellN1}, {Pos: vector.Of(1, 1), Val: CellN1}, {Pos: vector.Of(2, 1), Val: CellN0},
				{Pos: vector.Of(0, 2), Val: CellN0}, {Pos: vector.Of(1, 2), Val: CellN0}, {Pos: vector.Of(2, 2), Val: CellN0},
			},
		},
		{
			name: "multi reveal hides",
			game: MakeFromGenerator(MakeGameGenFromString("" +
				"*1 \n" +
				"22 \n" +
				"*1 ")),
			input: vector.Of(2, 2),
			want: []ChangeEventData{
				{Pos: vector.Of(1, 0), Val: CellN1}, {Pos: vector.Of(2, 0), Val: CellN0},
				{Pos: vector.Of(1, 1), Val: CellN2}, {Pos: vector.Of(2, 1), Val: CellN0},
				{Pos: vector.Of(1, 2), Val: CellN1}, {Pos: vector.Of(2, 2), Val: CellN0},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.game.Reveal(tc.input)
			f.Assert(t, list2Map(got...)).Eq(list2Map(tc.want...), "should return proper value")
		})
	}
}
