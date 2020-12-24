package main

type Operation interface {
	F(int, int) int
	Precedence() int
}

type Addition struct{}

func (o Addition) F(a, b int) int {
	return a + b
}
func (o Addition) Precedence() int {
	return 10
}

type Multiplication struct{}

func (o Multiplication) F(a, b int) int {
	return a * b
}
func (o Multiplication) Precedence() int {
	return 1
}
