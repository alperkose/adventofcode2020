package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleBlackTile(t *testing.T) {
	floor := TileFloor([]TileReference{
		TileReferenceFromString("esew"),
	})
	assert.Equal(t, 1, floor.BlackTileCount())
}
func TestTwoBlackTiles(t *testing.T) {
	floor := TileFloor([]TileReference{
		TileReferenceFromString("esew"),
		TileReferenceFromString("esewe"),
	})
	assert.Equal(t, 2, floor.BlackTileCount())
}
func TestOneBlackTileConvertedToWhite(t *testing.T) {
	floor := TileFloor([]TileReference{
		TileReferenceFromString("esew"),
		TileReferenceFromString("esewe"),
		TileReferenceFromString("swseeenw"),
	})
	fmt.Println(floor.tileLocations)
	assert.Equal(t, 1, floor.BlackTileCount())
}
