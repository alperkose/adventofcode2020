package main

type Program struct {
	currentMask  Mask
	memory       map[uint64]uint64
	instrcutions []Instruction
}

type Instruction interface {
	Apply(*Program)
}

func NewProgram(instr []Instruction) *Program {
	return &Program{Mask{0, 0}, map[uint64]uint64{}, instr}
}
func (p *Program) Run() map[uint64]uint64 {
	for _, instr := range p.instrcutions {
		instr.Apply(p)
	}
	return p.memory
}
