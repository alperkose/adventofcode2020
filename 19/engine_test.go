package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesWithSinglePath(t *testing.T) {
	engine := NewEngine([]Rule{
		SinglePathRule(0, []int{1, 2}),
		ValueRule(1, 'a'),
		ValueRule(2, 'b'),
	})

	assert.True(t, engine.Matches("ab"))
	assert.False(t, engine.Matches("ba"))
	assert.False(t, engine.Matches("aba"))
	assert.False(t, engine.Matches("aa"))
	assert.False(t, engine.Matches("bb"))
}
func TestMatchesWithDoublePath(t *testing.T) {
	engine := NewEngine([]Rule{
		SinglePathRule(0, []int{1, 2}),
		ValueRule(1, 'a'),
		DoublePathRule(2, []int{1, 3}, []int{3, 1}),
		ValueRule(3, 'b'),
	})

	// assert.True(t, engine.Matches("aab"))
	// assert.True(t, engine.Matches("aba"))
	assert.False(t, engine.Matches("aa"))
	// assert.False(t, engine.Matches("abb"))
	// assert.False(t, engine.Matches("baa"))

}
