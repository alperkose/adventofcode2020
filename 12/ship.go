package main

import "fmt"

type Heading string

const East = Heading("E")
const North = Heading("N")
const West = Heading("W")
const South = Heading("S")

type Ship struct {
	Heading  Heading
	X, Y     int
	WpX, WpY int
}

func (s Ship) String() string {
	return fmt.Sprintf("%s(%d, %d)-Wp(%d, %d)", s.Heading, s.X, s.Y, s.WpX, s.WpY)
}
