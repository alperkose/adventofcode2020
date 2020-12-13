package main

import (
	"fmt"
	"strconv"
)

type Instruction interface {
	Apply(ship Ship) Ship
}

func FromString(instr string) Instruction {
	instrCode := instr[0]
	instrValue, _ := strconv.Atoi(instr[1:])

	switch instrCode {
	case 'F':
		return ForwardByWp(instrValue)
		// return Forward(instrValue)
	case 'L':
		return RotateWpLeft(instrValue)
		// return RotateLeft(instrValue)
	case 'R':
		return RotateWpRight(instrValue)
		// return RotateRight(instrValue)
	case 'E':
		return MoveWpEast(instrValue)
		// return MoveEast(instrValue)
	case 'N':
		return MoveWpNorth(instrValue)
		// return MoveNorth(instrValue)
	case 'W':
		return MoveWpWest(instrValue)
		// return MoveWest(instrValue)
	case 'S':
		return MoveWpSouth(instrValue)
		// return MoveSouth(instrValue)
	}
	return DummyInstruction{}
}

type DummyInstruction struct {
}

func (d DummyInstruction) Apply(ship Ship) Ship {
	return ship
}

func Move(ship Ship, instructions []Instruction) Ship {
	for _, instruction := range instructions {
		ship = instruction.Apply(ship)
		fmt.Println(ship)
	}
	return ship
}
