package problem23

import "math"

type sparseMatrix struct {
	coords map[coord]bool
}

func newSparseMatrix() *sparseMatrix {
	return &sparseMatrix{
		coords: make(map[coord]bool),
	}
}

func (sm *sparseMatrix) addCoord(co coord) {
	sm.coords[co] = true
}

func (sm *sparseMatrix) delCoord(co coord) {
	delete(sm.coords, co)
}

func (sm *sparseMatrix) contains(co coord) bool {
	return sm.coords[co]
}

func (sm *sparseMatrix) getBox() (topLeft, bottomRight coord) {
	minr := math.MaxInt
	maxr := math.MinInt
	minc := minr
	maxc := maxr

	for k, _ := range sm.coords {
		if k.r < minr {
			minr = k.r
		}
		if k.r > maxr {
			maxr = k.r
		}
		if k.c < minc {
			minc = k.c
		}
		if k.c > maxc {
			maxc = k.c
		}
	}

	return newCoord(minr, minc), newCoord(maxr, maxc)
}

func (sm sparseMatrix) String() string {
	ret := ""

	topLeft, bottomRight := sm.getBox()
	for r := topLeft.r; r <= bottomRight.r; r++ {
		for c := topLeft.c; c <= bottomRight.c; c++ {
			if sm.coords[newCoord(r, c)] {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}

	return ret
}
