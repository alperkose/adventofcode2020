package main

import (
	"bufio"
	"io"
)

func Parse(source io.Reader) []Instruction {
	bufReader := bufio.NewReader(source)
	instructions := []Instruction{}
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		instructions = append(instructions, FromString(string(line)))
	}
	return instructions
}
