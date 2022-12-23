package problem23

func Part2(lines []string) int {
	mat := newSparseMatrix()
	for r, line := range lines {
		for c, elem := range line {
			if elem == '.' {
				// nop
			} else if elem == '#' {
				mat.addCoord(newCoord(r, c))
			} else {
				panic("wrong elem")
			}
		}
	}

	directionFuncs := []func(coord, bool, bool, bool, bool, bool, bool, bool, bool) (bool, coord){
		checkNorth,
		checkSouth,
		checkWest,
		checkEast,
	}
	directionIndex := 0

	round := 0
	for {
		round++
		proposals := make(map[coord]coord) // proposal > who proposed it
		for k, _ := range mat.coords {
			nw := mat.contains(newCoord(k.r-1, k.c-1))
			n := mat.contains(newCoord(k.r-1, k.c))
			ne := mat.contains(newCoord(k.r-1, k.c+1))
			e := mat.contains(newCoord(k.r, k.c+1))
			se := mat.contains(newCoord(k.r+1, k.c+1))
			s := mat.contains(newCoord(k.r+1, k.c))
			sw := mat.contains(newCoord(k.r+1, k.c-1))
			w := mat.contains(newCoord(k.r, k.c-1))

			if !nw && !n && !ne && !e && !se && !s && !sw && !w {
				continue
			}

			for i := 0; i < 4; i++ {
				checkFunc := directionFuncs[(directionIndex+i)%4]
				ok, proposedPos := checkFunc(k, nw, n, ne, e, se, s, sw, w)
				if ok {
					_, exists := proposals[proposedPos]
					if exists {
						// remove previous proposal and don't propose any current change
						delete(proposals, proposedPos)
					} else {
						proposals[proposedPos] = k
					}
					break
				}
			}
		}
		directionIndex++

		if len(proposals) == 0 {
			break
		}
		// apply remaining proposals
		for proposedPos, who := range proposals {
			mat.addCoord(proposedPos)
			mat.delCoord(who)
		}
	}

	// fmt.Println(round)
	// fmt.Println(mat)

	return round
}

func checkNorth(pos coord, nw, n, ne, e, se, s, sw, w bool) (ok bool, proposal coord) {
	if !nw && !n && !ne {
		return true, pos.north()
	}
	return false, coord{}
}

func checkSouth(pos coord, nw, n, ne, e, se, s, sw, w bool) (ok bool, proposal coord) {
	if !sw && !s && !se {
		return true, pos.south()
	}
	return false, coord{}
}

func checkWest(pos coord, nw, n, ne, e, se, s, sw, w bool) (ok bool, proposal coord) {
	if !nw && !w && !sw {
		return true, pos.west()
	}
	return false, coord{}
}

func checkEast(pos coord, nw, n, ne, e, se, s, sw, w bool) (ok bool, proposal coord) {
	if !ne && !e && !se {
		return true, pos.east()
	}
	return false, coord{}
}
