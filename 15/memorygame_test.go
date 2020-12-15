package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		input   []int
		guessNo int
		answer  int
	}{
		{"sample1", []int{0, 3, 6}, 2020, 436},
		{"sample2", []int{1, 3, 2}, 2020, 1},
		{"sample3", []int{2, 1, 3}, 2020, 10},
		{"sample4", []int{1, 2, 3}, 2020, 27},
		{"sample5", []int{2, 3, 1}, 2020, 78},
		{"sample6", []int{3, 2, 1}, 2020, 438},
		{"sample7", []int{3, 1, 2}, 2020, 1836},
		{"longer sample 1", []int{0, 3, 6}, 30000000, 175594},
		// {"longer sample 2", []int{1, 3, 2}, 30000000, 2578},
		// {"longer sample 3", []int{2, 1, 3}, 30000000, 3544142},
		// {"longer sample 4", []int{1, 2, 3}, 30000000, 261214},
		// {"longer sample 5", []int{2, 3, 1}, 30000000, 6895259},
		// {"longer sample 6", []int{3, 2, 1}, 30000000, 18},
		// {"longer sample 7", []int{3, 1, 2}, 30000000, 362},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.answer, NewGame(tC.input).Guess(tC.guessNo))
		})
	}
}
