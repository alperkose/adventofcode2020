package main

import (
	"fmt"
	"sort"
)

type VerificationError struct {
	val, index int
}

func (e VerificationError) Error() string {
	return fmt.Sprintf("Doesn't add up to %d", e.val)
}

func VerifyWithPreamble(numbers []int, preambleSize int) error {
	preamble := append([]int{}, numbers[:preambleSize]...)
	sort.Ints(preamble)
	for i, n := range numbers[preambleSize:] {
		if !canBeCalculatedFrom(n, preamble) {
			return VerificationError{n, i + preambleSize}
		}
		preamble[sort.SearchInts(preamble, numbers[i])] = n
		sort.Ints(preamble)
	}
	return nil
}

func FindContiguousSum(val int, preamble []int) int {
	for i := len(preamble) - 1; i >= 1; i-- {
		// fmt.Println(preamble[:i])
		sum := preamble[i]
		for j := i - 1; j >= 0; j-- {
			// fmt.Println(sum, preamble[j], val)
			sum += preamble[j]
			if sum == val {
				interval := preamble[j : i+1]
				sort.Ints(interval)
				fmt.Println(interval)
				return interval[0] + interval[len(interval)-1]
			}
			if sum > val {
				break
			}
		}
	}
	return -1
}

func canBeCalculatedFrom(n int, preamble []int) bool {
	// fmt.Println("canBeCalculatedFrom", n, preamble)
	for i := len(preamble) - 1; i >= 1; i-- {
		if preamble[i] > n {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if (preamble[i] + preamble[j]) == n {
				return true
			}
			if (preamble[i] + preamble[j]) < n {
				break
			}
		}
	}

	return false
}
