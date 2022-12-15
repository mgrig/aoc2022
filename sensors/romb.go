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

func (r romb) height(c coord) int {
	h := 1 + r.size - r.center.manhattanDist(c)
	if h < 0 {
		h = 0
	}
	return h
}
