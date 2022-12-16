package sensors

import (
	"aoc2022/common"
	"fmt"
)

type romb struct {
	center coord
	size   int
}

func newRomb(center coord, size int) romb {
	return romb{
		center: center,
		size:   size,
	}
}

func (r romb) String() string {
	return fmt.Sprintf("<%s, %d>", r.center, r.size)
}

// Returns all intersection points of a line at given y with the area covered by this rhombus.
func (r romb) intersectHorizLine(y int) (ret []coord) {
	vertDistToCenter := common.IntAbs(y - r.center.y)
	if vertDistToCenter > r.size {
		// no intersection, horiz line too far
		return
	}

	dx := r.size - vertDistToCenter

	for x := r.center.x - dx; x <= r.center.x+dx; x++ {
		ret = append(ret, newCoord(x, y))
	}

	return
}

func (r romb) covers(c coord) bool {
	return r.center.manhattanDist(c) <= r.size
}

// Assume a pyramid with this rhombus as base and max value on the center.
// This method computes the height at the given coord.
func (r romb) height(c coord) int {
	h := 1 + r.size - r.center.manhattanDist(c)
	if h < 0 {
		h = 0
	}
	return h
}

func (r romb) minDist(other romb) int {
	return r.center.manhattanDist(other.center) - r.size - other.size - 1
}

// Returns all coords at given manhattan dist from the center
func (r romb) isoline(dist int) []coord {
	coords := make([]coord, 4*dist)
	for i := 0; i < dist; i++ {
		coords[i] = newCoord(r.center.x+i, r.center.y-dist+i)
		coords[dist+i] = newCoord(r.center.x+dist-i, r.center.y+i)
		coords[2*dist+i] = newCoord(r.center.x-i, r.center.y+dist-i)
		coords[3*dist+i] = newCoord(r.center.x-dist+i, r.center.y-i)
	}
	return coords
}
