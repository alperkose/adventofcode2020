package main

import (
	"fmt"
	"os"
)

func main() {
	cubes := NewParser().Parse(os.Stdin)
	s := NewSolver(cubes, Neighbors3D)
	fmt.Println("3D: after 6 iterations", len(s.Next().Next().Next().Next().Next().Next().ActiveCubes()))

	s = NewSolver(cubes, Neighbors4D)
	fmt.Println("4D: after 6 iterations", len(s.Next().Next().Next().Next().Next().Next().ActiveCubes()))
}
