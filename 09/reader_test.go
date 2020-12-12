package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadingIntegerArray(t *testing.T) {
	source := strings.NewReader("1\n3\n5\n8")

	inputs := ReadSource(source)

	assert.Equal(t, []int{1, 3, 5, 8}, inputs)
}

func TestReadingIntegerArrayIgnoresEmptyLines(t *testing.T) {
	source := strings.NewReader("1\n2\n\n3\n")

	inputs := ReadSource(source)

	assert.Equal(t, []int{1, 2, 3}, inputs)
}
