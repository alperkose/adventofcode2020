package main

func Solve(cups []int, numberOfMoves int) []int {
	// fmt.Println("solving for", cups)
	// state := map[string]int{}
	// defer func() {
	// 	fmt.Println("state len", len(state))
	// }()
	current, index := toRing(cups)
	maxCup := len(cups)
	// start := time.Now()
	for i := 0; i < numberOfMoves; i++ {
		// stateStr := fmt.Sprint(allToArray(current))
		// if previousOccurence, ok := state[stateStr]; ok {
		// 	fmt.Printf("%3d. previous occurence %d state %s", i, previousOccurence, stateStr)
		// } else {
		// 	state[stateStr] = i
		// }
		// if i%1000 == 0 {
		// fmt.Printf("%3d. current %d, next: %d (%p), prev %d (%p) [%.2fm]\n", i, current.val, current.next.val, current.next, current.prev.val, current.prev, time.Now().Sub(start).Minutes())
		// }
		cVal := current.val

		destVal := correctVal(cVal-1, maxCup)
		for k := 1; k <= 3; k++ {
			// fmt.Printf("%3d. is dest (%d) picked %d, %d, %d\n", i, destVal, current.next.val, current.next.next.val, current.next.next.next.val)
			if destVal != current.next.val &&
				destVal != current.next.next.val &&
				destVal != current.next.next.next.val {
				break
			}
			destVal = correctVal(destVal-1, maxCup)
		}
		// candidateFwd := current.next.next.next.next
		// candidateBwd := current.prev
		// fmt.Printf("%3d. dest candidate %d\n", i, destVal)
		var destination *CupInARing
		// for candidateFwd.val != cVal {
		// fmt.Printf("%3d. candidates fwd:%d bwd:%d bwdNxt:%d\n", i, candidateFwd.val, candidateBwd.val, candidateBwd.next.val)
		destination = index[destVal]
		// for {
		// 	if candidateFwd.val == destVal {
		// 		destination = candidateFwd
		// 		break
		// 	}
		// 	if candidateBwd.val == destVal {
		// 		destination = candidateBwd
		// 		break
		// 	}
		// 	if candidateBwd.val == candidateFwd.next.val {
		// 		// will panic if this happens
		// 		break
		// 	}
		// 	candidateFwd = candidateFwd.next
		// 	candidateBwd = candidateBwd.prev
		// 	// fmt.Printf("%3d. candidates fwd:%d bwd:%d bwdNxt:%d\n", i, candidateFwd.val, candidateBwd.val, candidateBwd.next.val)
		// }

		if destination == nil {
			panic("Whaaaat!!! Couldn't find destination.")
		}

		reArrangeRing(current, destination)
		current = current.next
		// fmt.Printf("%3d. status %v\n", i, allToArray(current))
	}

	return toArray(current, 1, len(cups)-1)
}

func correctVal(val, max int) int {
	if val <= 0 {
		return val + max
	}
	return val
}

func reArrangeRing(current, destination *CupInARing) {
	firstPicked := current.next
	thirdPicked := current.next.next.next
	// 1 2 3 4 . . 7 8
	current.next = thirdPicked.next
	current.next.prev = current

	thirdPicked.next = destination.next
	thirdPicked.next.prev = thirdPicked

	destination.next = firstPicked
	firstPicked.prev = destination

}

func toArray(start *CupInARing, startFromValue, size int) []int {
	result := make([]int, size)
	for start.val != startFromValue {
		start = start.next
	}
	start = start.next
	for i := 0; i < size; i++ {
		result[i] = start.val
		start = start.next
	}
	return result
}
func allToArray(start *CupInARing) []int {
	result := []int{start.val}
	for start = start.next; start.val != result[0]; start = start.next {
		result = append(result, start.val)
	}

	return result
}

func toRing(cups []int) (*CupInARing, map[int]*CupInARing) {

	first := &CupInARing{cups[0], nil, nil}
	easyAccessMap := map[int]*CupInARing{first.val: first}
	var prev = first
	var curr *CupInARing
	for i := 1; i < len(cups); i++ {
		curr = &CupInARing{cups[i], nil, prev}
		easyAccessMap[cups[i]] = curr
		prev.next = curr
		prev = curr
	}
	prev.next = first
	first.prev = prev
	return first, easyAccessMap
}

type CupInARing struct {
	val  int
	next *CupInARing
	prev *CupInARing
}
