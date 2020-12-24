package main

import "fmt"

type Engine struct {
	rules []Rule
}

func NewEngine(rules []Rule) *Engine {
	return &Engine{rules}
}

func (e *Engine) Matches(line string) bool {
	match, matchedChars := e.startFrom(0, line)
	return match && len(line) == matchedChars
}

func (e *Engine) startFrom(ruleInd int, str string) (bool, int) {
	fmt.Printf("%3d) Trying rule : %+v on %s\n", ruleInd, e.rules[ruleInd], str)
	if len(str) == 0 {
		return true, 0
	}
	rule := e.rules[ruleInd]
	if !rule.HasNext() {
		return str[0] == byte(rule.val), 1
	}

	match, matchedChars := e.goToPath(rule.next1, str)
	fmt.Printf("%3d) match %v matchedCars %d\n", rule.order, match, matchedChars)

	if !match && len(rule.next2) > 0 {
		match, matchedChars = e.goToPath(rule.next2, str)
	}

	return match, matchedChars
}

func (e *Engine) goToPath(path []int, str string) (bool, int) {
	matchedChars := 0
	var match bool
	var nmatched int
	for i := 0; i < len(path); i++ {
		match, nmatched = e.startFrom(path[i], str[matchedChars:])
		if !match {
			break
		}
		matchedChars += nmatched
	}
	return match, matchedChars
}

func contains(list []int, n int) bool {
	for _, item := range list {
		if item == n {
			return true
		}
	}
	return false
}
