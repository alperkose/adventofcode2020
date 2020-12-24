package main

type Cruncher struct {
	ranges       []Range
	ticket       []int
	otherTickets [][]int
}

type Range struct {
	begin, end int
	field      string
}

func (r Range) inRange(n int) bool {
	return r.begin <= n && n <= r.end
}
func NewCruncher() *Cruncher {
	return &Cruncher{
		[]Range{},
		[]int{},
		[][]int{},
	}
}

func (c *Cruncher) Ranges() []Range {
	return c.ranges
}

func (c *Cruncher) Ticket() []int {
	return c.ticket
}
func (c *Cruncher) OtherTickets() [][]int {
	return c.otherTickets
}

func (c *Cruncher) AddRanges(r ...Range) {
	c.ranges = append(c.ranges, r...)
}
func (c *Cruncher) SetTicket(ticket []int) {
	c.ticket = ticket
}
func (c *Cruncher) AddOtherTicket(ticket []int) {
	c.otherTickets = append(c.otherTickets, ticket)
}

func (c *Cruncher) SumOfInvalidNumbers() int {
	var t []int
	var sum int = 0

	for i := 0; i < len(c.otherTickets); i++ {
		t = c.otherTickets[i]
		_, v := isValidTicket(t, c.ranges)
		sum += v
	}
	return sum
}

func isValidTicket(t []int, ranges []Range) (bool, int) {
	for j := 0; j < len(t); j++ {
		tn := t[j]
		ok := false
		for k := 0; k < len(ranges); k++ {
			rng := ranges[k]
			if rng.inRange(tn) {
				ok = true
			}
		}
		if !ok {
			return false, tn
		}
	}
	return true, 0
}

func (c *Cruncher) FindFieldOrder() []string {

	fields := make(map[string][]Range, len(c.ticket))
	for _, rng := range c.ranges {
		r, ok := fields[rng.field]
		if ok {
			fields[rng.field] = append(r, rng)
		} else {
			fields[rng.field] = []Range{rng}
		}
	}

	orderedFieldMaps := make([]map[string]bool, len(c.ticket))
	for colInd := 0; colInd < len(c.ticket); colInd++ {
		remainingFields := make(map[string]bool, len(fields))
		for k, _ := range fields {
			remainingFields[k] = true
		}

		for j := 0; j < len(c.otherTickets); j++ {
			tckt := c.otherTickets[j]
			if ok, _ := isValidTicket(tckt, c.ranges); !ok {
				// fmt.Printf("%d)\tNot a valid ticket (%d)\n", colInd, num)
				continue
			}
			for fName, _ := range remainingFields {
				// fmt.Printf("%d)\tChecking %s tckt %d\n", colInd, fName, tckt[colInd])
				ok := false
				for _, rng := range fields[fName] {
					// fmt.Printf("%d)\t%d is in range %+v?\n", colInd, tckt[colInd], rng)
					if rng.inRange(tckt[colInd]) {
						// fmt.Printf("%d)\t%s is in range\n", colInd, fName)
						ok = true
						break
					}
				}
				if !ok {
					// fmt.Printf("%d)\tDeleting %s from remainingFields\n", colInd, fName)
					delete(remainingFields, fName)
				}
			}
			if len(remainingFields) == 1 {
				for k, _ := range remainingFields {
					delete(fields, k)
				}
				break
			}
		}
		orderedFieldMaps[colInd] = map[string]bool{}
		// fmt.Println(colInd, "remaining", remainingFields)
		// fmt.Println()
		for k, _ := range remainingFields {
			orderedFieldMaps[colInd][k] = true
		}
	}

	// fmt.Printf("%#+v\n", orderedFieldMaps)
	orderedFields := make([]string, len(orderedFieldMaps))
	for i := 0; i < 100; i++ {
		for ofmInd, ofm := range orderedFieldMaps {

			if len(ofm) == 1 {
				var keyToDelete string
				for keyToDelete, _ = range ofm {
				}
				// fmt.Printf("%v is a single item %s, deleting from other maps\n", ofm, keyToDelete)

				for j, ofm2 := range orderedFieldMaps {
					if ofmInd == j {
						continue
					}
					// fmt.Printf("\tdeleting %s from %v\n", keyToDelete, ofm2)
					delete(ofm2, keyToDelete)
				}
				orderedFields[ofmInd] = keyToDelete
			}

			// for keyToCheck, _ := range ofm {
			// 	existsInOtherMaps := false
			// 	for j, ofm2 := range orderedFieldMaps {
			// 		if ofmInd == j {
			// 			continue
			// 		}
			// 		_, ok := ofm2[keyToCheck]
			// 		if ok {
			// 			existsInOtherMaps = true
			// 			break
			// 		}
			// 	}
			// 	if !existsInOtherMaps {
			// 		orderedFieldMaps[ofmInd] = map[string]bool{keyToCheck: true}
			// 		orderedFields[ofmInd] = keyToCheck
			// 		break
			// 	}
			// }
		}
	}

	return orderedFields
}
