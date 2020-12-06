package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var counterFlag string
	flag.StringVar(&counterFlag, "c", "common", "calculation method: common or union")
	flag.Parse()

	fGroupCount := func(g *Group) int {
		return g.CommonAnswers()
	}
	if counterFlag == "union" {
		fGroupCount = func(g *Group) int {
			return g.CountAnswers()
		}
	}

	groups := ReadSource(os.Stdin)

	sum := 0
	for _, group := range groups {
		sum += fGroupCount(group)
	}

	fmt.Println(sum)
}
