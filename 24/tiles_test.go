package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTileReference(t *testing.T) {
	testCases := []struct {
		input                string
		expectedNoOfLocators int
		expectedEndLocation  Point
	}{
		{"esew", 3, Point{0, -1}},
		{"nwwswee", 5, Point{0, 0}},
		{"wswnwsene", 5, Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			tileRef := TileReferenceFromString(tC.input)
			assert.Len(t, tileRef.locators, tC.expectedNoOfLocators)
			locX, locY := tileRef.LocateFromReference(0, 0)
			assert.Equal(t, tC.expectedEndLocation.x, locX)
			assert.Equal(t, tC.expectedEndLocation.y, locY)
		})
	}
}
