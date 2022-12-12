package hill

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

func (co coord) neighbors(nrR, nrC int) (coords []coord) {
	coords = make([]coord, 0)

	if co.c > 0 {
		coords = append(coords, newCoord(co.r, co.c-1))
	}

	if co.c < nrC-1 {
		coords = append(coords, newCoord(co.r, co.c+1))
	}

	if co.r > 0 {
		coords = append(coords, newCoord(co.r-1, co.c))
	}

	if co.r < nrR-1 {
		coords = append(coords, newCoord(co.r+1, co.c))
	}

	return
}
