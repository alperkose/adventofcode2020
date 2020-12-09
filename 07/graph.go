package main

import "fmt"

type Rules struct {
	bags        map[string]Bag
	containedIn map[string][]string
}

func NewRules() *Rules {
	return &Rules{map[string]Bag{}, map[string][]string{}}
}

func (r *Rules) Add(b Bag) {
	r.bags[b.color] = b
	for cc := range b.canContain {
		currCnt, ok := r.containedIn[cc]
		if !ok {
			r.containedIn[cc] = []string{b.color}
		} else {
			r.containedIn[cc] = append(currCnt, b.color)
		}
	}
}

func (r *Rules) BagsCanContain(color string) []string {
	canContain := map[string]bool{}
	colorsToCheck := []string{color}
	var currentColor string
	for {
		if len(colorsToCheck) == 0 {
			break
		}
		currentColor = colorsToCheck[0]
		colorsToCheck = colorsToCheck[1:]

		if _, ok := canContain[currentColor]; ok {
			continue
		}

		canContain[currentColor] = true
		cntIn, ok := r.containedIn[currentColor]
		if !ok {
			continue
		}
		for _, toCheck := range cntIn {
			colorsToCheck = append(colorsToCheck, toCheck)
		}
	}
	delete(canContain, color)
	return keys(canContain)
}

func (r *Rules) CountBagsWithin(color string) int {
	sum := 0
	curr := r.bags[color]
	for k, v := range curr.canContain {
		fmt.Println(k, v)
		sum += v * (r.CountBagsWithin(k) + 1)
	}
	return sum
}

func keys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Bag struct {
	color      string
	canContain map[string]int
}
type Content struct {
	color  string
	amount int
}

func NewBag(color string, content ...Content) Bag {
	canContain := make(map[string]int, len(content))
	for _, c := range content {
		canContain[c.color] = c.amount
	}
	return Bag{color, canContain}
}
