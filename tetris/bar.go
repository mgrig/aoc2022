package tetris

import "fmt"

type bar struct {
	pos coord
}

func newBar(pos coord) shape {
	return bar{pos: pos}
}

func (b bar) getPos() coord {
	return b.pos
}

func (b bar) getShapeCoords() []coord {
	return []coord{
		newCoord(0, 0),
		newCoord(0, 1),
		newCoord(0, 2),
		newCoord(0, 3),
	}
}

func (b bar) String() string {
	return fmt.Sprintf("bar@%s", b.pos)
}
