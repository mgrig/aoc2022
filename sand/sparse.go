package sand

import (
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func newCoord(x, y int) coord {
	return coord{x: x, y: y}
}

// x,y
func parseCoord(str string) coord {
	tokens := strings.Split(str, ",")
	if len(tokens) != 2 {
		panic("wrong coord format " + str)
	}

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	return newCoord(x, y)
}

// ***

type matrix interface {
	add(...coord)
	contains(coord) bool
	maxY() int
	addHorizLine(y int)
}

var _ matrix = &sparseMatrix{}

// ***

type sparseMatrix struct {
	coords    map[coord]bool // to have a faster .contains()
	horizLine *int
}

func newSparseMatrix() *sparseMatrix {
	return &sparseMatrix{
		coords:    make(map[coord]bool),
		horizLine: nil,
	}
}

func (sm *sparseMatrix) add(coords ...coord) {
	for i := range coords {
		sm.coords[coords[i]] = true
	}
}

func (sm *sparseMatrix) addHorizLine(y int) {
	sm.horizLine = &y
}

func (sm sparseMatrix) contains(c coord) bool {
	if sm.horizLine != nil && c.y >= *sm.horizLine {
		return true
	}
	return sm.coords[c]
}

func (sm sparseMatrix) maxY() int {
	max := math.MinInt
	for k, _ := range sm.coords {
		if k.y > max {
			max = k.y
		}
	}
	return max
}
