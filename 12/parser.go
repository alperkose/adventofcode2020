package main

import (
	"bufio"
	"io"
)

func Parse(source io.Reader, toInstruction func(string) Instruction) []Instruction {
	bufReader := bufio.NewReader(source)
	instructions := []Instruction{}
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		instructions = append(instructions, toInstruction(string(line)))
	}
	return instructions
}
