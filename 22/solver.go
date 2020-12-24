package main

import (
	"fmt"
)

type Solver struct {
	player1 []int
	player2 []int
}

func (s *Solver) Combat() []int {
	for len(s.player1) > 0 && len(s.player2) > 0 {
		c1 := s.player1[0]
		c2 := s.player2[0]
		if c1 < c2 {
			s.player2 = append(s.player2[1:], c2, c1)
			s.player1 = s.player1[1:]
		} else {
			s.player1 = append(s.player1[1:], c1, c2)
			s.player2 = s.player2[1:]
		}
	}
	if len(s.player1) < len(s.player2) {
		return s.player2
	}
	return s.player1
}

func (s *Solver) RecursiveCombat() []int {
	rp1, rp2 := recursiveCombatEngine(s.player1, s.player2)
	if len(rp1) < len(rp2) {
		return rp2
	}
	return rp1
}

func recursiveCombatEngine(p1, p2 []int) ([]int, []int) {
	previousTurns := map[string]bool{}
	// fmt.Println("========================================")

	for len(p1) > 0 && len(p2) > 0 {
		// fmt.Printf("(%d)%v;(%d)%v\n", len(p1), p1, len(p2), p2)
		c1 := p1[0]
		c2 := p2[0]
		currTurn := fmt.Sprintf("%v;%v", p1, p2)
		if _, ok := previousTurns[currTurn]; ok {
			// fmt.Println("this turn was played before", currTurn)
			// p1 = append(p1, p2...)
			return p1, []int{}
		} else {
			previousTurns[currTurn] = true
		}
		if c1 < len(p1) && c2 < len(p2) {
			rp1, rp2 := recursiveCombatEngine(append([]int{}, p1[1:c1+1]...), append([]int{}, p2[1:c2+1]...))
			// fmt.Printf("(%d)%v;(%d)%v\n", len(p1), p1, len(p2), p2)

			// fmt.Println("----------------------------------------")
			if len(rp1) < len(rp2) {
				// fmt.Println("Player 2 wins")
				p2 = append(p2[1:], c2, c1)
				p1 = p1[1:]
			} else {
				// fmt.Println("Player 1 wins")
				p1 = append(p1[1:], c1, c2)
				p2 = p2[1:]
			}
		} else if c1 < c2 {
			// fmt.Println("Player 2 wins")
			p2 = append(p2[1:], c2, c1)
			p1 = p1[1:]
		} else {
			// fmt.Println("Player 2 wins")
			p1 = append(p1[1:], c1, c2)
			p2 = p2[1:]
		}
	}
	return p1, p2
}
