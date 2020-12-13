package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingInput(t *testing.T) {
	testCases := []struct {
		desc         string
		input        string
		earliestTime int
		listOfBuses  []Bus
	}{
		{
			desc:         "Input with all buses",
			input:        "425\n1,2,3,5,7,11",
			earliestTime: 425,
			listOfBuses:  []Bus{Bus(1), Bus(2), Bus(3), Bus(5), Bus(7), Bus(11)},
		}, {
			desc:         "Input with some ignored buses",
			input:        "623\n13,17,x,29,x,37",
			earliestTime: 623,
			listOfBuses:  []Bus{Bus(13), Bus(17), IgnoredBus, Bus(29), IgnoredBus, Bus(37)},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			source := strings.NewReader(tC.input)
			earliest, buses := Parse(source)
			assert.Equal(t, tC.earliestTime, earliest)
			assert.Equal(t, tC.listOfBuses, buses)
		})
	}
}
