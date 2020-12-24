package main

type Term interface {
	Value() int
	Add(Term) Term
	Mult(Term) Term
}

type InitialTerm struct {
}

func (t InitialTerm) Value() int {
	return -1
}

func (t InitialTerm) Add(nt Term) Term {
	return nt
}
func (t InitialTerm) Mult(nt Term) Term {
	return nt
}

type Val int

func (t Val) Value() int {
	return int(t)
}

func (t Val) Add(nt Term) Term {
	return Var{t, nt, Addition{}}
}
func (t Val) Mult(nt Term) Term {
	return Var{t, nt, Multiplication{}}
}

type Var struct {
	left, right Term
	op          Operation
}

func (t Var) Value() int {
	return t.op.F(t.left.Value(), t.right.Value())
}

func (t Var) Add(nt Term) Term {
	return Var{t.left, Var{t.right, nt, Addition{}}, t.op}
	// return Var{t, nt, Addition{}}
}
func (t Var) Mult(nt Term) Term {
	return Var{t, nt, Multiplication{}}
}

type Paranthesis struct {
	term Term
}

func (t Paranthesis) Value() int {
	return t.term.Value()
}

func (t Paranthesis) Add(nt Term) Term {
	return Var{t.term, nt, Addition{}}
}
func (t Paranthesis) Mult(nt Term) Term {
	return Var{t.term, nt, Multiplication{}}
}
