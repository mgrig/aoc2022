package cubes

type coord struct {
	x, y, z int
}

func newCoord(x, y, z int) coord {
	return coord{x: x, y: y, z: z}
}

func (c coord) plus(other coord) coord {
	return newCoord(c.x+other.x, c.y+other.y, c.z+other.z)
}
