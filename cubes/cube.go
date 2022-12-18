package cubes

import "fmt"

type cube struct {
	c              coord
	connectedFaces map[int]bool
}

func newCube(c coord) cube {
	return cube{
		c:              c,
		connectedFaces: make(map[int]bool, 7),
	}
}

func (cb *cube) markConnected(faceIndex int) {
	if faceIndex < 1 || faceIndex > 6 {
		panic(fmt.Sprintf("wrong face index %s", faceIndex))
	}
	cb.connectedFaces[faceIndex] = true
}
