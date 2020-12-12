package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var preambleSize int
	flag.IntVar(&preambleSize, "p", 25, "(int) Preamble size")
	flag.Parse()

	numbers := ReadSource(os.Stdin)
	err := VerifyWithPreamble(numbers, preambleSize)
	if err != error(nil) {
		fmt.Println(err)
		verf, _ := err.(VerificationError)
		fmt.Println(FindContiguousSum(verf.val, numbers[:verf.index]))
	}
}
