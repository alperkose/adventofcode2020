package main

import (
	"fmt"
	"os"
)

func main() {
	solver := Parse(os.Stdin)
	fmt.Println(solver.NonAllergenicIngredients())
}
