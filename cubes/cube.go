package cubes

import "fmt"

type cube struct {
	c            coord
	checkedFaces map[int]bool
}

func newCube(c coord) cube {
	return cube{
		c:            c,
		checkedFaces: make(map[int]bool, 7),
	}
}

func (cb *cube) markChecked(faceIndex int) {
	if faceIndex < 1 || faceIndex > 6 {
		panic(fmt.Sprintf("wrong face index %d", faceIndex))
	}
	cb.checkedFaces[faceIndex] = true
}
