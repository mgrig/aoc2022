package cubes

type coord struct {
	x, y, z int
}

func newCoord(x, y, z int) coord {
	return coord{x: x, y: y, z: z}
}
