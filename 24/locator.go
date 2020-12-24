package main

type Locator func(x, y int) (int, int)

func LocatorFromRef(ref string) Locator {
	switch ref {
	case "e":
		return ToEast
	case "se":
		return ToSouthEast
	case "sw":
		return ToSouthWest
	case "w":
		return ToWest
	case "nw":
		return ToNorthWest
	case "ne":
		return ToNorthEast
	}
	panic("Illegal direction")
}

func ToEast(x, y int) (int, int) {
	return x + 1, y
}
func ToSouthEast(x, y int) (int, int) {
	return x, y - 1
}
func ToSouthWest(x, y int) (int, int) {
	return x - 1, y - 1
}
func ToWest(x, y int) (int, int) {
	return x - 1, y
}
func ToNorthWest(x, y int) (int, int) {
	return x, y + 1
}
func ToNorthEast(x, y int) (int, int) {
	return x + 1, y + 1
}
