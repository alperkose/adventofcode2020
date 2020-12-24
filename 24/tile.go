package main

func TileReferenceFromString(input string) TileReference {
	locators := []Locator{}
	for i := 0; i < len(input); i++ {
		curr := input[i]
		switch curr {
		case 'e':
			locators = append(locators, LocatorFromRef(string(curr)))
		case 's':
			if input[i+1] == 'e' {
				locators = append(locators, LocatorFromRef("se"))
			} else {
				locators = append(locators, LocatorFromRef("sw"))
			}
			i++
		case 'w':
			locators = append(locators, LocatorFromRef(string(curr)))
		case 'n':
			if input[i+1] == 'e' {
				locators = append(locators, LocatorFromRef("ne"))
			} else {
				locators = append(locators, LocatorFromRef("nw"))
			}
			i++
		}
	}
	return TileReference{input, locators}
}

type TileReference struct {
	ref      string
	locators []Locator
}

func (tr TileReference) LocateFromReference(x, y int) (int, int) {
	var locX, locY int = x, y
	for _, locate := range tr.locators {
		locX, locY = locate(locX, locY)
	}
	return locX, locY
}
