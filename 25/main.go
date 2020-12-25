package main

import "fmt"

func main() {
	cardLoopSize := FindLoopSize(7, 8987316)
	// doorLoopSize := FindLoopSize(7, 14681524)
	fmt.Println(Transform(14681524, cardLoopSize))
}
