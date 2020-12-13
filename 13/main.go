package main

import (
	"fmt"
	"os"
)

func main() {
	earliestTimeToLeave, buses := Parse(os.Stdin)
	bus, waitTime := FindEarliestBus(earliestTimeToLeave, buses)
	fmt.Println("Earliest bus to take:", bus, ", wait time:", waitTime, "-->", (waitTime * int(bus)))
	// schduleStart := ParallelFind(buses, 142889000000251, 100000000000)
	schduleStart := EarliestConsecutiveDepartureTime(buses)
	fmt.Println("Contest earliest bus time:", schduleStart)
}
