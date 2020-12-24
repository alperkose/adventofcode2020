package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestParsingFullFile(t *testing.T) {
	p := NewParser()
	var cr *Cruncher = p.Parse(strings.NewReader(sampleInput))

	assert.Len(t, cr.Ranges(), 6)

	assert.Equal(t, []int{7, 1, 14}, cr.Ticket())

	assert.Len(t, cr.OtherTickets(), 4)

	assert.Equal(t, 71, cr.SumOfInvalidNumbers())
}

func TestRangeParserWithClassField(t *testing.T) {
	c := NewCruncher()
	ParseRange("class: 1-3 or 5-7", c)
	assert.ElementsMatch(t, []Range{{1, 3, "class"}, {5, 7, "class"}}, c.ranges)
}

func TestRangeParserWithRowField(t *testing.T) {
	c := NewCruncher()
	ParseRange("row: 6-11 or 33-44", c)
	assert.ElementsMatch(t, []Range{{6, 11, "row"}, {33, 44, "row"}}, c.ranges)
}

func TestTicketParser(t *testing.T) {
	c := NewCruncher()
	ParseTicket("7,1,14", c)
	assert.ElementsMatch(t, []int{7, 1, 14}, c.Ticket())
}
func TestOtherTicketParser(t *testing.T) {
	c := NewCruncher()
	ParseOtherTicket("7,1,14", c)
	assert.ElementsMatch(t, []int{7, 1, 14}, c.OtherTickets()[0])
	ParseOtherTicket("7,3,47", c)
	assert.ElementsMatch(t, []int{7, 3, 47}, c.OtherTickets()[1])
}
