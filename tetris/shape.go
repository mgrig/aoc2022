package tetris

type shape interface {
	getPos() coord
	getShapeCoords() []coord
	String() string
}

var _ shape = minus{}
var _ shape = plus{}
var _ shape = el{}
var _ shape = bar{}
var _ shape = square{}

func GetShape(index int, pos coord) shape {
	rem := index % 5
	switch rem {
	case 0:
		return newMinus(pos)
	case 1:
		return newPlus(pos)
	case 2:
		return newEl(pos)
	case 3:
		return newBar(pos)
	case 4:
		return newSquare(pos)
	default:
		panic("oops")
	}
}

func getAbsoluteShapeCoords(s shape) []coord {
	relCoords := s.getShapeCoords()
	coords := make([]coord, len(relCoords))
	for i, rC := range relCoords {
		coords[i] = s.getPos().plus(rC)
	}
	return coords
}

func moveRight(s shape, b *bottom, mapWidth int) (ok bool, newShape shape) {
	newShape = newPos(s, s.getPos().plus(newCoord(1, 0)))
	if isCollision(newShape, b, mapWidth) {
		return false, s
	}
	return true, newShape
}

func moveLeft(s shape, b *bottom, mapWidth int) (ok bool, newShape shape) {
	newShape = newPos(s, s.getPos().plus(newCoord(-1, 0)))
	if isCollision(newShape, b, mapWidth) {
		return false, s
	}
	return true, newShape
}

func moveDown(s shape, b *bottom) (ok bool, newShape shape) {
	newShape = newPos(s, s.getPos().plus(newCoord(0, -1)))
	if isCollision(newShape, b, 7) {
		return false, s
	}
	return true, newShape
}

func isCollision(s shape, b *bottom, mapWidth int) bool {
	absCoords := getAbsoluteShapeCoords(s)
	for _, aC := range absCoords {
		if aC.x < 0 || aC.x >= mapWidth || b.contains(aC) {
			return true
		}
	}
	return false
}

func newPos(s shape, pos coord) shape {
	_, ok := s.(minus)
	if ok {
		return newMinus(pos)
	}

	_, ok = s.(plus)
	if ok {
		return newPlus(pos)
	}

	_, ok = s.(el)
	if ok {
		return newEl(pos)
	}

	_, ok = s.(bar)
	if ok {
		return newBar(pos)
	}

	_, ok = s.(square)
	if ok {
		return newSquare(pos)
	}

	panic("wrong shape type")
}
