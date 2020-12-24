package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cr := NewParser().Parse(os.Stdin)
	fmt.Println(cr.SumOfInvalidNumbers())
	fieldOrders := cr.FindFieldOrder()
	fmt.Println(fieldOrders)
	result := 1
	for i, field := range fieldOrders {
		if strings.HasPrefix(field, "departure") {
			result *= cr.ticket[i]
		}
	}

	fmt.Println(result)
}
