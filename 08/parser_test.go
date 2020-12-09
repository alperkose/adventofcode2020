package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingNop(t *testing.T) {
	source := strings.NewReader("nop +12")

	pr := Parse(source)
	acc, err := pr.Run()
	assert.NoError(t, err)
	assert.Equal(t, 0, acc)
}
func TestParsingAccumulator(t *testing.T) {
	source := strings.NewReader("acc +42")

	pr := Parse(source)

	acc, err := pr.Run()
	assert.NoError(t, err)
	assert.Equal(t, 42, acc)
}
func TestParsingJump(t *testing.T) {
	source := strings.NewReader("jmp +2\nacc +2\nnop +3")

	pr := Parse(source)

	acc, err := pr.Run()
	assert.NoError(t, err)
	assert.Equal(t, 0, acc)
}
func TestParsingAndMovingMultipleJumps(t *testing.T) {
	source := strings.NewReader("jmp +2\njmp +3\nnop +3\njmp -2\nacc -42")

	pr := Parse(source)

	acc, err := pr.Run()
	assert.NoError(t, err)
	assert.Equal(t, -42, acc)
}

func TestDetectAndExitOnInfiniteLoop(t *testing.T) {
	source := strings.NewReader("jmp +2\nacc +3\nnop +3\njmp -2\nacc -42")

	pr := Parse(source)

	acc, err := pr.Run()
	assert.Error(t, err)
	assert.Equal(t, 3, acc)
}

func TestDetectAndFixOnInfiniteLoop(t *testing.T) {
	source := strings.NewReader("jmp +2\nacc +3\nnop +2\njmp -2\nacc -42")

	pr := Parse(source)

	acc, err := pr.FixNRun()
	assert.NoError(t, err)
	assert.Equal(t, -39, acc)
}

func TestDetectAndFixFromExample(t *testing.T) {
	source := strings.NewReader("nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6")

	pr := Parse(source)

	acc, err := pr.FixNRun()
	assert.NoError(t, err)
	assert.Equal(t, 8, acc)
}
