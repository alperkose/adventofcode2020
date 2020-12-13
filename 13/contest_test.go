package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestFindOrdering(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []Bus
		expected int
	}{
		{"Single Bus", []Bus{Bus(7)}, 7},
		{"Two Buses", []Bus{Bus(7), Bus(11)}, 21},
		{"Two Buses and an ignored one", []Bus{Bus(7), IgnoredBus, Bus(11)}, 42},
		{"More combinations", []Bus{Bus(7), Bus(13), IgnoredBus, IgnoredBus, Bus(59), IgnoredBus, Bus(31), Bus(19)}, 1068781},
		{"Sample #1", []Bus{Bus(17), IgnoredBus, Bus(13), Bus(19)}, 3417},
		{"Sample #2", []Bus{Bus(67), Bus(7), Bus(59), Bus(61)}, 754018},
		{"Sample #3", []Bus{Bus(67), IgnoredBus, Bus(7), Bus(59), Bus(61)}, 779210},
		{"Sample #4", []Bus{Bus(67), Bus(7), IgnoredBus, Bus(59), Bus(61)}, 1261476},
		{"Sample #5", []Bus{Bus(1789), Bus(37), Bus(47), Bus(1889)}, 1202161486},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			earliestTime := EarliestOrderedTime(tC.input)
			assert.Equal(t, tC.expected, earliestTime)
		})
	}
}
