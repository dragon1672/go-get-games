package sudoku

import (
	"errors"
	"fmt"
)

type Move uint16

const (
	Width       = 9
	Height      = 9
	UNSET  Move = 99
	S1     Move = 1
	S2     Move = 2
	S3     Move = 3
	S4     Move = 4
	S5     Move = 5
	S6     Move = 6
	S7     Move = 7
	S8     Move = 8
	S9     Move = 9
)

var (
	Zones = struct {
		// each row and their indexes
		RowIndexes [][]int
		// each col and their indexes
		ColIndexes [][]int
		// each square and their indexes
		SquareIndexes [][]int
	}{
		RowIndexes: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{9, 10, 11, 12, 13, 14, 15, 16, 17},
			{18, 19, 20, 21, 22, 23, 24, 25, 26},
			{27, 28, 29, 30, 31, 32, 33, 34, 35},
			{36, 37, 38, 39, 40, 41, 42, 43, 44},
			{45, 46, 47, 48, 49, 50, 51, 52, 53},
			{54, 55, 56, 57, 58, 59, 60, 61, 62},
			{63, 64, 65, 66, 67, 68, 69, 70, 71},
			{72, 73, 74, 75, 76, 77, 78, 79, 80},
		},
		ColIndexes: [][]int{
			{0, 9, 18, 27, 36, 45, 54, 63, 72},
			{1, 10, 19, 28, 37, 46, 55, 64, 73},
			{2, 11, 20, 29, 38, 47, 56, 65, 74},
			{3, 12, 21, 30, 39, 48, 57, 66, 75},
			{4, 13, 22, 31, 40, 49, 58, 67, 76},
			{5, 14, 23, 32, 41, 50, 59, 68, 77},
			{6, 15, 24, 33, 42, 51, 60, 69, 78},
			{7, 16, 25, 34, 43, 52, 61, 70, 79},
			{8, 17, 26, 35, 44, 53, 62, 71, 80},
		},
		SquareIndexes: [][]int{
			{0, 1, 2, 9, 10, 11, 18, 19, 20},
			{3, 4, 5, 12, 13, 14, 21, 22, 23},
			{6, 7, 8, 15, 16, 17, 24, 25, 26},
			{27, 28, 29, 36, 37, 38, 45, 46, 47},
			{30, 31, 32, 39, 40, 41, 48, 49, 50},
			{33, 34, 35, 42, 43, 44, 51, 52, 53},
			{54, 55, 56, 63, 64, 65, 72, 73, 74},
			{57, 58, 59, 66, 67, 68, 75, 76, 77},
			{60, 61, 62, 69, 70, 71, 78, 79, 80},
		},
	}
	AllZones = combineSlices(Zones.SquareIndexes, Zones.RowIndexes, Zones.SquareIndexes)
)

func combineSlices[T any](a ...[]T) (ret []T) {
	for _, v := range a {
		ret = append(ret, v...)
	}
	return ret
}

func (s Move) Num() uint16 {
	return uint16(s)
}

func (s Move) String() string {
	return string([]byte{byte(s.Rune())})
}

func (s Move) Rune() rune {
	if s == UNSET {
		return '.'
	}
	return '0' + int32(s.Num())
}

func validate(board ReadOnlyBoard, combine bool) (err error) {
	validOptions := map[Move]struct{}{UNSET: {}, S1: {}, S2: {}, S3: {}, S4: {}, S5: {}, S6: {}, S7: {}, S8: {}, S9: {}}
	for _, moveSet := range AllZones {
		// set Options so that their # correlates to their value
		// these options will be replaced with UNSET as they are "used" to ensure that numbers are not duplicated
		options := []Move{UNSET, S1, S2, S3, S4, S5, S6, S7, S8, S9}
		for _, i := range moveSet {
			v := board.GetI(i)
			if _, ok := validOptions[v]; !ok {
				err = errors.Join(err, fmt.Errorf("entry %v not a valid option: %v", v, validOptions))
				if !combine {
					return
				}
			}
			if v != UNSET {
				if options[v.Num()] == UNSET {
					err = errors.Join(err, fmt.Errorf("discovered duplicate entry %v in moveset %v", v, moveSet))
					if !combine {
						return
					}
				}
				options[v.Num()] = UNSET
			}
		}
	}
	return
}

func ValidateAll(board ReadOnlyBoard) error {
	return validate(board, true)
}

func Validate(board ReadOnlyBoard) error {
	return validate(board, false)
}

func TextPrintBoard(board ReadOnlyBoard) string {
	return fmt.Sprintf(""+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"----- + ----- + -----\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"----- + ----- + -----\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"%v %v %v | %v %v %v | %v %v %v\n"+
		"",
		board.GetI(0), board.GetI(1), board.GetI(2), board.GetI(3), board.GetI(4), board.GetI(5), board.GetI(6), board.GetI(7), board.GetI(8),
		board.GetI(9), board.GetI(10), board.GetI(11), board.GetI(12), board.GetI(13), board.GetI(14), board.GetI(15), board.GetI(16), board.GetI(17),
		board.GetI(18), board.GetI(19), board.GetI(20), board.GetI(21), board.GetI(22), board.GetI(23), board.GetI(24), board.GetI(25), board.GetI(26),
		board.GetI(27), board.GetI(28), board.GetI(29), board.GetI(30), board.GetI(31), board.GetI(32), board.GetI(33), board.GetI(34), board.GetI(35),
		board.GetI(36), board.GetI(37), board.GetI(38), board.GetI(39), board.GetI(40), board.GetI(41), board.GetI(42), board.GetI(43), board.GetI(44),
		board.GetI(45), board.GetI(46), board.GetI(47), board.GetI(48), board.GetI(49), board.GetI(50), board.GetI(51), board.GetI(52), board.GetI(53),
		board.GetI(54), board.GetI(55), board.GetI(56), board.GetI(57), board.GetI(58), board.GetI(59), board.GetI(60), board.GetI(61), board.GetI(62),
		board.GetI(63), board.GetI(64), board.GetI(65), board.GetI(66), board.GetI(67), board.GetI(68), board.GetI(69), board.GetI(70), board.GetI(71),
		board.GetI(72), board.GetI(73), board.GetI(74), board.GetI(75), board.GetI(76), board.GetI(77), board.GetI(78), board.GetI(79), board.GetI(80))
}
