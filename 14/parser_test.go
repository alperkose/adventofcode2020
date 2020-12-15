package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingInput(t *testing.T) {

	var inputStr = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	var program *Program = Parse(strings.NewReader(inputStr), Part1)
	mem := program.Run()
	assert.Len(t, mem, 2)
	assert.Equal(t, uint64(101), mem[7])
	assert.Equal(t, uint64(64), mem[8])
}
