package main

import (
	"sort"
)

func JoltDifferenceGroups(deviceJolts []int) (int, int, int) {
	deviceJolts = append([]int{}, deviceJolts...)
	deviceJolts = append(deviceJolts, 0)
	sort.Ints(deviceJolts)
	// fmt.Println(len(deviceJolts), deviceJolts)
	groups := []int{0, 0, 1}
	for i := 1; i < len(deviceJolts); i++ {
		diff := deviceJolts[i] - deviceJolts[i-1]
		groups[diff-1] += 1
	}
	return groups[0], groups[1], groups[2]
}

func JoltCombinations(deviceJolts []int, fCombination func([]int) int) int {
	deviceJolts = append([]int{0}, deviceJolts...)
	sort.Ints(deviceJolts)
	// fmt.Println(len(deviceJolts), deviceJolts)

	comb := 1
	marker := 0
	for i := 1; i < len(deviceJolts); i++ {
		// fmt.Println(i, deviceJolts[i], deviceJolts[i-1])
		if deviceJolts[i]-deviceJolts[i-1] == 3 {
			comb *= fCombination(deviceJolts[marker:i])
			marker = i
		}
	}
	// return comb * validCombinations(deviceJolts[marker:])
	return comb * fCombination(deviceJolts[marker:])
}

func validCombinations(deviceJolts []int) int {
	joltCount := len(deviceJolts)
	if joltCount == 1 {
		return 1
	}
	comb := validCombinations(deviceJolts[1:])
	if joltCount > 2 && (deviceJolts[2]-deviceJolts[0]) <= 3 {
		comb += validCombinations(deviceJolts[2:])
	}
	if joltCount > 3 && (deviceJolts[3]-deviceJolts[0]) <= 3 {
		comb += validCombinations(deviceJolts[3:])
	}
	return comb
}

func tribonacci(deviceJolts []int) int {
	return tribonacciNumber(len(deviceJolts) - 1)
}
func tribonacciNumber(c int) int {
	if c <= 1 {
		return 1
	}
	if c == 2 {
		return 2
	}
	if c == 3 {
		return 4
	}

	return tribonacciNumber(c-1) + tribonacciNumber(c-2) + tribonacciNumber(c-3)
}
