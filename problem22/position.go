package problem22

import "fmt"

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

func (p position) String() string {
	return fmt.Sprintf("(%d, %d)f%d", p.r, p.c, p.facing)
}

func (p position) getCoord() coord {
	return newCoord(p.r, p.c)
}

func (p position) turnRight() position {
	newFacing := (p.facing + 1) % 4
	return newPosition(p.r, p.c, newFacing)
}

func (p position) turnLeft() position {
	newFacing := (p.facing + 4 - 1) % 4
	return newPosition(p.r, p.c, newFacing)
}

func (p position) moveAhead(steps int, b *board, faces map[int]face) position {
	for step := 0; step < steps; step++ {
		fmt.Println("p:", p)
		nextPos := b.moveWithWrap(p, faces)
		fmt.Println("nextPos:", nextPos)
		if b.matrix[nextPos.r][nextPos.c] == 2 { // wall
			return p
		}
		p = nextPos
	}
	return p
}
