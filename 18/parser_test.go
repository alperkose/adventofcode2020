package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingLine(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			input:    "3 + 4",
			expected: 7,
		}, {
			input:    "3 * 4",
			expected: 12,
		}, {
			input:    "3 * 4 * 5",
			expected: 60,
		}, {
			input:    "3 + 4 + 5",
			expected: 12,
		}, {
			input:    "3 + 4 * 5",
			expected: 35,
		}, {
			input:    "3 * 4 + 5",
			expected: 27,
		}, {
			input:    "(3 * 4) + 5",
			expected: 17,
		}, {
			input:    "5 + (3 * 4) + 5",
			expected: 22,
		}, {
			input:    "(5 *5) + (3 * 4)",
			expected: 37,
		}, {
			input:    "5 * (5 + (3 * 4))",
			expected: 85,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			term, _ := NewParser().ParseTerm(tC.input)
			assert.Equal(t, tC.expected, term.Value())
		})
	}
}

func TestParseSource(t *testing.T) {
	source := `(3 * 4) + 5
	3 * 4 + 5
	2 * 3 + (4 * 5)
	5 + (8 * 3 + 9 + 3 * 4 * 3)
	5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
	((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`
	terms := NewParser().Parse(strings.NewReader(source))
	assert.Len(t, terms, 6)
	assert.Equal(t, 17, terms[0].Value())
	assert.Equal(t, 27, terms[1].Value())
	assert.Equal(t, 26, terms[2].Value())
	assert.Equal(t, 437, terms[3].Value())
	assert.Equal(t, 12240, terms[4].Value())
	assert.Equal(t, 13632, terms[5].Value())
}
