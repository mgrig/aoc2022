package problem22

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

func (co coord) String() string {
	return fmt.Sprintf("(%d, %d)", co.r, co.c)
}
