package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoInvalidTicket(t *testing.T) {
	cr := NewCruncher()
	cr.AddRanges(Range{1, 3, "f1"}, Range{5, 7, "f1"}, Range{11, 13, "f2"}, Range{15, 27, "f2"})
	cr.AddOtherTicket([]int{2, 5, 11, 13})

	assert.Equal(t, 0, cr.SumOfInvalidNumbers())
}
func TestSingleInvalidTicket(t *testing.T) {
	cr := NewCruncher()
	cr.AddRanges(Range{1, 3, "f1"}, Range{5, 7, "f1"}, Range{11, 13, "f2"}, Range{15, 27, "f2"})
	cr.AddOtherTicket([]int{2, 5, 10, 13})

	assert.Equal(t, 10, cr.SumOfInvalidNumbers())
}
func TestMultipleInvalidTickets(t *testing.T) {
	cr := NewCruncher()
	cr.AddRanges(Range{1, 3, "f1"}, Range{5, 7, "f1"}, Range{11, 13, "f2"}, Range{15, 27, "f2"})
	cr.AddOtherTicket([]int{2, 5, 10, 13})
	cr.AddOtherTicket([]int{0, 5, 12, 13})
	cr.AddOtherTicket([]int{2, 4, 11, 13})

	assert.Equal(t, 14, cr.SumOfInvalidNumbers())
}

func TestFields(t *testing.T) {
	cr := NewCruncher()
	cr.AddRanges(Range{0, 1, "class"}, Range{4, 19, "class"}, Range{0, 5, "row"}, Range{8, 19, "row"}, Range{0, 13, "seat"}, Range{16, 19, "seat"})
	cr.SetTicket([]int{11, 12, 13})
	cr.AddOtherTicket([]int{3, 9, 18})
	cr.AddOtherTicket([]int{15, 1, 5})
	cr.AddOtherTicket([]int{5, 14, 9})

	assert.Equal(t, []string{"row", "class", "seat"}, cr.FindFieldOrder())
}
