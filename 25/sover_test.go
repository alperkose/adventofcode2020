package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLoopSize(t *testing.T) {
	assert.Equal(t, FindLoopSize(7, 5764801), 8)
	assert.Equal(t, FindLoopSize(7, 17807724), 11)
}

func TestTransforming(t *testing.T) {
	assert.Equal(t, Transform(17807724, 8), 14897079)
	assert.Equal(t, Transform(5764801, 11), 14897079)
}
