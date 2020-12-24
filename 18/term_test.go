package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueTerm(t *testing.T) {
	assert.Equal(t, 5, Val(5).Value())
	assert.Equal(t, -1, Val(-1).Value())
}

func TestVariableTerm(t *testing.T) {
	assert.Equal(t, 11, Var{Val(5), Val(6), Addition{}}.Value())
	assert.Equal(t, 30, Var{Val(5), Val(6), Multiplication{}}.Value())
}

func TestAddingNewTerm(t *testing.T) {
	assert.Equal(t, 7, InitialTerm{}.Add(Val(7)).Value())
	assert.Equal(t, 12, Val(7).Add(Val(5)).Value())
	assert.Equal(t, 15, Var{Val(5), Val(7), Addition{}}.Add(Val(3)).Value())

	assert.Equal(t, 50, Var{Val(5), Val(7), Multiplication{}}.Add(Val(3)).Value())
}
func TestMultiplyingNewTerm(t *testing.T) {
	assert.Equal(t, 35, Val(7).Mult(Val(5)).Value())
	assert.Equal(t, 105, Var{Val(5), Val(7), Multiplication{}}.Mult(Val(3)).Value())

	assert.Equal(t, 36, Var{Val(5), Val(7), Addition{}}.Mult(Val(3)).Value())
}

func TestParanthesisTerm(t *testing.T) {
	assert.Equal(t, 7, Paranthesis{Val(7)}.Value())
	assert.Equal(t, 105, Paranthesis{Var{Val(5), Val(7), Multiplication{}}}.Mult(Val(3)).Value())
	assert.Equal(t, 36, Paranthesis{Var{Val(5), Val(7), Addition{}}}.Mult(Val(3)).Value())
	assert.Equal(t, 38, Paranthesis{Var{Val(5), Val(7), Multiplication{}}}.Add(Val(3)).Value())
}
