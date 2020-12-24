package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	game := flag.String("game", "recursive", "recursive or combat")
	flag.Parse()
	solver := Parse(os.Stdin)
	if *game == "combat" {
		printResult(solver.Combat())
	} else {
		printResult(solver.RecursiveCombat())
	}

}

func printResult(winner []int) {
	coef := len(winner)
	sum := 0
	for i := 0; i < len(winner); i++ {
		sum += coef * winner[i]
		coef--
	}
	fmt.Println(sum)
}
