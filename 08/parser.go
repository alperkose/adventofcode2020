package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(source io.Reader) *Program {
	bufferedSource := bufio.NewReader(source)
	pr := &Program{[]Instruction{}, 0, 0}
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		parts := strings.Split(string(line), " ")
		arg, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "acc":
			pr.addInstruction(Accumulator{arg})
		case "jmp":
			pr.addInstruction(Jump{arg})
		case "nop":
			pr.addInstruction(NOp{arg})
		}
	}
	return pr
}

type Instruction interface {
	Exec(*Program)
	Type() string
	Arg() int
}

type Program struct {
	instructions []Instruction
	acc          int
	pointer      int
}

func (pr *Program) Run() (int, error) {
	// fmt.Println("===================")
	pr.acc = 0
	pr.pointer = 0
	executedLines := map[int]bool{}
	for {
		if _, ok := executedLines[pr.pointer]; ok {
			return pr.acc, fmt.Errorf("Loop Detected")
		}
		// fmt.Println(pr.pointer, pr.instructions[pr.pointer].Type(), pr.instructions[pr.pointer].Arg())
		executedLines[pr.pointer] = true
		pr.instructions[pr.pointer].Exec(pr)
		pr.pointer++
		if pr.pointer >= len(pr.instructions) {
			break
		}
	}
	return pr.acc, nil
}

func (pr *Program) FixNRun() (int, error) {
	for i := 0; i < len(pr.instructions); i++ {
		currInst := pr.instructions[i]
		if currInst.Type() == "jmp" {
			pr.instructions[i] = NOp{currInst.Arg()}
			if acc, err := pr.Run(); err == nil {
				return acc, err
			}
		} else if currInst.Type() == "nop" {
			pr.instructions[i] = Jump{currInst.Arg()}
			if acc, err := pr.Run(); err == nil {
				return acc, err
			}
		}
		pr.instructions[i] = currInst
	}
	return 0, fmt.Errorf("Couldn't find a fix")
}

func (pr *Program) swapAndRun()

func (pr *Program) addInstruction(inst Instruction) {
	pr.instructions = append(pr.instructions, inst)
}

func (pr *Program) Increment(delta int) {
	pr.acc += delta
}

func (pr *Program) Offset(o int) {
	pr.pointer += (o - 1)
}

type Accumulator struct {
	delta int
}

func (a Accumulator) Exec(pr *Program) {
	pr.Increment(a.delta)
}
func (a Accumulator) Type() string {
	return "acc"
}

func (a Accumulator) Arg() int {
	return a.delta
}

type Jump struct {
	offset int
}

func (a Jump) Exec(pr *Program) {
	pr.Offset(a.offset)
}
func (a Jump) Type() string {
	return "jmp"
}
func (a Jump) Arg() int {
	return a.offset
}

type NOp struct {
	arg int
}

func (a NOp) Exec(pr *Program) {
}
func (a NOp) Type() string {
	return "nop"
}
func (a NOp) Arg() int {
	return a.arg
}
