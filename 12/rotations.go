package main

import "strings"

const headings = "ENWS"

type RotateLeft int

func (l RotateLeft) Apply(ship Ship) Ship {
	amount := int(l) / 90
	curr := strings.Index(headings, string(ship.Heading))
	ship.Heading = Heading(string(headings[(curr+amount)%4]))
	return ship
}

type RotateRight int

func (r RotateRight) Apply(ship Ship) Ship {
	amount := int(r) / 90
	curr := strings.Index(headings, string(ship.Heading))
	ship.Heading = Heading(string(headings[(curr-amount+4)%4]))
	return ship
}

var wpRotations = []struct {
	rotate func(Ship) Ship
}{
	{
		func(s Ship) Ship {
			s.WpX, s.WpY = s.WpY*-1, s.WpX
			return s
		},
	}, {
		func(s Ship) Ship {
			s.WpX, s.WpY = s.WpX*-1, s.WpY*-1
			return s
		},
	}, {
		func(s Ship) Ship {
			s.WpX, s.WpY = s.WpY, s.WpX*-1
			return s
		},
	},
}

type RotateWpLeft int

func (l RotateWpLeft) Apply(ship Ship) Ship {
	amount := (int(l) / 90) - 1
	return wpRotations[amount].rotate(ship)
}

type RotateWpRight int

func (r RotateWpRight) Apply(ship Ship) Ship {
	amount := 3 - (int(r) / 90)
	return wpRotations[amount].rotate(ship)
}
