package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolverActiveWithThreeNeighborsStaysActive(t *testing.T) {
	assert.Subset(t,
		NewSolver([]ActiveCube{AC3(0, 0, 0), AC3(0, 1, 0), AC3(1, 0, 0)}, Neighbors3D).Next().ActiveCubes(),
		[]ActiveCube{AC3(0, 0, 0), AC3(0, 1, 0), AC3(1, 0, 0)},
	)
}
func TestSolverActiveWithTwoNeighborsStaysActive(t *testing.T) {
	newSpace := NewSolver([]ActiveCube{AC3(0, 0, 0), AC3(0, 1, 0), AC3(0, 2, 0)}, Neighbors3D).Next().ActiveCubes()
	assert.Subset(t,
		newSpace,
		[]ActiveCube{AC3(0, 1, 1)},
	)
	assert.NotContains(t, newSpace, AC3(0, 0, 0))
	assert.NotContains(t, newSpace, AC3(0, 2, 0))
}
func TestSolverCubesBecomeActive(t *testing.T) {
	assert.Subset(t,
		NewSolver([]ActiveCube{AC3(0, 0, 0), AC3(0, 1, 0), AC3(1, 0, 0)}, Neighbors3D).Next().ActiveCubes(),
		[]ActiveCube{AC3(0, 0, -1), AC3(1, 1, 0), AC3(0, 0, 1)},
	)
}
