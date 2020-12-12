package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var method string
	flag.StringVar(&method, "m", "combinatorial", "calculation method: combinatorial or tribonacci")
	flag.Parse()

	combinationCalculator := validCombinations
	if method == "tribonacci" {
		combinationCalculator = tribonacci
	} else {
		fmt.Printf("No such method: %s. Using default combinatorial method\n", method)
	}

	input := ReadSource(os.Stdin)
	// fmt.Println(len(input), input)
	oneJ, twoJ, threeJ := JoltDifferenceGroups(input)
	fmt.Printf("Count of one jolt difference: %d\n", oneJ)
	fmt.Printf("Count of two jolt difference: %d\n", twoJ)
	fmt.Printf("Count of three jolt difference: %d\n", threeJ)
	fmt.Printf("%d * %d = %d\n", oneJ, threeJ, oneJ*threeJ)
	// fmt.Println(len(input), input)
	fmt.Println("Combinations:", JoltCombinations(input, combinationCalculator))
	// fmt.Println(len(input), input)

}
