package problem23

type coord struct {
	r, c int
}

func newCoord(r, c int) coord {
	return coord{
		r: r,
		c: c,
	}
}

func (co coord) north() coord {
	return newCoord(co.r-1, co.c)
}

func (co coord) south() coord {
	return newCoord(co.r+1, co.c)
}

func (co coord) east() coord {
	return newCoord(co.r, co.c+1)
}

func (co coord) west() coord {
	return newCoord(co.r, co.c-1)
}
