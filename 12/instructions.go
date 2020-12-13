package main

import (
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
		return Forward(instrValue)
	case 'L':
		return RotateLeft(instrValue)
	case 'R':
		return RotateRight(instrValue)
	case 'E':
		return MoveEast(instrValue)
	case 'N':
		return MoveNorth(instrValue)
	case 'W':
		return MoveWest(instrValue)
	case 'S':
		return MoveSouth(instrValue)
	}
	return DummyInstruction{}
}

func FromStringUsingWP(instr string) Instruction {
	instrCode := instr[0]
	instrValue, _ := strconv.Atoi(instr[1:])

	switch instrCode {
	case 'F':
		return ForwardByWp(instrValue)
	case 'L':
		return RotateWpLeft(instrValue)
	case 'R':
		return RotateWpRight(instrValue)
	case 'E':
		return MoveWpEast(instrValue)
	case 'N':
		return MoveWpNorth(instrValue)
	case 'W':
		return MoveWpWest(instrValue)
	case 'S':
		return MoveWpSouth(instrValue)
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
		// fmt.Println(ship)
	}
	return ship
}
