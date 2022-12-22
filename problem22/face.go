package problem22

type face struct {
	id         int
	toAbsolute coordSystem
	size       int
	borders    map[int]transformation // 0 is right border on map, 1 - bottom border, etc TO ABSOLUTE
}

func newFace(id, size int, translationToAbsolute coord, borders map[int]transformation) face {
	return face{
		id:         id,
		size:       size,
		toAbsolute: newCoordSystem(translationToAbsolute, 0),
		borders:    borders,
	}
}

func (f face) contains(pos coord) bool {
	relPos := f.toAbsolute.fromOriginToThis(pos)
	return relPos.r >= 0 && relPos.r < f.size && relPos.c >= 0 && relPos.c < f.size
}
