package main

import (
	"fmt"
	"os"
)

func main() {
	rules, samples := Parse(os.Stdin)
	engine := NewEngine(rules)
	sum := 0
	for _, sample := range samples {
		if engine.Matches(sample) {
			fmt.Println(sample)
			sum++
		}
	}
	fmt.Println(sum)

}
