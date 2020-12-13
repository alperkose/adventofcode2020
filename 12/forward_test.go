package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func DisabledTestForwardInstruction(t *testing.T) {
	testCases := []struct {
		desc             string
		heading          Heading
		instr            string
		expectedPosition string
	}{
		{"Forward east", East, "F10", "(10, 0)"},
		{"Forward north", North, "F9", "(0, 9)"},
		{"Forward north", West, "F8", "(-8, 0)"},
		{"Forward north", South, "F7", "(0, -7)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{tC.heading, 0, 0, 0, 0}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expectedPosition, fmt.Sprintf("(%d, %d)", finalShip.X, finalShip.Y))
		})
	}
}
func TestForwardByWaypointInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		instr    string
		wpx, wpy int
		expected string
	}{
		{"Forward east", "F10", 10, 1, "E(100, 10)-Wp(10, 1)"},
		{"Forward north", "F9", -1, 10, "E(-9, 90)-Wp(-1, 10)"},
		{"Forward north", "F8", -3, -4, "E(-24, -32)-Wp(-3, -4)"},
		{"Forward north", "F7", 2, -7, "E(14, -49)-Wp(2, -7)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{East, 0, 0, tC.wpx, tC.wpy}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.String())
		})
	}
}
