package main

import "fmt"

type TileLocation struct {
	x, y int
}

type Floor struct {
	tileLocations map[string]TileLocation
}

func TileFloor(tiles []TileReference) *Floor {
	locations := map[string]TileLocation{}
	for _, tRef := range tiles {
		tX, tY := tRef.LocateFromReference(0, 0)
		key := fmt.Sprint(tX, tY)
		if _, ok := locations[key]; ok {
			delete(locations, key)
		} else {
			locations[key] = TileLocation{tX, tY}
		}
	}
	return &Floor{locations}
}

func (f *Floor) BlackTileCount() int {
	return len(f.tileLocations)
}

func (f *Floor) FlipForDays(days int) {
	floorMap := f.tileLocations
	for i := 0; i < days; i++ {
		newFloorMap := map[string]TileLocation{}
		for k, tile := range floorMap {
			adjBl := findAdjacentBlacks(tile.x, tile.y, floorMap)
			if adjBl == 1 || adjBl == 2 {
				newFloorMap[k] = tile
			}

			for _, adj := range adjacency {
				adjX, adjY := adj(tile.x, tile.y)
				adjKey := fmt.Sprint(adjX, adjY)
				if _, ok := floorMap[adjKey]; !ok {
					if findAdjacentBlacks(adjX, adjY, floorMap) == 2 {
						newFloorMap[adjKey] = TileLocation{adjX, adjY}
					}
				}
			}
		}
		floorMap = newFloorMap
	}
	f.tileLocations = floorMap
}

var adjacency = []Locator{ToEast, ToSouthEast, ToSouthWest, ToWest, ToNorthWest, ToNorthEast}

func findAdjacentBlacks(x, y int, floorMap map[string]TileLocation) int {
	blCount := 0
	for _, adj := range adjacency {
		if _, ok := floorMap[fmt.Sprint(adj(x, y))]; ok {
			blCount++
		}
	}
	return blCount
}
