package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoltGroups(t *testing.T) {
	testCases := []struct {
		input                         []int
		oneJolt, twoJolts, threeJolts int
		desc                          string
	}{
		{[]int{1}, 1, 0, 1, "single jolt"},
		{[]int{2}, 0, 1, 1, "single but 2 step above plug"},
		{[]int{2, 3}, 1, 1, 1, "smallest difference"},
		{[]int{1, 3, 5}, 1, 2, 1, "steps of two"},
		{[]int{1, 4, 5, 7, 9}, 2, 2, 2, "mixed but balanced"},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 7, 0, 5, "sample from puzzle"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			oneJ, twoJ, threeJ := JoltDifferenceGroups(tC.input)
			assert.Equal(t, tC.oneJolt, oneJ)
			assert.Equal(t, tC.twoJolts, twoJ)
			assert.Equal(t, tC.threeJolts, threeJ)
		})
	}
}

func TestJoltCombinations(t *testing.T) {
	testCases := []struct {
		input        []int
		combinations int
		desc         string
	}{
		{[]int{2, 3}, 2, "combinations with 0"},
		{[]int{3, 4}, 1, "no combination with 0"},
		{[]int{2, 4, 6}, 1, "steps of two"},
		{[]int{1, 4, 5, 8}, 1, "few combinations"},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 8, "sample from puzzle"},
		{[]int{1, 2, 3}, 4, "steps of one"},
		{[]int{1, 2, 3, 4}, 7, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5}, 13, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6}, 24, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 44, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 81, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 149, "sample from puzzle"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.combinations, JoltCombinations(tC.input, validCombinations))
		})
	}
}
func TestJoltCombinationsWithTribonacci(t *testing.T) {
	testCases := []struct {
		input        []int
		combinations int
		desc         string
	}{
		{[]int{1, 2, 3}, 4, "steps of one"},
		{[]int{1, 2, 3, 4}, 7, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5}, 13, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6}, 24, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 44, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 81, "sample from puzzle"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 149, "sample from puzzle"},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 8, "sample from puzzle"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.combinations, JoltCombinations(tC.input, tribonacci))
		})
	}
}

func BenchmarkCombinatorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoltCombinations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, validCombinations)
		JoltCombinations([]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, validCombinations)
	}

}
func BenchmarkTribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoltCombinations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tribonacci)
		JoltCombinations([]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, tribonacci)
	}

}
