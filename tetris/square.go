package tetris

import "fmt"

type square struct {
	pos coord
}

func newSquare(pos coord) shape {
	return square{pos: pos}
}

func (s square) getPos() coord {
	return s.pos
}

func (s square) getShapeCoords() []coord {
	return []coord{
		newCoord(0, 0),
		newCoord(1, 0),
		newCoord(0, 1),
		newCoord(1, 1),
	}
}

func (s square) String() string {
	return fmt.Sprintf("square@%s", s.pos)
}
