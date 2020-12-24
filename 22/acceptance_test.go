package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingCombat(t *testing.T) {
	var input = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
	solver := Parse(strings.NewReader(input))
	winnersCards := solver.Combat()
	assert.Equal(t, []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}, winnersCards)
}

func TestPlayingRecursiveCombat(t *testing.T) {
	var input = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
	solver := Parse(strings.NewReader(input))
	winnersCards := solver.RecursiveCombat()
	assert.Equal(t, []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}, winnersCards)
}
func XTestPlayingRecursiveCombatWithInfiniteLoopPrevention(t *testing.T) {
	var input = `Player 1:
2
43
19

Player 2:
3
2
29
14`
	solver := Parse(strings.NewReader(input))
	winnersCards := solver.RecursiveCombat()
	assert.Equal(t, []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}, winnersCards)
}
