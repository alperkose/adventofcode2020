package main

import (
	"fmt"
	"os"
)

func main() {
	groups := ReadSource(os.Stdin)
	sum := 0
	for _, group := range groups {
		sum += group.CommonAnswers()
	}

	fmt.Println(sum)
}
