package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Parse(source io.Reader) *Solver {
	buffer := bufio.NewReader(source)
	p1 := []int{}
	p2 := []int{}
	var appender func(int)
	for {
		line, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		if len(strings.Trim(lineStr, " ")) == 0 {
			continue
		}
		if lineStr == "Player 1:" {
			appender = func(c int) { p1 = append(p1, c) }
			continue
		}
		if lineStr == "Player 2:" {
			appender = func(c int) { p2 = append(p2, c) }
			continue
		}
		card, _ := strconv.Atoi(lineStr)
		appender(card)
	}
	return &Solver{p1, p2}
}
