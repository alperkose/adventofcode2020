package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingCards(t *testing.T) {
	testCases := []struct {
		desc     string
		solver   *Solver
		expected []int
	}{
		{
			"Single turn player 1 wins",
			&Solver{[]int{3}, []int{2}},
			[]int{3, 2},
		}, {
			"Single turn player 2 wins",
			&Solver{[]int{1}, []int{5}},
			[]int{5, 1},
		}, {
			"Two turns",
			&Solver{[]int{1, 10}, []int{5, 9}},
			[]int{10, 5, 9, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, tC.solver.Combat())
		})
	}
}
