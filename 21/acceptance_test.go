package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWith4Foods(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	solver := Parse(strings.NewReader(input))
	ingredients := solver.NonAllergenicIngredients()
	assert.ElementsMatch(t, []string{"kfcds", "nhms", "sbzzf", "trh"}, ingredients)
}
