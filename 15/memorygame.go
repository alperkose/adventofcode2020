package main

func NewGame(input []int) *Game {
	mem := make(map[int][]int, len(input))
	for i, v := range input {
		l, ok := mem[v]
		if !ok {
			mem[v] = []int{i}
		} else {
			mem[v] = append(l, i)
		}
	}
	return &Game{input, mem}
}

type Game struct {
	input []int
	mem   map[int][]int
}

func (g *Game) Guess(n int) int {
	last := g.input[len(g.input)-1]
	var curr, lenLastLocs int
	// fmt.Printf("Inputs   %v\n", g.input)
	for i := len(g.input); i < n; i++ {
		l := g.mem[last]
		// fmt.Printf("%4d) %4d %v\n", i, last, l)
		lenLastLocs = len(l)
		if lenLastLocs == 1 {
			curr = 0
		} else {
			curr = l[lenLastLocs-1] - l[lenLastLocs-2]
		}
		l, ok := g.mem[curr]
		if !ok {
			g.mem[curr] = []int{i}
		} else {
			g.mem[curr] = append(l, i)
		}
		// fmt.Printf("      %4d %v\n", curr, l)
		last = curr
	}
	return curr
}
