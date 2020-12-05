package main

import "testing"

func TestCreatingAValidBoardingPass(t *testing.T) {

}

func TestCreatingBoadingPasses(t *testing.T) {
	testCases := []struct {
		input         string
		expectedError bool
		desc          string
	}{
		{"BFFFBBFRLR", false, "Valid boarding pass with all letters"},
		{"BBBBBBBLLL", false, "Valid boarding pass with only B and L"},
		{"FFFFFFFRRR", false, "Valid boarding pass with only F and R"},
		{"AAADFFARRR", true, "Invalid boarding pass with not allowed row letters"},
		{"FBFBFBFCLR", true, "Invalid boarding pass with not allowed column letters"},
		{"FBFBFBLLR", true, "Invalid boarding pass with not enough row letters"},
		{"FBFBFBFBLLR", true, "Invalid boarding pass with too much row letters"},
		{"FBFBFBFLL", true, "Invalid boarding pass with not enough column letters"},
		{"FBFBFBFLLRR", true, "Invalid boarding pass with too much column letters"},
		{"FBFBFBFFLR", true, "Invalid boarding pass with too much row and not enough column letters"},
		{"", true, "Invalid boarding pass as empty string"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := NewBoadingPass(tC.input)
			if (err == nil) == tC.expectedError {
				t.Errorf("err: %#+v\n", err)
			}
		})
	}
}

func TestSeatIDCalculation(t *testing.T) {
	testCases := []struct {
		input  string
		row    int
		col    int
		seatID int
		desc   string
	}{
		{"FFFFFFFLLL", 0, 0, 0, "row 0, column 0"},
		{"BFFFFFFLLL", 64, 0, 512, "row 64, column 0"},
		{"FBFFFFFLLL", 32, 0, 256, "row 32, column 0"},
		{"FFBFFFFLLL", 16, 0, 128, "row 16, column 0"},
		{"FFFBFFFLLL", 8, 0, 64, "row 8, column 0"},
		{"FFFFBFFLLL", 4, 0, 32, "row 4, column 0"},
		{"FFFFFBFLLL", 2, 0, 16, "row 2, column 0"},
		{"FFFFFFBLLL", 1, 0, 8, "row 1, column 0"},
		{"FFFFFFFRLL", 0, 4, 4, "row 0, column 4"},
		{"FFFFFFFLRL", 0, 2, 2, "row 0, column 2"},
		{"FFFFFFFLLR", 0, 1, 1, "row 0, column 1"},
		{"BFFFBBFRRR", 70, 7, 567, "row 70, column 7"},
		{"FFFBBBFRRR", 14, 7, 119, "row 14, column 7"},
		{"BBFFBBFRLL", 102, 4, 820, "row 102, column 4"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bpass, _ := NewBoadingPass(tC.input)
			if bpass.SeatID() != tC.seatID {
				t.Errorf("Wrong seatID for %#+v expected %d -> got %d\n", tC.input, tC.seatID, bpass.SeatID())
			}
		})
	}
}
