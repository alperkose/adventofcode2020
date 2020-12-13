package main

import "fmt"

func EarliestConsecutiveDepartureTime(buses []Bus) int {
	lastScheduledBus := int(buses[0])
	lastIndex := 0
	increment := lastScheduledBus
	for i := 1; i < len(buses); i++ {
		if buses[i] == IgnoredBus {
			continue
		}
		// fmt.Println("loop", commonNumber, increment, buses[i], i-lastIndex)
		lastScheduledBus, increment = crt(lastScheduledBus, increment, int(buses[i]), i-lastIndex)
		lastIndex = i
	}
	return lastScheduledBus - len(buses) + 1
}

func crt(cumulativeTimestamp, quotient, nextBus, timeInterval int) (int, int) {
	var currInt int
	var a, b = cumulativeTimestamp, nextBus

	for i := 0; i < 1000; i++ {
		currInt = b - a
		// fmt.Println(n1, n2, currDiff)
		if currInt == timeInterval {
			return b, quotient * nextBus
		}
		if currInt > timeInterval {
			a = a + quotient
		}
		if currInt < timeInterval {
			if a > b {
				b = a - (a % nextBus) + nextBus
			} else {
				b = b + nextBus
			}
		}
	}

	panic(fmt.Sprint("NOOOO!", cumulativeTimestamp, quotient, nextBus, timeInterval, a, b, currInt))
}
