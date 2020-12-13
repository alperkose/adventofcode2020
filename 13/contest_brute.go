package main

import "fmt"

func FindOrderingStartingFrom(buses []Bus, startTime, bound int) int {
	highestFreqBus, offset := max(buses)
	intervalLength := len(buses)
	// fmt.Println("Starting loop over:", buses)
	fmt.Println("intervalLength:", intervalLength)
	fmt.Println("highestFreqBus:", highestFreqBus, ", offset:", offset)
	fmt.Println("StartTime", startTime)
	var upperBoundary = startTime + bound
	var allMatch = false
	var loopIteration = int(highestFreqBus)
	var earliestBus = (startTime % loopIteration) + startTime - offset
	for {
		earliestBus += int(highestFreqBus)
		if earliestBus > upperBoundary {
			break
		}
		for i := 0; i < intervalLength; i++ {
			// fmt.Println(i, "bus", buses[i], "time", waitTimeTo(earliestBus, buses[i]))
			allMatch = false
			if buses[i] == IgnoredBus {
				continue
			}
			if waitTimeTo(earliestBus, buses[i]) != i {
				break
			}
			allMatch = true
		}
		if allMatch {
			return earliestBus
		}
	}
	return -1
}

func ParallelFind(buses []Bus, startTime, boundary int) int {

	buffer := 100
	c := make(chan int, buffer)

	start := startTime
	for {
		for i := 0; i < buffer; i++ {
			go func(c chan int, start, bound int) {
				c <- FindOrderingStartingFrom(buses, start, bound)
			}(c, start, boundary)
			start += boundary
		}
		buffer = 0

		result := <-c
		if result < 0 {
			buffer++
		} else {
			return result
		}
	}

}
func max(buses []Bus) (Bus, int) {
	var max = buses[0]
	var locMax = 0
	for loc, bus := range buses {
		if bus > max {
			max = bus
			locMax = loc
		}
	}
	return max, locMax
}
