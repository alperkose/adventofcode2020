package main

type Forward int

func (f Forward) Apply(ship Ship) Ship {
	switch ship.Heading {
	case East:
		ship.X += int(f)
	case North:
		ship.Y += int(f)
	case West:
		ship.X -= int(f)
	case South:
		ship.Y -= int(f)
	}
	return ship
}

type ForwardByWp int

func (f ForwardByWp) Apply(ship Ship) Ship {
	ship.X += int(f) * ship.WpX
	ship.Y += int(f) * ship.WpY
	return ship
}
