package cubes

import (
	"aoc2022/common"
	"fmt"
	"math"
	"strings"
)

func ExteriorSides(lines []string) int {
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

	// find external face
	// - find any cube with min y
	// - pick face index 1
	minY := math.MaxInt
	var minYCube *cube
	for c, cb := range s.coords {
		if c.y < minY {
			minY = c.y
			minYCube = cb
		}
	}
	if minYCube == nil {
		panic("no start cube found")
	}
	startFace := newFace(minYCube.c, 1)

	facesChecked := newFaceSet()
	facesToCheck := newFaceSet()
	facesToCheck.addFace(startFace)

	for len(facesToCheck.faces) > 0 {
		f := facesToCheck.getAny()
		fCoord := newCoord(f.x, f.y, f.z)
		// fmt.Printf("checking %v\n", f)

		dirs := getDirections(f.faceIndex)
		for _, dir := range dirs {
			if dir == 0 {
				panic(fmt.Sprintf("wrong dir %d", dir))
			}
			var toCheck face
			// check edge cube first

			if edgeCoord := fCoord.plus(getEdgeCoord(f.faceIndex, dir)); s.contains(edgeCoord) {
				toCheck = newFace(edgeCoord, oppositeFaceIndex(dir))
			} else if connectedCoord := fCoord.plus(getConnectedCoord(dir)); s.contains(connectedCoord) {
				toCheck = newFace(connectedCoord, f.faceIndex)
			} else {
				toCheck = newFace(fCoord, dir)
			}

			if !facesChecked.contains(toCheck) {
				facesToCheck.addFace(toCheck)
			}
		}
		facesToCheck.removeFace(f)
		facesChecked.addFace(f)

		// fmt.Printf("  toCheck: %s\n", facesToCheck)
		// fmt.Printf("  checked: %s\n", facesChecked)
	}

	return len(facesChecked.faces)
}
