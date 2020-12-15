package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type InstructionSet uint8

const (
	Part1 = InstructionSet(1)
	Part2 = InstructionSet(2)
)

func Parse(source io.Reader, instrSet InstructionSet) *Program {
	buffer := bufio.NewReader(source)

	instructions := []Instruction{}
	for {
		line, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		statement := string(line)
		if strings.HasPrefix(statement, "mask = ") {
			instructions = append(instructions, FromMaskString(statement[7:]))
		} else if strings.HasPrefix(statement, "mem[") {
			ind := strings.IndexByte(statement[4:], ']')
			memLoc, _ := strconv.ParseUint(statement[4:4+ind], 10, 64)
			ind = strings.Index(statement, " = ")
			memVal, _ := strconv.ParseUint(statement[ind+3:], 10, 64)
			if instrSet == Part1 {
				instructions = append(instructions, SetMemory{memLoc, memVal})
			} else {
				instructions = append(instructions, SetMemoryAlternatives{memLoc, memVal})
			}
		}
	}
	return NewProgram(instructions)
}
