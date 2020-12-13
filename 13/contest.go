package main

import "fmt"

func EarliestOrderedTime(buses []Bus) int {
	commonNumber := int(buses[0])
	lastIndex := 0
	increment := commonNumber
	for i := 1; i < len(buses); i++ {
		if buses[i] == IgnoredBus {
			continue
		}
		// fmt.Println("loop", commonNumber, increment, buses[i], i-lastIndex)
		commonNumber, increment = findCommonBase(commonNumber, increment, int(buses[i]), i-lastIndex)
		lastIndex = i
	}
	return commonNumber - len(buses) + 1
}

func findCommonBase(commonNumber, inc, numberToCheck, diff int) (int, int) {
	var currDiff int
	var n1, n2 = commonNumber, numberToCheck

	for i := 0; i < 1000; i++ {
		currDiff = n2 - n1
		// fmt.Println(n1, n2, currDiff)
		if currDiff == diff {
			return n2, inc * numberToCheck
		}
		if currDiff > diff {
			n1 = n1 + inc
		}
		if currDiff < diff {
			if n1 > n2 {
				n2 = n1 - (n1 % numberToCheck) + numberToCheck
			} else {
				n2 = n2 + numberToCheck
			}
		}
	}

	panic(fmt.Sprint("NOOOO!", commonNumber, inc, numberToCheck, diff, n1, n2, currDiff))
}
