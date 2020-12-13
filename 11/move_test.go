package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUntilStabilized_Part1(t *testing.T) {
	area := ReadSource(strings.NewReader(
		`L.LL.LL.LL
		LLLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLLL
		L.LLLLLL.L
		L.LLLLL.LL`))
	finalArea, occupants, turns := MoveUntilStabilized(area, Occupy)
	assert.Equal(t, 37, occupants)
	assert.Equal(t, 6, turns)
	assert.Equal(t, []string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.#.L..#..",
		"#L##.##.L#",
		"#.#L.LL.LL",
		"#.#L#L#.##",
		"..L.L.....",
		"#L#L##L#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	}, finalArea)

}

func TestMoveUntilStabilized_Part2(t *testing.T) {
	area := ReadSource(strings.NewReader(
		`L.LL.LL.LL
		LLLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLLL
		L.LLLLLL.L
		L.LLLLL.LL`))
	finalArea, occupants, turns := MoveUntilStabilized(area, OccupyExtendedView)
	assert.Equal(t, 26, occupants)
	assert.Equal(t, 7, turns)
	assert.Equal(t, []string{
		"#.L#.L#.L#",
		"#LLLLLL.LL",
		"L.L.L..#..",
		"##L#.#L.L#",
		"L.L#.LL.L#",
		"#.LLLL#.LL",
		"..#.L.....",
		"LLL###LLL#",
		"#.LLLLL#.L",
		"#.L#LL#.L#",
	}, finalArea)

}
