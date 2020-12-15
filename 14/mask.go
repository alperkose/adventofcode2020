package main

type Mask struct {
	Mask  uint64
	Value uint64
}

func (m Mask) Apply(p *Program) {
	p.currentMask = m
}

const upperBound = uint64(1 << 35)

func (m Mask) MapTo(in uint64) uint64 {

	// fmt.Printf("in       %36b\n", in)
	// fmt.Printf("mask     %36b\n", m.Mask)
	// fmt.Printf("maskVal  %36b\n", m.Value)
	// fmt.Println()

	out := uint64(0)
	var maskBit, valueBit, inBit uint64
	for shiftingBit := uint64(1); shiftingBit <= upperBound; shiftingBit <<= 1 {
		maskBit = m.Mask & shiftingBit
		// fmt.Printf("maskBit  %36b\n", maskBit)
		inBit = in & shiftingBit
		// fmt.Printf("inBit    %36b\n", inBit)
		if maskBit == 0 {
			out |= inBit
			continue
		}
		valueBit = m.Value & shiftingBit
		// fmt.Printf("valueBit %36b\n", valueBit)
		if valueBit > 0 {
			out |= valueBit
		} else if inBit > 0 {
			out &= ^maskBit
		}
		// fmt.Printf("outBit   %36b\n", out)
		// fmt.Println()
	}
	return out
}

func (m Mask) MapToAlternatives(in uint64) []uint64 {
	floatingAlternatives := map[uint64]bool{0: true}
	out := uint64(0)
	var maskBit, valueBit, inBit uint64
	for shiftingBit := uint64(1); shiftingBit <= upperBound; shiftingBit <<= 1 {
		maskBit = m.Mask & shiftingBit
		inBit = in & shiftingBit
		if maskBit == 0 { // 'X'
			out &= ^shiftingBit
			addToAlternatives(floatingAlternatives, shiftingBit)
			continue
		}
		valueBit = m.Value & shiftingBit
		if valueBit > 0 {
			out |= valueBit
		} else {
			out |= inBit
		}
	}
	alternatives := make([]uint64, len(floatingAlternatives))
	altInd := 0
	for k, _ := range floatingAlternatives {
		alternatives[altInd] = out + k
		altInd++
	}
	return alternatives
}

func addToAlternatives(alt map[uint64]bool, bit uint64) {
	t := make([]uint64, len(alt)*2)
	ind := 0
	for k, _ := range alt {
		t[ind] = k &^ bit
		t[ind+1] = k | bit
		ind += 2
	}
	for _, v := range t {
		alt[v] = true
	}

}

func FromMaskString(input string) Mask {

	mask, maskedValue := uint64(0), uint64(0)
	shiftingBit := uint64(1)
	for i := len(input) - 1; i >= 0; i-- {
		switch input[i] {
		case '1':
			maskedValue |= shiftingBit
			mask |= shiftingBit
		case '0':
			mask |= shiftingBit
		}
		shiftingBit <<= 1
	}
	return Mask{mask, maskedValue}
}
