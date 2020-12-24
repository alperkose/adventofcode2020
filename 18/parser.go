package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(source io.Reader) []Term {
	buffer := bufio.NewReader(source)
	terms := []Term{}
	for {
		line, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		term, _ := p.ParseTerm(string(line))
		terms = append(terms, term)
	}
	return terms
}

func (p *Parser) ParseTerm(inp string) (Term, int) {
	// fmt.Println("Parsing", inp)
	var term Term = InitialTerm{}
	var op byte = '+'
	var c byte
	for ind := 0; ind < len(inp); ind++ {
		c = inp[ind]
		if c >= '0' && c <= '9' {
			var num int
			endInd := strings.IndexAny(inp[ind:], " +*()")
			if endInd == -1 {
				num, _ = strconv.Atoi(inp[ind:])
				ind = len(inp)
			} else {
				num, _ = strconv.Atoi(inp[ind : ind+endInd])
				ind += endInd - 1
			}
			if op == '+' {
				term = term.Add(Val(num))
			} else {
				term = term.Mult(Val(num))
			}
			// fmt.Printf("%d) %c -> found num %d (op: %c, term: %#+v)\n", ind, c, num, op, term)
		} else if c == '+' || c == '*' {
			op = c
		} else if c == '(' {
			// fmt.Printf("%d) Paranthesis opened. (term:%#+v, op:%c)\n", ind, term, op)
			newTerm, parsedLen := p.ParseTerm(inp[ind+1:])
			if op == '+' {
				term = term.Add(Paranthesis{newTerm})
			} else {
				term = term.Mult(Paranthesis{newTerm})
			}
			ind += parsedLen
			// fmt.Printf("%d) After param. (term:%#+v, op:%c)\n", ind, term, op)
		} else if c == ')' {
			// fmt.Printf("%d) Paranthesis closed. (term:%#+v, op:%c)\n", ind, term, op)
			return term, ind + 1
		}
	}
	return term, len(inp)
}
