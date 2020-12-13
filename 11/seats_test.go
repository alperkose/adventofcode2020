package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstRound(t *testing.T) {
	area := []string{".LL.LLL..", "LL.LL.LL."}

	occupiedArea := Occupy(area)
	assert.Len(t, occupiedArea, len(area))
	assert.Equal(t, occupiedArea[0], ".##.###..")
	assert.Equal(t, occupiedArea[1], "##.##.##.")
}

func TestOccupiedSeatStaysOccupied(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  []string
		expected []string
	}{
		{"3 adjacent seats taken", []string{"...", "###", ".#."}, []string{"...", "###", ".#."}},
		{"2 adjacent seats taken", []string{".#.", "#.#", ".#."}, []string{".#.", "#.#", ".#."}},
		{"1 adjacent seat taken", []string{"##.", "...", ".##"}, []string{"##.", "...", ".##"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Occupy(tC.initial)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
func TestOccupiedSeatAreVacated(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  []string
		expected []string
	}{
		{"4 adjacent seats taken", []string{".#.", "###", ".#."}, []string{".#.", "#L#", ".#."}},
		{"5 adjacent seats taken", []string{"##.", "###", ".#."}, []string{"#L.", "LL#", ".#."}},
		{"6 adjacent seat taken", []string{"##.", "###", ".##"}, []string{"#L.", "LLL", ".L#"}},
		{"7 adjacent seat taken", []string{"###", "###", ".##"}, []string{"#L#", "LLL", ".L#"}},
		{"8 adjacent seat taken", []string{"###", "###", "###"}, []string{"#L#", "LLL", "#L#"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Occupy(tC.initial)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
func TestEmptySeatsAreOccupied(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  []string
		expected []string
	}{
		{"no adjacent seats", []string{"L.L", "...", "L.L"}, []string{"#.#", "...", "#.#"}},
		{"1 adjacent empty seat", []string{"LL.", "...", ".LL"}, []string{"##.", "...", ".##"}},
		{"2 adjacent empty seats", []string{"LLL", "...", "LLL"}, []string{"###", "...", "###"}},
		{"3 and 6 adjacent seats taken", []string{"LLL", ".L.", "LLL"}, []string{"###", ".#.", "###"}},
		{"4, 5 and 7 adjacent seats taken", []string{"LLL", "LL.", "LLL"}, []string{"###", "##.", "###"}},
		{"all adjacent seats taken", []string{"LLL", "LLL", "LLL"}, []string{"###", "###", "###"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Occupy(tC.initial)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
func TestEmptySeatsRemainVacant(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  []string
		expected []string
	}{
		{"1 adjacent occupant", []string{"L#.", "...", ".L#"}, []string{"L#.", "...", ".L#"}},
		{"2 adjacent occupants", []string{"L#.", "#.#", ".#L"}, []string{"L#.", "#.#", ".#L"}},
		{"3 adjacent occupants", []string{"L#L", "###", "..."}, []string{"L#L", "###", "..."}},
		{"4 adjacent occupants", []string{"###", "L.L", "###"}, []string{"###", "L.L", "###"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Occupy(tC.initial)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestExtendedView(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  []string
		expected []string
	}{
		{
			"top left diagonal, one sees vacant one sees occupied",
			[]string{"L....", ".....", "..L..", ".....", "....#"},
			[]string{"#....", ".....", "..L..", ".....", "....#"},
		}, {
			"top right diagonal, one sees vacant one sees occupied",
			[]string{"....L", ".....", "..L..", ".....", "#...."},
			[]string{"....#", ".....", "..L..", ".....", "#...."},
		}, {
			"horizontal, one sees vacant one sees occupied",
			[]string{"L.L.#"},
			[]string{"#.L.#"},
		}, {
			"vertical, one sees vacant one sees occupied",
			[]string{"L", ".", "L", ".", "#"},
			[]string{"#", ".", "L", ".", "#"},
		}, {
			"occupied seat remains occupied with 4 or less occupied seats in sight",
			[]string{"#.##L", ".#...", ".#..#", ".....", "#L.#."},
			[]string{"#.##L", ".#...", ".#..#", ".....", "#L.#."},
		}, {
			"occupied seat vacated with 5 or 6 occupied seats in sight",
			[]string{"#.##L", ".#..#", "##..#", ".....", "#L.##"},
			[]string{"#.##L", ".L..#", "#L..#", ".....", "#L.##"},
		}, {
			"occupied seat vacated with 7 occupied seats in sight",
			[]string{".#.#.", ".....", "##..#", ".....", ".#.#."},
			[]string{".#.#.", ".....", "#L..#", ".....", ".#.#."},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := OccupyExtendedView(tC.initial)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
