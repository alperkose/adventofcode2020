package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEarliestBusToTake(t *testing.T) {
	testCases := []struct {
		desc         string
		earliestTime int
		buses        []Bus
		expectedBus  Bus
		waitTime     int
	}{
		{"Single bus", 10, []Bus{Bus(7)}, Bus(7), 4},
		{"Two buses", 10, []Bus{Bus(7), Bus(11)}, Bus(11), 1},
		{"Two buses and an ignored bus", 10, []Bus{Bus(7), IgnoredBus, Bus(11)}, Bus(11), 1},
		{"More buses", 939, []Bus{Bus(7), Bus(13), IgnoredBus, IgnoredBus, Bus(59), IgnoredBus, Bus(31), Bus(19)}, Bus(59), 5},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actualBusToTake, actualWaitTime := FindEarliestBus(tC.earliestTime, tC.buses)
			assert.Equal(t, tC.expectedBus, actualBusToTake)
			assert.Equal(t, tC.waitTime, actualWaitTime)
		})
	}
}
