package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	ship := Ship{East, 0, 0, 10, 1}
	instructions := Parse(os.Stdin)
	ship = Move(ship, instructions)
	fmt.Println(ship)
	fmt.Println(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}
