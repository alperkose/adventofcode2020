package main

import (
	"fmt"
	"os"
)

func main() {
	pr := Parse(os.Stdin)

	fmt.Println(pr.FixNRun())
}
