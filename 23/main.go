package main

import "fmt"

func main() {
	input := []int{6, 1, 4, 7, 5, 2, 8, 3, 9}
	fmt.Println(Solve(input, 100))

	extendedInput := make([]int, 1000000)
	var i int
	for i = 0; i < len(input); i++ {
		extendedInput[i] = input[i]
	}
	for ; i < len(extendedInput); i++ {
		extendedInput[i] = i + 1
	}
	result := Solve(extendedInput, 10000000)
	fmt.Println(result[0], result[1], len(result))
}
