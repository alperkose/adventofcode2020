package main

type Rule struct {
	order int
	next1 []int
	next2 []int
	val   rune
}

func ValueRule(order int, val rune) Rule {
	return Rule{order, []int{}, []int{}, val}
}
func SinglePathRule(order int, n1 []int) Rule {
	return Rule{order, n1, []int{}, '-'}
}
func DoublePathRule(order int, n1, n2 []int) Rule {
	return Rule{order, n1, n2, '-'}
}

func (r Rule) HasNext() bool {
	return len(r.next1) > 0 || len(r.next2) > 0
}
