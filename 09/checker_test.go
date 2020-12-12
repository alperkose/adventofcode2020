package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksWithoutIterating(t *testing.T) {
	assert.NoError(t, VerifyWithPreamble([]int{1, 2, 3, 4, 5, 6}, 5))
	assert.NoError(t, VerifyWithPreamble([]int{1, 2, 3, 4, 5, 7}, 5))
	assert.EqualError(t, VerifyWithPreamble([]int{1, 2, 3, 4, 5, 10}, 5), "Doesn't add up to 10")
	assert.EqualError(t, VerifyWithPreamble([]int{5, 3, 1, 8, 13, 12}, 5), "Doesn't add up to 12")
	assert.NoError(t, VerifyWithPreamble([]int{5, 3, 2, 8, 13, 7}, 5))
	assert.EqualError(t, VerifyWithPreamble([]int{5, 3, 2, 8, 10, 7}, 4), "Doesn't add up to 7")
}

func TestChecksWithIterating(t *testing.T) {
	assert.NoError(t, VerifyWithPreamble([]int{1, 2, 3, 4, 5, 7, 9, 11}, 4))
	assert.EqualError(t, VerifyWithPreamble([]int{1, 2, 3, 4, 5, 9, 12, 20}, 4), "Doesn't add up to 20")
	assert.EqualError(t, VerifyWithPreamble([]int{8, 3, 1, 5, 13, 21}, 4), "Doesn't add up to 21")
}

func TestContiguousSum(t *testing.T) {
	assert.Equal(t, 62, FindContiguousSum(127, []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182}))
}
