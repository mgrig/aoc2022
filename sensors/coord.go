package sensors

import (
	"aoc2022/common"
	"fmt"
)

type coord struct {
	x, y int
}

func newCoord(x, y int) coord {
	return coord{x: x, y: y}
}

func (c coord) manhattanDist(other coord) int {
	return common.IntAbs(c.x-other.x) + common.IntAbs(c.y-other.y)
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}
