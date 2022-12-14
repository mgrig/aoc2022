package sand

import (
	"aoc2022/common"
	"strings"
)

func parseCoods(line string) (coords []coord) {
	tokens := strings.Split(line, " -> ")
	if len(tokens) < 2 {
		panic("wrong line format " + line)
	}

	for i := 0; i < len(tokens)-1; i++ {
		from := parseCoord(tokens[i])
		to := parseCoord(tokens[i+1])
		if from.y == to.y {
			minx := common.IntMin(from.x, to.x)
			maxx := common.IntMax(from.x, to.x)
			for ix := minx; ix <= maxx; ix++ {
				coords = append(coords, newCoord(ix, from.y))
			}
		} else if from.x == to.x {
			miny := common.IntMin(from.y, to.y)
			maxy := common.IntMax(from.y, to.y)
			for iy := miny; iy <= maxy; iy++ {
				coords = append(coords, newCoord(from.x, iy))
			}
		} else {
			panic("diagonal line? " + line)
		}
	}
	return
}

func Part1(lines []string) int {
	var rocks matrix = newSparseMatrix()

	for _, line := range lines {
		rocks.add(parseCoods(line)...)
	}

	maxY := rocks.maxY()
	start := newCoord(500, 0)

	sandAtRest := 0
	for { // sand units
		endedAtRest := sandFall(start, rocks, start, maxY+3)
		if endedAtRest {
			sandAtRest++
		} else {
			break
		}
	}

	return sandAtRest
}

func sandFall(sandPos coord, rocks matrix, start coord, maxY int) (endedAtRest bool) {
	for { // falling
		if sandPos.y > maxY {
			// falling through
			return false
		}

		down := newCoord(sandPos.x, sandPos.y+1)
		if !rocks.contains(down) {
			return sandFall(down, rocks, start, maxY)
		}

		downLeft := newCoord(sandPos.x-1, sandPos.y+1)
		if !rocks.contains(downLeft) {
			return sandFall(downLeft, rocks, start, maxY)
		}

		downRight := newCoord(sandPos.x+1, sandPos.y+1)
		if !rocks.contains(downRight) {
			return sandFall(downRight, rocks, start, maxY)
		}

		rocks.add(sandPos)
		if sandPos == start {
			// pretend it falls through to force end
			return false
		}
		return true
	}
}

func Part2(lines []string) int {
	var rocks matrix = newSparseMatrix()

	for _, line := range lines {
		rocks.add(parseCoods(line)...)
	}

	bottom := rocks.maxY() + 2
	rocks.addHorizLine(bottom)
	start := newCoord(500, 0)

	sandAtRest := 0
	for { // sand units
		endedAtRest := sandFall(start, rocks, start, bottom+1)
		if endedAtRest {
			sandAtRest++
		} else {
			break
		}
	}

	return sandAtRest + 1
}
