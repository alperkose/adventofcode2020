package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActiveCubeString(t *testing.T) {
	assert.Equal(t, "0,0,0,0", AC3(0, 0, 0).String())
	assert.Equal(t, "0,1,2,0", AC3(0, 1, 2).String())
	assert.Equal(t, "4,-3,-5,0", AC3(4, -3, -5).String())
	assert.Equal(t, "4,-3,-5,-1", AC4(4, -3, -5, -1).String())
	assert.Equal(t, "4,-3,-5,4", AC4(4, -3, -5, 4).String())
}

func TestNeighbor3DList(t *testing.T) {
	assert.ElementsMatch(t, []ActiveCube{
		AC3(0, 0, -1), AC3(1, 0, -1), AC3(2, 0, -1),
		AC3(0, 1, -1), AC3(1, 1, -1), AC3(2, 1, -1),
		AC3(0, 2, -1), AC3(1, 2, -1), AC3(2, 2, -1),
		AC3(0, 0, 0), AC3(1, 0, 0), AC3(2, 0, 0),
		AC3(0, 1, 0), AC3(2, 1, 0),
		AC3(0, 2, 0), AC3(1, 2, 0), AC3(2, 2, 0),
		AC3(0, 0, 1), AC3(1, 0, 1), AC3(2, 0, 1),
		AC3(0, 1, 1), AC3(1, 1, 1), AC3(2, 1, 1),
		AC3(0, 2, 1), AC3(1, 2, 1), AC3(2, 2, 1),
	},
		Neighbors3D(AC3(1, 1, 0)),
	)
}
func TestNeighbor4DList(t *testing.T) {
	assert.Len(t, Neighbors4D(ActiveCube{1, 1, 0, 1}), 80)
}
