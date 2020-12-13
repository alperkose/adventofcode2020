package main

import (
	"math"
)

func Occupy(area []string) []string {
	occupiedArea := make([]string, len(area))
	for i := 0; i < len(area); i++ {
		row := area[i]
		newRow := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			switch row[j] {
			case 'L':
				occupant := occAdjacent(area, i, j)
				if occupant == 0 {
					newRow[j] = '#'
				} else {
					newRow[j] = 'L'
				}
			case '#':
				occupant := occAdjacent(area, i, j) - 1
				if occupant > 3 {
					newRow[j] = 'L'
				} else {
					newRow[j] = '#'
				}
			default:
				newRow[j] = '.'
			}
		}
		occupiedArea[i] = string(newRow)
	}

	return occupiedArea
}

func OccupyExtendedView(area []string) []string {
	occupiedArea := make([]string, len(area))
	limit := int(math.Max(float64(len(area)), float64(len(area[0]))))
	for i := 0; i < len(area); i++ {
		row := area[i]
		newRow := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			switch row[j] {
			case 'L':
				occupant := inSight(area, i, j, limit)
				if occupant == 0 {
					newRow[j] = '#'
				} else {
					newRow[j] = 'L'
				}
			case '#':
				occupant := inSight(area, i, j, limit)
				if occupant > 4 {
					newRow[j] = 'L'
				} else {
					newRow[j] = '#'
				}
			default:
				newRow[j] = '.'
			}
		}
		occupiedArea[i] = string(newRow)
	}

	return occupiedArea
}

func occAdjacent(area []string, i, j int) int {
	count := 0
	if i > 0 {
		count += occRowAdjacent(area[i-1], j)
	}
	count += occRowAdjacent(area[i], j)
	if i < len(area)-1 {
		count += occRowAdjacent(area[i+1], j)
	}
	return count
}

func occRowAdjacent(row string, j int) int {
	count := 0
	if j > 0 && row[j-1] == '#' {
		count++
	}
	if row[j] == '#' {
		count++
	}
	if j < len(row)-1 && row[j+1] == '#' {
		count++
	}
	return count
}

var di = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dj = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func inSight(area []string, i, j, limit int) int {
	H := len(area)
	W := len(area[0])
	// fmt.Println("insight", i, j, H, W, limit)
	eightD := []int{-1, -1, -1, -1, -1, -1, -1, -1}
	count := 0

	for d := 1; d < limit; d++ {
		// fmt.Println(i, j, "distance", d)
		if count >= 8 {
			break
		}
		for k := 0; k < 8; k++ {
			// fmt.Println("\tdirection", k)
			if eightD[k] == -1 {
				ci := di[k]*d + i
				if ci < 0 || ci >= H {
					eightD[k] = 0
					count++
					continue
				}
				cj := dj[k]*d + j
				if cj < 0 || cj >= W {
					eightD[k] = 0
					count++
					continue
				}
				// fmt.Printf("\tarea[%d][%d] = %s\n", ci, cj, string(area[ci][cj]))
				if area[ci][cj] == '#' {
					eightD[k] = 1
					count++
				}
				if area[ci][cj] == 'L' {
					eightD[k] = 0
					count++
				}
			}
		}
	}
	sum := 0
	for _, oneD := range eightD {
		if oneD != -1 {
			sum += oneD
		}
	}
	// fmt.Println(i, j, "sum", sum)
	return sum
}
