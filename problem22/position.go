package problem22

type position struct {
	r, c   int // 0-based
	facing int // 0 - right, 1 - down, 2 - left, 3 - up
}

func newPosition(r, c, facing int) position {
	return position{
		r:      r,
		c:      c,
		facing: facing,
	}
}

func (p position) turnRight() position {
	newFacing := (p.facing + 1) % 4
	return newPosition(p.r, p.c, newFacing)
}

func (p position) turnLeft() position {
	newFacing := (p.facing + 4 - 1) % 4
	return newPosition(p.r, p.c, newFacing)
}

func (p position) moveAhead(steps int, b *board) position {
	for step := 0; step < steps; step++ {
		nextPos := b.moveWithWrap(p)
		if b.matrix[nextPos.r][nextPos.c] == 2 { // wall
			return p
		}
		p = nextPos
	}
	return p
}
