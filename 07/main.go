package main

import (
	"fmt"
	"os"
)

func main() {
	rules := ReadSource(os.Stdin)
	output := rules.BagsCanContain("shiny gold")
	fmt.Printf("%q\n%d\n", output, len(output))
	fmt.Printf("%d\n", rules.CountBagsWithin("shiny gold"))
}
