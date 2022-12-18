package cubes

import (
	"aoc2022/common"
	"strings"
)

var faceDeltas map[int]coord = make(map[int]coord, 7)

func UnconnectedSides(lines []string) int {
	faceDeltas[1] = newCoord(0, -1, 0)
	faceDeltas[2] = newCoord(0, 0, 1)
	faceDeltas[3] = newCoord(-1, 0, 0)
	faceDeltas[4] = newCoord(1, 0, 0)
	faceDeltas[5] = newCoord(0, 0, -1)
	faceDeltas[6] = newCoord(0, 1, 0)

	s := newShape()
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		if len(tokens) != 3 {
			panic("wrong line " + line)
		}
		x := common.StringToInt(tokens[0])
		y := common.StringToInt(tokens[1])
		z := common.StringToInt(tokens[2])
		cb := newCube(newCoord(x, y, z))
		s.addCoord(&cb)
	}

	for c, cb := range s.coords {
		for face := 1; face <= 6; face++ {
			if cb.connectedFaces[face] {
				continue
			}
			cNeigh := newCoord(c.x+faceDeltas[face].x, c.y+faceDeltas[face].y, c.z+faceDeltas[face].z)
			if s.contains(cNeigh) {
				cb.markConnected(face)
				s.coords[cNeigh].markConnected(7 - face)
			}
		}
	}

	sum := 0
	for _, cb := range s.coords {
		sum += 6 - len(cb.connectedFaces)
	}

	return sum
}
