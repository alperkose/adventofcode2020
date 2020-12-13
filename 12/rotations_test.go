package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaypointLeftRotationInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		instr    string
		expected string
	}{
		{"Rotate  90 degrees left", "L90", "E(0, 0)-Wp(-1, 10)"},
		{"Rotate 180 degrees left", "L180", "E(0, 0)-Wp(-10, -1)"},
		{"Rotate 270 degrees left", "L270", "E(0, 0)-Wp(1, -10)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{East, 0, 0, 10, 1}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.String())
		})
	}
}

func TestWaypointRightRotationInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		instr    string
		expected string
	}{
		{"Rotate  90 degrees right", "R90", "E(0, 0)-Wp(1, -10)"},
		{"Rotate 180 degrees right", "R180", "E(0, 0)-Wp(-10, -1)"},
		{"Rotate 270 degrees right", "R270", "E(0, 0)-Wp(-1, 10)"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{East, 0, 0, 10, 1}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.String())
		})
	}
}
func DisabledTestLeftRotationInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		heading  Heading
		instr    string
		expected Heading
	}{
		{"Rotate  90 degrees left from east", East, "L90", North},
		{"Rotate 180 degrees left from east", East, "L180", West},
		{"Rotate 270 degrees left from east", East, "L270", South},
		{"Rotate  90 degrees left from north", North, "L90", West},
		{"Rotate 180 degrees left from north", North, "L180", South},
		{"Rotate 270 degrees left from north", North, "L270", East},
		{"Rotate  90 degrees left from west", West, "L90", South},
		{"Rotate 180 degrees left from west", West, "L180", East},
		{"Rotate 270 degrees left from west", West, "L270", North},
		{"Rotate  90 degrees left from south", South, "L90", East},
		{"Rotate 180 degrees left from south", South, "L180", North},
		{"Rotate 270 degrees left from south", South, "L270", West},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{tC.heading, 0, 0, 0, 0}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.Heading)
		})
	}
}

func DisabledTestRightRotationInstruction(t *testing.T) {
	testCases := []struct {
		desc     string
		heading  Heading
		instr    string
		expected Heading
	}{
		{"Rotate  90 degrees right from east", East, "R90", South},
		{"Rotate 180 degrees right from east", East, "R180", West},
		{"Rotate 270 degrees right from east", East, "R270", North},
		{"Rotate  90 degrees right from north", North, "R90", East},
		{"Rotate 180 degrees right from north", North, "R180", South},
		{"Rotate 270 degrees right from north", North, "R270", West},
		{"Rotate  90 degrees right from west", West, "R90", North},
		{"Rotate 180 degrees right from west", West, "R180", East},
		{"Rotate 270 degrees right from west", West, "R270", South},
		{"Rotate  90 degrees right from south", South, "R90", West},
		{"Rotate 180 degrees right from south", South, "R180", North},
		{"Rotate 270 degrees right from south", South, "R270", East},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ship := Ship{tC.heading, 0, 0, 0, 0}
			instruction := FromString(tC.instr)
			finalShip := instruction.Apply(ship)
			assert.Equal(t, tC.expected, finalShip.Heading)
		})
	}
}
