package main

import (
	"fmt"
	"os"
)

func main() {
	passports := ReadSource(os.Stdin, PassportFromString)

	validCounter := 0
	for _, pass := range passports {
		if pass.Valid() {
			validCounter++
		}
	}
	fmt.Println(validCounter)
}
