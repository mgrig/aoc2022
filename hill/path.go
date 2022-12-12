package hill

type path struct {
	coords []coord
}

func newPath() path {
	return path{
		coords: make([]coord, 0),
	}
}

func clonePath(other path) path {
	return path{
		coords: other.coords,
	}
}

func (p path) withNewCoord(co coord) path {
	ret := clonePath(p)
	ret.coords = append(ret.coords, co)
	return ret
}

func (p path) nrSteps() int {
	return len(p.coords) - 1
}
