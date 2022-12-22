package problem22

type transformation struct {
	from, to coordSystem
}

func newTransformation(fromX, fromY, fromRot, toX, toY, toRot int) transformation {
	return transformation{
		from: newCoordSystem(newCoord(fromX, fromY), fromRot),
		to:   newCoordSystem(newCoord(toX, toY), toRot),
	}
}

func (t transformation) transform(posAbs position) position {
	relPos := t.from.fromOriginToThis(newCoord(posAbs.r, posAbs.c)) // same rel pos in both from and to
	newPosInAbs := t.to.fromThisToOrigin(relPos)
	newFacing := (posAbs.facing - (t.to.rotation-t.from.rotation)/90 + 4) % 4
	return newPosition(newPosInAbs.r, newPosInAbs.c, newFacing)
}

func (t transformation) reverse() transformation {
	return transformation{
		from: t.to,
		to:   t.from,
	}
}
