package cubes

type shape struct {
	coords map[coord]*cube
}

func newShape() shape {
	return shape{
		coords: make(map[coord]*cube),
	}
}

func (s *shape) addCoord(cb *cube) {
	s.coords[cb.c] = cb
}

func (s shape) contains(c coord) bool {
	_, exists := s.coords[c]
	return exists
}
