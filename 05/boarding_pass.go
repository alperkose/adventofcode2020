package main

import "fmt"

const NumRowLetters = 7
const NumSeatLetters = 3
const TotalLength = NumRowLetters + NumSeatLetters
const RowCount = 128
const ColumnCount = 8

type BoardingPass struct {
	input               string
	row, column, seatID int
}

func (bp BoardingPass) SeatID() int {
	return bp.seatID
}

func NewBoadingPass(input string) (BoardingPass, error) {
	if len(input) != (TotalLength) {
		return BoardingPass{input, -1, -1, -1}, fmt.Errorf("%s does not have 10 letters", input)
	}
	row := 0
	mask := RowCount >> 1
	for i := 0; i < NumRowLetters; i++ {
		switch input[i] {
		case 'F':
		case 'B':
			row += mask
		default:
			return BoardingPass{input, -1, -1, -1}, fmt.Errorf("Boarding pass cannot have %c as row letter", input[i])
		}
		mask = mask >> 1
	}

	col := 0
	mask = ColumnCount >> 1
	for i := NumRowLetters; i < TotalLength; i++ {
		switch input[i] {
		case 'R':
			col += mask
		case 'L':
		default:
			return BoardingPass{input, -1, -1, -1}, fmt.Errorf("Boarding pass cannot have %c as seat letter", input[i])
		}
		mask = mask >> 1
	}
	return BoardingPass{input, row, -1, row*ColumnCount + col}, nil
}
