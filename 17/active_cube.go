package main

import "fmt"

func AC3(x, y, z int) ActiveCube {
	return ActiveCube{x, y, z, 0}
}
func AC4(x, y, z, w int) ActiveCube {
	return ActiveCube{x, y, z, w}
}

type ActiveCube struct {
	x, y, z, w int
}

func (c ActiveCube) String() string {
	return fmt.Sprint(c.x, ",", c.y, ",", c.z, ",", c.w)
}

var neigborCalculations = []struct {
	x, y, z int
}{
	{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
	{-1, 0, -1}, {0, 0, -1}, {1, 0, -1},
	{-1, 1, -1}, {0, 1, -1}, {1, 1, -1},
	{-1, -1, 0}, {0, -1, 0}, {1, -1, 0},
	{-1, 0, 0}, {0, 0, 0}, {1, 0, 0},
	{-1, 1, 0}, {0, 1, 0}, {1, 1, 0},
	{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
	{-1, 0, 1}, {0, 0, 1}, {1, 0, 1},
	{-1, 1, 1}, {0, 1, 1}, {1, 1, 1},
}

func Neighbors3D(c ActiveCube) []ActiveCube {
	neighbors := make([]ActiveCube, 26)
	ind := 0
	for _, delta := range neigborCalculations {
		if delta.x == 0 && delta.y == 0 && delta.z == 0 {
			continue
		}
		neighbors[ind] = ActiveCube{c.x + delta.x, c.y + delta.y, c.z + delta.z, c.w}
		ind++
	}
	return neighbors
}

func Neighbors4D(c ActiveCube) []ActiveCube {
	neighbors := make([]ActiveCube, 80)
	ind := 0
	for _, w := range []int{-1, 0, 1} {
		for _, delta := range neigborCalculations {
			if delta.x == 0 && delta.y == 0 && delta.z == 0 && w == 0 {
				continue
			}
			neighbors[ind] = ActiveCube{c.x + delta.x, c.y + delta.y, c.z + delta.z, c.w + w}
			ind++
		}
	}
	return neighbors
}
