package rope

import (
	"aoc2022/common"
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func newCoord(x, y int) coord {
	return coord{
		x: x,
		y: y,
	}
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c *coord) move(directionDeg int) {
	switch directionDeg {
	case 0:
		c.x += 1
	case 90:
		c.y += 1
	case 180:
		c.x -= 1
	case 270:
		c.y -= 1
	case 45:
		c.x += 1
		c.y += 1
	case 135:
		c.x -= 1
		c.y += 1
	case 225:
		c.x -= 1
		c.y -= 1
	case 315:
		c.x += 1
		c.y -= 1
	default:
		panic(fmt.Sprintf("wrong direction %d", directionDeg))
	}
}

// ****

type headTail struct {
	head coord
	tail coord
}

func newHeadTail(head, tail coord) *headTail {
	return &headTail{
		head: head,
		tail: tail,
	}
}

func (ht *headTail) moveHead(directionDeg int) {
	ht.head.move(directionDeg)
	ht.pullTail()
}

func (ht *headTail) pullTail() {
	dx := ht.head.x - ht.tail.x
	dy := ht.head.y - ht.tail.y

	if common.IntAbs(dx) < 2 && common.IntAbs(dy) < 2 {
		return
	}

	if dx > 0 {
		if dy < 0 {
			ht.tail.move(315)
		} else if dy == 0 {
			ht.tail.move(0)
		} else { // dy > 0
			ht.tail.move(45)
		}
	} else if dx == 0 {
		if dy < 0 {
			ht.tail.move(270)
		} else if dy == 0 {
			// cannot be, could panic
		} else { // dy > 0
			ht.tail.move(90)
		}
	} else { // dx < 0
		if dy < 0 {
			ht.tail.move(225)
		} else if dy == 0 {
			ht.tail.move(180)
		} else { // dy > 0
			ht.tail.move(135)
		}
	}
}

func (ht headTail) String() string {
	return fmt.Sprintf("%s - %s", ht.head, ht.tail)
}

// ****

func CountTail(lines []string) int {
	tailPositions := make(map[coord]bool)
	ht := newHeadTail(newCoord(0, 0), newCoord(0, 0))
	tailPositions[ht.tail] = true

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic("wrong line format " + line)
		}

		directionDeg := 0
		switch tokens[0] {
		case "R":
			directionDeg = 0
		case "U":
			directionDeg = 90
		case "L":
			directionDeg = 180
		case "D":
			directionDeg = 270
		default:
			panic("wrong direction " + tokens[0])
		}

		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		for s := 0; s < steps; s++ {
			ht.moveHead(directionDeg)
			tailPositions[ht.tail] = true
		}
	}

	return len(tailPositions)
}

// ****

type rope struct {
	knots []coord
}

func newRope() *rope {
	ret := rope{
		knots: make([]coord, 10),
	}
	for i := 0; i < len(ret.knots); i++ {
		ret.knots[i] = newCoord(0, 0)
	}
	return &ret
}

func (rope *rope) moveHead(directionDeg int) {
	for i := 0; i < len(rope.knots)-1; i++ {
		ht := newHeadTail(rope.knots[i], rope.knots[i+1])
		if i == 0 {
			ht.head.move(directionDeg)
		}
		ht.pullTail()
		rope.knots[i] = ht.head
		tailMoved := rope.knots[i+1] != ht.tail
		rope.knots[i+1] = ht.tail
		if !tailMoved {
			break
		}
	}
}

func (r rope) getTail() coord {
	return r.knots[len(r.knots)-1]
}

func (rope rope) String() string {
	mat := make([][]string, 31)
	for r := 0; r < len(mat); r++ {
		mat[r] = make([]string, 31)
		for c := 0; c < len(mat[r]); c++ {
			mat[r][c] = "."
		}
	}

	mat[14][14] = "s"
	for i := len(rope.knots) - 1; i >= 0; i-- {
		knot := rope.knots[i]
		c := 15 + knot.x
		r := 15 - knot.y
		name := "H"
		if i > 0 {
			name = fmt.Sprintf("%d", i)
		}
		mat[r][c] = name
	}

	ret := ""
	for r := 0; r < len(mat); r++ {
		ret += fmt.Sprintf("%s\n", strings.Join(mat[r], ""))
	}
	return ret
}

// ****

func CountTail10(lines []string) int {
	tailPositions := make(map[coord]bool)
	rope := newRope()
	tailPositions[rope.getTail()] = true

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic("wrong line format " + line)
		}

		directionDeg := 0
		switch tokens[0] {
		case "R":
			directionDeg = 0
		case "U":
			directionDeg = 90
		case "L":
			directionDeg = 180
		case "D":
			directionDeg = 270
		default:
			panic("wrong direction " + tokens[0])
		}

		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		for s := 0; s < steps; s++ {
			rope.moveHead(directionDeg)
			tailPositions[rope.getTail()] = true
		}
		// fmt.Println(rope.String())
	}

	return len(tailPositions)
}
