package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingRules(t *testing.T) {

	source := `0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"

aab
baa
aba`
	rules, samples := Parse(strings.NewReader(source))

	assert.Len(t, rules, 4)
	assert.True(t, rules[0].HasNext())
	assert.False(t, rules[1].HasNext())
	assert.True(t, rules[2].HasNext())
	assert.False(t, rules[3].HasNext())

	assert.Equal(t, []string{"aab", "baa", "aba"}, samples)
}

func TestParsingRulesWithOrder(t *testing.T) {

	source := `2: 1 3 | 3 1
	0: 1 2	
1: "a"
3: "b"

aab
baa
aba`
	rules, samples := Parse(strings.NewReader(source))

	assert.Len(t, rules, 4)
	assert.True(t, rules[0].HasNext())
	assert.False(t, rules[1].HasNext())
	assert.True(t, rules[2].HasNext())
	assert.False(t, rules[3].HasNext())
	assert.Equal(t, []string{"aab", "baa", "aba"}, samples)

}

func TestParsingSingleBlock(t *testing.T) {
	rule := ParseLine("0: 1 2")
	assert.Equal(t, 0, rule.order)
	assert.Equal(t, []int{1, 2}, rule.next1)
	assert.Equal(t, []int{}, rule.next2)
	assert.True(t, rule.HasNext())
}
func TestParsingDoubleBlock(t *testing.T) {
	rule := ParseLine("2: 1 3 | 3 1")
	assert.Equal(t, 2, rule.order)
	assert.Equal(t, []int{1, 3}, rule.next1)
	assert.Equal(t, []int{3, 1}, rule.next2)
	assert.True(t, rule.HasNext())
}
func TestParsingValue(t *testing.T) {
	rule := ParseLine("3: \"b\"")
	assert.Equal(t, 3, rule.order)
	assert.Equal(t, 'b', rule.val)
	assert.False(t, rule.HasNext())
	assert.Equal(t, []int{}, rule.next1)
	assert.Equal(t, []int{}, rule.next2)
}
