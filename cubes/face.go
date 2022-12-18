package cubes

import "fmt"

/*
Face index encoding is similar to a dice, and aligns with xyz axes as follows:
1 - towards you, decreasing y
2 - upwards, increasing z
3 - to the left, decreasing x
4 - to the right
5 - downwards
6 - to the back, away from you

Encoding allows easy computation of the opposite face, by computing: 7 - faceIndex
*/

// faceIndex > coord moving outwards from given face
var outDelta = map[int]coord{
	1: newCoord(0, -1, 0),
	2: newCoord(0, 0, 1),
	3: newCoord(-1, 0, 0),
	4: newCoord(1, 0, 0),
	5: newCoord(0, 0, -1),
	6: newCoord(0, 1, 0),
}

type face struct {
	x, y, z   int // expanded for easier comparison (not using coord here)
	faceIndex int
}

func newFace(c coord, faceIndex int) face {
	return face{
		x:         c.x,
		y:         c.y,
		z:         c.z,
		faceIndex: faceIndex,
	}
}

func getDirections(faceIndex int) []int {
	ret := make([]int, 0)
	for i := 1; i <= 6; i++ {
		if i != faceIndex && (7-i) != faceIndex {
			ret = append(ret, i)
		}
	}
	return ret
}

func oppositeFaceIndex(faceIndex int) int {
	return 7 - faceIndex
}

func getEdgeCoord(faceIndex, dir int) coord {
	return outDelta[faceIndex].plus(outDelta[dir])
}

func getConnectedCoord(dir int) coord {
	return outDelta[dir]
}

func (f face) String() string {
	return fmt.Sprintf("%d@(%d, %d, %d)", f.faceIndex, f.x, f.y, f.z)
}
