package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	assert.Equal(t, []int{3, 5, 2, 4}, Solve([]int{3, 2, 4, 1, 5}, 1))
	assert.Equal(t, []int{5, 2, 4, 3}, Solve([]int{3, 5, 2, 4, 1}, 1))

}

var setup sync.Once
var benchmarkInput []int

func XTestSolverWithIncreasingValues(t *testing.T) {

	setup.Do(func() {
		input := []int{3, 5, 2, 4, 1}
		benchmarkInput = make([]int, 1000)
		var i int
		for i = 0; i < len(input); i++ {
			benchmarkInput[i] = input[i]
		}
		for ; i < len(benchmarkInput); i++ {
			benchmarkInput[i] = i + 1
		}
	})
	Solve(benchmarkInput, 10000)
	t.FailNow()
}

func BenchmarkSolverWithIncreasingValues(b *testing.B) {

	setup.Do(func() {
		input := []int{3, 5, 2, 4, 1}
		benchmarkInput = make([]int, 1000)
		var i int
		for i = 0; i < len(input); i++ {
			benchmarkInput[i] = input[i]
		}
		for ; i < len(benchmarkInput); i++ {
			benchmarkInput[i] = i + 1
		}
	})
	for i := 0; i < b.N; i++ {
		Solve(benchmarkInput, 10000)
	}
}
