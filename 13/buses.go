package main

type Bus int

var IgnoredBus = Bus(-1)

func FindEarliestBus(earliestTime int, buses []Bus) (Bus, int) {
	var earliestBus = buses[0]
	var waitTimeToEarliestBus = waitTimeTo(earliestTime, buses[0])
	var waitTime int
	for _, bus := range buses {
		if bus == IgnoredBus {
			continue
		}
		waitTime = waitTimeTo(earliestTime, bus)
		if waitTime < waitTimeToEarliestBus {
			earliestBus, waitTimeToEarliestBus = bus, waitTime
		}
	}
	return earliestBus, waitTimeToEarliestBus
}

func waitTimeTo(earliestTime int, bus Bus) int {
	return (int(bus) - (earliestTime % int(bus))) % int(bus)
}
