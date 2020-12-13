package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	var methodFlag string
	flag.StringVar(&methodFlag, "m", "waypoint", "navigation method: waypoint or ship")
	flag.Parse()

	method := FromStringUsingWP
	if methodFlag == "ship" {
		method = FromString
	} else if methodFlag != "waypoint" {
		fmt.Printf("No such method: %s. Using default waypoint method\n", methodFlag)
	}

	ship := Ship{East, 0, 0, 10, 1}
	instructions := Parse(os.Stdin, method)
	ship = Move(ship, instructions)
	fmt.Println(ship)
	fmt.Println(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}
