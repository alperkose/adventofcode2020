package main

import (
	"fmt"
	"os"
)

func main() {
	tileRefs := Parse(os.Stdin)
	floor := TileFloor(tileRefs)
	fmt.Println(floor.BlackTileCount())
	floor.FlipForDays(100)
	fmt.Print(floor.BlackTileCount())
}
