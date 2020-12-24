package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

type LineParser func(string, *Cruncher)

func (p Parser) Parse(source io.Reader) *Cruncher {
	cruncher := NewCruncher()
	buffer := bufio.NewReader(source)

	var lineParser LineParser = ParseRange

	for {
		lineAsByte, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		line := strings.Trim(string(lineAsByte), " ")
		if len(line) == 0 {
			continue
		}

		if line == "your ticket:" {
			lineParser = ParseTicket
			continue
		}
		if line == "nearby tickets:" {
			lineParser = ParseOtherTicket
			continue
		}
		lineParser(line, cruncher)
	}

	return cruncher
}

var rangeRE = regexp.MustCompile(`([\w\s]+):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)`)

func ParseRange(line string, cr *Cruncher) {
	parts := rangeRE.FindStringSubmatch(line)
	var begin, end int
	field := parts[1]
	begin, _ = strconv.Atoi(parts[2])
	end, _ = strconv.Atoi(parts[3])
	r1 := Range{begin, end, field}
	begin, _ = strconv.Atoi(parts[4])
	end, _ = strconv.Atoi(parts[5])
	r2 := Range{begin, end, field}
	cr.AddRanges(r1, r2)
}

func ParseTicket(line string, cr *Cruncher) {
	cr.SetTicket(parseTicket(line))
}

func ParseOtherTicket(line string, cr *Cruncher) {
	cr.AddOtherTicket(parseTicket(line))
}

func parseTicket(line string) []int {
	parts := strings.Split(line, ",")
	ticket := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		ticket[i], _ = strconv.Atoi(parts[i])
	}
	return ticket
}
