package problem24

import "fmt"

type coord struct {
	r, c int
}

func newCoord(r, c int) coord {
	return coord{
		r: r,
		c: c,
	}
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.r, c.c)
}
