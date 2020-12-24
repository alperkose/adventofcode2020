package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSampleAfter10Moves(t *testing.T) {
	result := Solve([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10)
	assert.Equal(t, []int{9, 2, 6, 5, 8, 3, 7, 4}, result)
}
func TestFromSampleAfter100Moves(t *testing.T) {
	result := Solve([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 100)
	assert.Equal(t, []int{6, 7, 3, 8, 4, 5, 2, 9}, result)
}
