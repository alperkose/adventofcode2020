package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingSample(t *testing.T) {
	initialState := `.#.
..#
###`
	cubes := NewParser().Parse(strings.NewReader(initialState))
	assert.Len(t, cubes, 5)
	assert.ElementsMatch(t, []ActiveCube{AC3(1, 0, 0), AC3(2, 1, 0), AC3(0, 2, 0), AC3(1, 2, 0), AC3(2, 2, 0)}, cubes)
}
