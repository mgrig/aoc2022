package tetris

import "fmt"

type plus struct {
	pos coord
}

func newPlus(pos coord) shape {
	return plus{pos: pos}
}

func (p plus) getPos() coord {
	return p.pos
}

func (p plus) getShapeCoords() []coord {
	return []coord{
		newCoord(0, 1),
		newCoord(1, 1),
		newCoord(2, 1),
		newCoord(1, 0),
		newCoord(1, 2),
	}
}

func (p plus) String() string {
	return fmt.Sprintf("plus@%s", p.pos)
}
