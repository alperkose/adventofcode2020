package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovement(t *testing.T) {
	instructions := []Instruction{
		FromString("F10"),
		FromString("N3"),
		FromString("F7"),
		FromString("R90"),
		FromString("F11"),
	}

	actual := Move(Ship{East, 0, 0, 10, 1}, instructions)
	assert.Equal(t, "E(214, -72)-Wp(4, -10)", actual.String())
	// assert.Equal(t, "S(17, -8)", fmt.Sprintf("%s(%d, %d)", actual.Heading, actual.X, actual.Y))
}
