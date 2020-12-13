package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func DisabledTestMoveInDirectionInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		heading  Heading
		instr    string
		expected string
	}{
		{"Move east", North, "E10", "N(10, 0)"},
		{"Move north", East, "N9", "E(0, 9)"},
		{"Move west", South, "W8", "S(-8, 0)"},
		{"Move south", West, "S7", "W(0, -7)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{tC.heading, 0, 0, 0, 0}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, fmt.Sprintf("%s(%d, %d)", finalShip.Heading, finalShip.X, finalShip.Y))
		})
	}
}
func TestMoveInDirectionInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		heading  Heading
		instr    string
		expected string
	}{
		{"Move east", North, "E10", "N(0, 0)-Wp(20, 1)"},
		{"Move north", East, "N9", "E(0, 0)-Wp(10, 10)"},
		{"Move west", South, "W8", "S(0, 0)-Wp(2, 1)"},
		{"Move south", West, "S7", "W(0, 0)-Wp(10, -6)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{tC.heading, 0, 0, 10, 1}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.String())
		})
	}
}
