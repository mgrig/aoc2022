package tetris

import "fmt"

type minus struct {
	pos coord
}

func newMinus(pos coord) shape {
	return minus{pos: pos}
}

func (m minus) getPos() coord {
	return m.pos
}

func (m minus) getShapeCoords() []coord {
	return []coord{
		newCoord(0, 0),
		newCoord(1, 0),
		newCoord(2, 0),
		newCoord(3, 0),
	}
}

func (m minus) String() string {
	return fmt.Sprintf("minus@%s", m.pos)
}
