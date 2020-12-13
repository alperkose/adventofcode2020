package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingNop(t *testing.T) {
	source := strings.NewReader(".L..L.\nLLL...")

	rows := ReadSource(source)

	assert.Len(t, rows, 2)
}

func TestTrimEmptySpace(t *testing.T) {
	source := strings.NewReader(`..LL..
								##LL..
								......`)

	rows := ReadSource(source)

	assert.Equal(t, []string{
		"..LL..",
		"##LL..",
		"......",
	}, rows)
}
