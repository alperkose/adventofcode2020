package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingCommands(t *testing.T) {
	source := strings.NewReader("F10\nN3\nF7\nR90\nF11")
	instructions := Parse(source)
	assert.Len(t, instructions, 5)
}
