package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePlayers(t *testing.T) {
	s := Parse(strings.NewReader(`Player 1:
4
2
10

Player 2:
5
3
7`))
	assert.Equal(t, []int{4, 2, 10}, s.player1)
	assert.Equal(t, []int{5, 3, 7}, s.player2)
}
