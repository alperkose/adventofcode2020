package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestSample(t *testing.T) {
	tileRefs := Parse(strings.NewReader(input))
	floor := TileFloor(tileRefs)
	// fmt.Println(floor.tileLocations)
	assert.Equal(t, 10, floor.BlackTileCount())
}

func TestFlippingFloorsForDays(t *testing.T) {
	testCases := []struct {
		days               int
		expectedBlackCount int
	}{
		{1, 15},
		{2, 12},
		{3, 25},
		{4, 14},
		{5, 23},
		{6, 28},
		{7, 41},
		{8, 37},
		{9, 49},
		{10, 37},
		{20, 132},
		{30, 259},
		{40, 406},
		{50, 566},
		{60, 788},
		{70, 1106},
		{80, 1373},
		{90, 1844},
		{100, 2208},
	}
	tileRefs := Parse(strings.NewReader(input))
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.days, "day(s)"), func(t *testing.T) {
			floor := TileFloor(tileRefs)
			// fmt.Println(floor)
			floor.FlipForDays(tC.days)
			// fmt.Println(floor)
			assert.Equal(t, tC.expectedBlackCount, floor.BlackTileCount())
		})
	}
}
