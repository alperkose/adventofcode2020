package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const sep_contains string = " (contains "

func Parse(source io.Reader) *Solver {
	buffer := bufio.NewReader(source)
	foods := []Food{}
	for {
		line, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lineStr := string(line)
		ingredientEnd := strings.Index(lineStr, sep_contains)
		ingredients := strings.Split(lineStr[:ingredientEnd], " ")
		allergenEnd := strings.LastIndex(lineStr, ")")
		allergens := strings.Split(lineStr[ingredientEnd+len(sep_contains):allergenEnd], ", ")
		foods = append(foods, Food{ingredients, allergens})
	}
	return &Solver{foods}
}

type Solver struct {
	food []Food
}

type Food struct {
	ingredients []string
	allergens   []string
}

func (s *Solver) NonAllergenicIngredients() []string {

	// fmt.Println(">> food", s.food)

	nonAllergicIngredientCounts := map[string]int{}
	for _, food := range s.food {
		for _, ingredient := range food.ingredients {
			count, ok := nonAllergicIngredientCounts[ingredient]
			if !ok {
				nonAllergicIngredientCounts[ingredient] = 1
			} else {
				nonAllergicIngredientCounts[ingredient] = count + 1
			}
		}
	}

	allergenToIngredients := map[string]map[string]bool{}
	for _, food := range s.food {
		// fmt.Println("\tfood", food)
		for _, allergen := range food.allergens {
			// fmt.Println("\t\tallergen", allergen)
			ingredientSet, ok := allergenToIngredients[allergen]
			if !ok {
				freshMap := map[string]bool{}
				for _, ingredient := range food.ingredients {
					// fmt.Println("\t\t\tingredient", ingredient)
					freshMap[ingredient] = true
				}
				allergenToIngredients[allergen] = freshMap
				continue
			}
			intersection := map[string]bool{}
			for _, ingredient := range food.ingredients {
				if _, ok := ingredientSet[ingredient]; ok {
					// fmt.Println("\t\t\tingredient", ingredient)
					intersection[ingredient] = true
				}
			}
			allergenToIngredients[allergen] = intersection
		}
	}

	for i := 0; i < 100; i++ {
		for allergen, ingredients := range allergenToIngredients {

			if len(ingredients) == 1 {
				var keyToDelete string
				for keyToDelete, _ = range ingredients {
				}
				delete(nonAllergicIngredientCounts, keyToDelete)
				// fmt.Printf("%v is a single item %s, deleting from other maps\n", ofm, keyToDelete)

				for alrg, ingr := range allergenToIngredients {
					if alrg == allergen {
						continue
					}
					// fmt.Printf("\tdeleting %s from %v\n", keyToDelete, ofm2)
					delete(ingr, keyToDelete)
				}
			}
		}
	}
	fmt.Println(">> allergenToIngredients", allergenToIngredients)

	nonAllergicIngredients := []string{}
	count := 0
	for ingredient, cnt := range nonAllergicIngredientCounts {
		nonAllergicIngredients = append(nonAllergicIngredients, ingredient)
		count += cnt
	}
	fmt.Println(">> nonAllergicIngredients", nonAllergicIngredients)
	fmt.Println(">> count", count)
	return nonAllergicIngredients
}
