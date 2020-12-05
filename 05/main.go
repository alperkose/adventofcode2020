package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	inputs := ReadSource(os.Stdin)
	maxSeatID := -1
	bpasses := make([]BoardingPass, len(inputs))
	for i, input := range inputs {
		pass, err := NewBoadingPass(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid Boarding Pass %s, ignoring...", input)
		}
		bpasses[i] = pass
		if pass.SeatID() > maxSeatID {
			maxSeatID = pass.SeatID()
		}
	}
	fmt.Printf("Max Seat ID is %d\n", maxSeatID)
	sort.Sort(SortBySeatID(bpasses))

	fmt.Printf("First seat: %d\n", bpasses[0].SeatID())
	for i := 1; i < len(bpasses); i++ {
		prevSeatID := bpasses[i-1].SeatID()
		nextSeatID := bpasses[i].SeatID()
		if (nextSeatID - prevSeatID) > 1 {
			fmt.Printf("previous: %d, next: %d\n", prevSeatID, nextSeatID)
		}
	}
	fmt.Printf("Last seat: %d\n", bpasses[len(bpasses)-1].SeatID())
}

type SortBySeatID []BoardingPass

func (a SortBySeatID) Len() int           { return len(a) }
func (a SortBySeatID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBySeatID) Less(i, j int) bool { return a[i].SeatID() < a[j].SeatID() }
