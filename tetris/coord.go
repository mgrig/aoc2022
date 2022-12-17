package tetris

import "fmt"

type coord struct {
	x, y int
}

func newCoord(x, y int) coord {
	return coord{
		x: x,
		y: y,
	}
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coord) plus(other coord) coord {
	return newCoord(c.x+other.x, c.y+other.y)
}
