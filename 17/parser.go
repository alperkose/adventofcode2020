package main

import (
	"bufio"
	"io"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(source io.Reader) []ActiveCube {
	buffer := bufio.NewReader(source)
	cubes := []ActiveCube{}
	row := 0
	for {
		line, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		for i, c := range line {
			if c == '#' {
				cubes = append(cubes, ActiveCube{i, row, 0, 0})
			}
		}
		row++
	}
	return cubes
}
