package tetris

import "fmt"

type el struct {
	pos coord
}

func newEl(pos coord) shape {
	return el{pos: pos}
}

func (l el) getPos() coord {
	return l.pos
}

func (l el) getShapeCoords() []coord {
	return []coord{
		newCoord(0, 0),
		newCoord(1, 0),
		newCoord(2, 0),
		newCoord(2, 1),
		newCoord(2, 2),
	}
}

func (l el) String() string {
	return fmt.Sprintf("el@%s", l.pos)
}
