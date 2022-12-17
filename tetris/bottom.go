package tetris

import "fmt"

type bottom struct {
	coords map[coord]bool
	high   int
}

func newBottom(width int) bottom {
	ret := make(map[coord]bool, width)
	for i := 0; i < width; i++ {
		ret[newCoord(i, 0)] = true
	}
	return bottom{
		coords: ret,
		high:   0,
	}
}

func (b *bottom) addCoord(c coord) {
	b.coords[c] = true
	if c.y > b.high {
		b.high = c.y
	}

	if len(b.coords) > 200 {
		b.clean()
	}
}

func (b bottom) highest() int {
	return b.high
	// ret := 0
	// for k := range b.coords {
	// 	if k.y > ret {
	// 		ret = k.y
	// 	}
	// }
	// return ret
}

func (b bottom) lowest() int {
	ret := b.high
	for k := range b.coords {
		if k.y < ret {
			ret = k.y
		}
	}
	return ret
}

func (b bottom) contains(c coord) bool {
	return b.coords[c]
}

func (b *bottom) clean() {
	toRemove := make([]coord, 0)
	for k := range b.coords {
		if k.y < b.high-40 {
			toRemove = append(toRemove, k)
		}
	}
	// fmt.Printf("remove %d from bottom\n", len(toRemove))
	for _, c := range toRemove {
		delete(b.coords, c)
	}
}

func (b bottom) plot(bottomFromZero bool) {
	miny := 0
	if !bottomFromZero {
		miny = b.lowest()
	}
	for y := b.high; y >= miny; y-- {
		for x := 0; x < 7; x++ {
			if b.contains(newCoord(x, y)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b bottom) skyline() skyline {
	ret := make([]coord, 7)
	for x := 0; x < 7; x++ {
		for y := b.high; y >= 0; y-- {
			cxy := newCoord(x, y)
			if b.contains(cxy) {
				ret[x] = cxy
				break
			}
		}
	}
	return skyline{
		coords: ret,
	}
}
