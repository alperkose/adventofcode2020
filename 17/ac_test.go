package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleFromExercise_Part1(t *testing.T) {
	initialState := `.#.
..#
###`
	cubes := NewParser().Parse(strings.NewReader(initialState))
	s := NewSolver(cubes, Neighbors3D)
	assert.Len(t, s.Next().ActiveCubes(), 11)
	assert.Len(t, s.Next().ActiveCubes(), 21)
	assert.Len(t, s.Next().ActiveCubes(), 38)
	assert.Len(t, s.Next().Next().Next().ActiveCubes(), 112)
}

func TestSampleFromExercise_Part2(t *testing.T) {
	initialState := `.#.
..#
###`
	cubes := NewParser().Parse(strings.NewReader(initialState))
	s := NewSolver(cubes, Neighbors4D)
	assert.Len(t, s.Next().ActiveCubes(), 29)
	assert.Len(t, s.Next().ActiveCubes(), 60)
	assert.Len(t, s.Next().Next().Next().Next().ActiveCubes(), 848)
}
