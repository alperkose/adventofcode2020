package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingAMemoryOfAllForced1Mask(t *testing.T) {

	mem := NewProgram([]Instruction{
		FromMaskString("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		SetMemory{1, 42},
	}).Run()

	assert.Len(t, mem, 1)
	assert.Equal(t, uint64(42), mem[1])

}
