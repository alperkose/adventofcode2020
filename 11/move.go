package main

import (
	"fmt"
	"strings"
)

func MoveUntilStabilized(area []string, fOccupy func([]string) []string) (occupied []string, occupants, turns int) {
	occupied = area
	turns = 0
	occupiedStr := ""
	for {
		occupied = fOccupy(occupied)
		updatedStr := fmt.Sprint(occupied)
		// fmt.Println(updatedStr)
		turns++
		if updatedStr == occupiedStr {
			return occupied, strings.Count(updatedStr, "#"), turns
		}
		occupiedStr = updatedStr
		if turns > 1000 {
			fmt.Println(turns, "and counting")
		}
	}
}
