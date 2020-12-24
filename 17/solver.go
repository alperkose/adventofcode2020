package main

type FNeigbors func(ActiveCube) []ActiveCube

type Solver struct {
	space     map[string]ActiveCube
	neighbors FNeigbors
}

func NewSolver(cubes []ActiveCube, neighbors FNeigbors) *Solver {
	space := make(map[string]ActiveCube, len(cubes))
	for _, c := range cubes {
		space[c.String()] = c
	}
	return &Solver{space, neighbors}
}

type CubeWithCount struct {
	c      ActiveCube
	nCount int
}

func (s *Solver) Next() *Solver {
	newSpace := make(map[string]ActiveCube, len(s.space)*2)
	neighbors := map[string]CubeWithCount{}
	// fmt.Println("space", s.space)
	for k, cb := range s.space {
		countNeighbors := 0
		for _, n := range s.neighbors(cb) {
			nStr := n.String()
			_, ok := s.space[nStr]
			if ok {
				countNeighbors++
			} else {
				if nc, ok := neighbors[nStr]; ok {
					nc.nCount += 1
					neighbors[nStr] = nc
				} else {
					neighbors[nStr] = CubeWithCount{n, 1}
				}
			}
		}
		if countNeighbors == 2 || countNeighbors == 3 {
			newSpace[k] = cb
		}
	}
	// fmt.Printf("%+v\n", neighbors)
	for k, n := range neighbors {
		if n.nCount == 3 {
			newSpace[k] = n.c
		}
	}
	s.space = newSpace
	return s
}

func (s *Solver) ActiveCubes() []ActiveCube {
	cubes := make([]ActiveCube, len(s.space))
	ind := 0
	for _, c := range s.space {
		cubes[ind] = c
		ind++
	}
	return cubes
}
