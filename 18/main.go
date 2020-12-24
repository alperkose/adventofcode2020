package main

import (
	"fmt"
	"os"
)

func main() {
	terms := NewParser().Parse(os.Stdin)
	sum := 0
	for _, term := range terms {
		sum += term.Value()
	}
	fmt.Println("sum =", sum)
}
