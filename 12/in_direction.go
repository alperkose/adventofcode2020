package main

type MoveEast int

func (d MoveEast) Apply(ship Ship) Ship {
	ship.X += int(d)
	return ship
}

type MoveNorth int

func (d MoveNorth) Apply(ship Ship) Ship {
	ship.Y += int(d)
	return ship
}

type MoveWest int

func (d MoveWest) Apply(ship Ship) Ship {
	ship.X -= int(d)
	return ship
}

type MoveSouth int

func (d MoveSouth) Apply(ship Ship) Ship {
	ship.Y -= int(d)
	return ship
}

// Wapypoint operations

type MoveWpEast int

func (d MoveWpEast) Apply(ship Ship) Ship {
	ship.WpX += int(d)
	return ship
}

type MoveWpNorth int

func (d MoveWpNorth) Apply(ship Ship) Ship {
	ship.WpY += int(d)
	return ship
}

type MoveWpWest int

func (d MoveWpWest) Apply(ship Ship) Ship {
	ship.WpX -= int(d)
	return ship
}

type MoveWpSouth int

func (d MoveWpSouth) Apply(ship Ship) Ship {
	ship.WpY -= int(d)
	return ship
}
