package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var setFlag string
	flag.StringVar(&setFlag, "set", "part2", "instruction set: part1 or part2")
	flag.Parse()
	instrSet := Part2
	if setFlag == "part1" {
		instrSet = Part1
	} else {
		fmt.Println("Using default instrcution set (Part2)")
	}

	mem := Parse(os.Stdin, instrSet).Run()
	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Printf("Memory:\nSum:\n\t%64b\n\t%64d\n", sum, sum)
}
