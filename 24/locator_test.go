package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Point struct{ x, y int }

func TestLocators(t *testing.T) {
	testCases := []struct {
		ref      string
		from     Point
		expected Point
	}{
		//e, se, sw, w, nw, ne
		//      (x, y+1) (x+1,y+1)
		//  (x-1,y)   (x,y)   (x+1,y)
		//    (x-1, y-1) (x,y-1)

		{"e", Point{1, 2}, Point{2, 2}},
		{"se", Point{1, 2}, Point{1, 1}},
		{"sw", Point{1, 2}, Point{0, 1}},
		{"w", Point{1, 2}, Point{0, 2}},
		{"nw", Point{1, 2}, Point{1, 3}},
		{"ne", Point{1, 2}, Point{2, 3}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("from %v to %s -> %v", tC.from, tC.ref, tC.expected), func(t *testing.T) {
			locate := LocatorFromRef(tC.ref)
			actualX, actualY := locate(tC.from.x, tC.from.y)
			assert.Equal(t, tC.expected.x, actualX)
			assert.Equal(t, tC.expected.y, actualY)
		})
	}
}
