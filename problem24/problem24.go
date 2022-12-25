package problem24

import (
	"fmt"
	"math"
)

// var lcm int = 12

var lcm int = 600

func Part1(lines []string) int {
	lineRows := len(lines)
	lineCols := len(lines[0])
	bo := newBoard(lineRows-2, lineCols-2)
	var startCol, endCol int
	for r, line := range lines {

		if r == 0 {
			for c, run := range line {
				if run == '.' {
					startCol = c - 1
					continue
				}
			}
		}

		if r == lineRows-1 {
			for c, run := range line {
				if run == '.' {
					endCol = c - 1
					continue
				}
			}
		}

		for c, run := range line {
			dir := ""
			if run == '>' {
				dir = "E"
			} else if run == '<' {
				dir = "W"
			} else if run == 'v' {
				dir = "S"
			} else if run == '^' {
				dir = "N"
			} else if run == '#' || run == '.' {
				continue
			} else {
				panic("wrong run")
			}
			bo.blizzards[r-1][c-1] = dir
		}
	}

	start := newCoord(-1, startCol)
	end := newCoord(bo.nrRows, endCol)
	fmt.Println(start, end)

	found, best1 := rec(start, start, end, 1, bo, newCache(), newPath())
	fmt.Println(found, best1)

	found, best2 := rec(end, end, start, best1+1, bo, newCache(), newPath())
	fmt.Println(found, best2)

	found, best3 := rec(start, start, end, best1+best2+1, bo, newCache(), newPath())
	fmt.Println(found, best3)

	return best1 + best2 + best3
}

func rec(prePos, start, end coord, step int, bo *board, c *cache, p path) (found bool, bestFromHereToEnd int) {
	if prePos == end {
		return true, 0
	}

	exists, score := c.get(step, prePos)
	if exists {
		return true, score
	}

	// fmt.Printf("step: %d, prePos: %v, cache: %d\n", step, prePos, len(c.entries))

	if step == 1500 {
		return false, math.MaxInt / 2
	}

	nextCoords := findPossibleNextPositions(prePos, start, end, bo.nrRows, bo.nrCols)
	best := math.MaxInt / 2
	anyFound := false
	for _, next := range nextCoords {
		if !bo.isOpen(next, start, end, step) {
			continue
		}

		key := newKey(step, next)
		if p.contains(key) {
			continue
		}
		newPath := p.clone()
		newPath.history[key] = true

		found, score = rec(next, start, end, step+1, bo, c, newPath)
		if !found {
			continue
		}
		anyFound = true
		score++
		if score < best {
			best = score
		}
	}

	if anyFound {
		c.addEntry(newKey(step, prePos), best)
		return true, best
	} else {
		// fmt.Printf("not found: %v\n", newKey(step, prePos))
		c.addEntry(newKey(step, prePos), math.MaxInt/2)
		return false, math.MaxInt / 2
	}
}

func findPossibleNextPositions(pos, start, end coord, nrRows, nrCols int) (coords []coord) {
	newPos := newCoord(pos.r+1, pos.c) // S
	if newPos == end {
		return []coord{newPos}
	}
	if newPos.r >= 0 && newPos.r < nrRows && newPos.c >= 0 && newPos.c < nrCols {
		coords = append(coords, newPos)
	}

	newPos = newCoord(pos.r, pos.c+1) // E
	if newPos == end {
		return []coord{newPos}
	}
	if newPos.r >= 0 && newPos.r < nrRows && newPos.c >= 0 && newPos.c < nrCols {
		coords = append(coords, newPos)
	}

	newPos = newCoord(pos.r-1, pos.c) // N
	if newPos == end {
		return []coord{newPos}
	}
	if newPos.r >= 0 && newPos.r < nrRows && newPos.c >= 0 && newPos.c < nrCols {
		coords = append(coords, newPos)
	}

	newPos = newCoord(pos.r, pos.c-1) // W
	if newPos == end {
		return []coord{newPos}
	}
	if newPos.r >= 0 && newPos.r < nrRows && newPos.c >= 0 && newPos.c < nrCols {
		coords = append(coords, newPos)
	}

	coords = append(coords, pos) // wait

	return
}
