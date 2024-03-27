package minesweeper

import (
	"github.com/dragon162/go-get-games/games/minesweeper/common/sliceutls"
	"github.com/dragon1672/go-collections/vector"
	"strings"
)

//MineSweeper.create(10, 10, 10); // 10% bombs

// Beginner has a total of ten mines and the board size is either 8 × 8, 9 × 9, or 10 × 10
// 8x8:   15.625% bombs
// 9x9:   12.346% bombs
// 10x10: 10% bombs
//goland:noinspection GoUnusedGlobalVariable
var (
	BeginnerGame       = &GameGenerator{Width: 10, Height: 10, Gen: &BombCountGen{BombCount: 10}}
	BeginnerDifficulty = &DifficultyBombGen{BombPercent: .1}
)

// Intermediate has 40 mines and also varies in size between 13 × 15 and 16 × 16
// 13x15: 20.512% bombs
// 16x16: 15.625% bombs
//goland:noinspection GoUnusedGlobalVariable
var (
	IntermediateGame       = &GameGenerator{Width: 16, Height: 16, Gen: &BombCountGen{BombCount: 40}}
	IntermediateDifficulty = &DifficultyBombGen{BombPercent: .16}
)

// Expert has 99 mines and is always 16 × 30 (or 30 × 16) : 20.625% bomb
//goland:noinspection GoUnusedGlobalVariable
var (
	ExpertGame       = &GameGenerator{Width: 16, Height: 30, Gen: &BombCountGen{BombCount: 99}}
	ExpertDifficulty = &DifficultyBombGen{BombPercent: .2}
)

//goland:noinspection GoUnusedGlobalVariable
var (
	InsaneGame       = &GameGenerator{Width: 50, Height: 30, Gen: InsaneDifficulty}
	InsaneDifficulty = &DifficultyBombGen{BombPercent: .3}
)

type StaticBombGen struct {
	GameStateGenerator
	Bombs map[vector.IntVec2]bool
}

func (s *StaticBombGen) GenerateBombs(_, _ int, _ ...vector.IntVec2) map[vector.IntVec2]bool {
	// Create a copy to avoid any accidental shady business
	ret := make(map[vector.IntVec2]bool, len(s.Bombs))
	for key, value := range s.Bombs {
		ret[key] = value
	}
	return ret
}

type BombCountGen struct {
	GameStateGenerator
	BombCount int
}

func (b *BombCountGen) GenerateBombs(width, height int, discouragedPositions ...vector.IntVec2) map[vector.IntVec2]bool {
	return generateNumBombs(width, height, b.BombCount, discouragedPositions...)
}

type DifficultyBombGen struct {
	GameStateGenerator
	BombPercent float64
}

func (s *DifficultyBombGen) GenerateBombs(width, height int, discouragedPositions ...vector.IntVec2) map[vector.IntVec2]bool {
	totalSize := width * height
	bombsToGenerate := int(float64(totalSize) * s.BombPercent)
	return generateNumBombs(width, height, bombsToGenerate, discouragedPositions...)
}

// MakeGameGenFromString creates a static minesweeper from a given string
// a `*` will be marked as a bomb and any other character ignored
// EG: 3x3 with 2 bombs
// ` * \n`
// `   \n`
// `*  \n`
func MakeGameGenFromString(s string) *GameGenerator {
	lines := strings.Split(s, "\n")
	height := len(lines)
	var width int
	var bombs []vector.IntVec2
	for y, line := range lines {
		if len(line) > width {
			width = len(line)
		}
		for x, val := range line {
			if val == '*' {
				bombs = append(bombs, vector.Of(x, y))
			}
		}
	}
	return &GameGenerator{
		Width:  width,
		Height: height,
		Gen:    &StaticBombGen{Bombs: sliceutls.List2Map(bombs...)},
	}
}

// generateNumBombs will create a bomb list and attempt to generate number of bombs, but will prioritize discouraged positions.
func generateNumBombs(width, height int, bombs int, discouragedPositions ...vector.IntVec2) map[vector.IntVec2]bool {
	var possiblePos []vector.IntVec2
	toFilter := sliceutls.List2Map(discouragedPositions...)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pos := vector.Of(x, y)
			if restricted := toFilter[pos]; !restricted {
				possiblePos = append(possiblePos, pos)
			}
		}
	}
	// shuffle positions
	sliceutls.Shuffle(possiblePos)
	return sliceutls.List2Map[vector.IntVec2](sliceutls.Truncate(possiblePos, bombs)...)
}
