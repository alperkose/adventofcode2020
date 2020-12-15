package main

type SetMemory struct {
	loc   uint64
	value uint64
}

func (sm SetMemory) Apply(p *Program) {
	p.memory[sm.loc] = p.currentMask.MapTo(sm.value)
}

type SetMemoryAlternatives struct {
	loc   uint64
	value uint64
}

func (sm SetMemoryAlternatives) Apply(p *Program) {
	// fmt.Println("setting memory alternatives", sm)
	alternativeLocs := p.currentMask.MapToAlternatives(sm.loc)
	for _, loc := range alternativeLocs {
		p.memory[loc] = sm.value
	}
}
