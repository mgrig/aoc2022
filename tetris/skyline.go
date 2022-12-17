package tetris

import (
	"fmt"
	"math"
)

type skyline struct {
	coords []coord
}

func (s skyline) sanityCheck() {
	if len(s.coords) != 7 {
		panic("wrong skyline size")
	}
}

func (s skyline) isFlatBottomForMinus() bool {
	_, max := s.minmax()
	maxGapOf4 := 0
	for x := 0; x < 7; x++ {
		if s.coords[x].y < max {
			// there's a gap at max
			maxGapOf4 += 1
			if maxGapOf4 == 4 {
				return false
			}
		} else {
			maxGapOf4 = 0 // reset gap
		}
	}
	return true
}

func (s skyline) isFlatBottomPostMinus() bool {
	if !s.isFlatBottomForMinus() {
		return false
	}
	_, max := s.minmax()
	return s.coords[0].y == max && s.coords[1].y == max && s.coords[6].y == max
}

func (s skyline) minmax() (min, max int) {
	max = 0
	min = math.MaxInt
	for _, c := range s.coords {
		if c.y > max {
			max = c.y
		}
		if c.y < min {
			min = c.y
		}
	}
	return min, max
}

func (s skyline) plot() {
	s.sanityCheck()
	min, max := s.minmax()

	for y := max; y >= min; y-- {
		for x := 0; x < 7; x++ {
			if s.coords[x].y == y {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (s skyline) normalizedEquals(other skyline) bool {
	diff := s.coords[0].y - other.coords[0].y
	for x := 1; x < 7; x++ {
		if diff != s.coords[x].y-other.coords[x].y {
			return false
		}
	}
	return true
}
