package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var methodFlag string
	flag.StringVar(&methodFlag, "m", "firstsight", "calculation method: firstsight or adjacency")
	flag.Parse()

	method := OccupyExtendedView
	if methodFlag == "adjacency" {
		method = Occupy
	} else if methodFlag != "firstsight" {
		fmt.Printf("No such method: %s. Using default firstsight method\n", methodFlag)
	}

	_, occupants, turns := MoveUntilStabilized(ReadSource(os.Stdin), method)
	fmt.Println("occupants", occupants, "turns", turns)
}
