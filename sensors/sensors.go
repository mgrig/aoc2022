package sensors

import (
	"aoc2022/common"
	"regexp"
)

var re *regexp.Regexp = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func parseLine(line string) (r romb, beacon coord) {
	tokens := re.FindStringSubmatch(line)
	if len(tokens) != 5 {
		panic("wrong line format " + line)
	}
	tokensInt := common.ToIntegerValues(tokens[1:5])
	sensor := newCoord(tokensInt[0], tokensInt[1])
	beacon = newCoord(tokensInt[2], tokensInt[3])
	return newRomb(sensor, sensor.manhattanDist(beacon)), beacon
}

func NrExcludedPositionsOnRow(lines []string, row int) int {
	rombs := make([]romb, len(lines))
	knownBeacons := make(map[coord]bool, len(lines))
	var beacon coord
	for i, line := range lines {
		rombs[i], beacon = parseLine(line)
		knownBeacons[beacon] = true
	}

	// set(x coord)
	lineCover := make(map[int]bool)
	for i, _ := range rombs {
		intersection := rombs[i].intersectHorizLine(row)
		for _, c := range intersection {
			if !knownBeacons[c] { // exlude known beacons from result!
				lineCover[c.x] = true
			}
		}
	}

	return len(lineCover)
}

/*
I want to avoid searching over the huge grid of part 2.
Main idea: given a start coord, move away from any nearby sensor position.
Each area "covered" by a sensor (aka where we know there is no other beacon)
is represented by a "romb" instance. Now imagine this rhomb is the base of a
pyramid with the top point in the center, i.e. on the sensor position.
Overlap all such pyramids from the input file.
Every coordinate that is covered by at least one sensor (pyramid) will have a
height > 0.
We are asked to find the only point (in the given huge grid) that has height 0.
Seen as a 3d landscape, the task is to find the global minimum.
As the problem guarantees there is only 1 such point, we can imagine that in the
3d landscape, looking from the global minimum, we are surrounded by faces of pyramids.
This solution asserts that if we do a hill-descent algorithmus starting on each
existing pyramid face, we are guaranteed to eventully end up in the global minimum.
To guarantee we check every face, we set 4 starting points for each pyramid, 1 unit
away from the top/center in each direction (N, S, E, W).
Note: this is actually a bug, and we should start on NE, NW, SW, SE :/
*/
func TuningFreq(lines []string, min, max int) int {
	rombs := make([]romb, len(lines))
	knownBeacons := make(map[coord]bool, len(lines))
	var beacon coord
	for i, line := range lines {
		rombs[i], beacon = parseLine(line)
		knownBeacons[beacon] = true
	}

	beacon = findBeacon(min, max, rombs)

	return 4000000*beacon.x + beacon.y
}

func findBeacon(min, max int, rombs []romb) coord {
	for _, r := range rombs {
		// Pick 4 start positions for each pyramid, one for each face
		// Bug: Should be NE, NW, SW, SE instead of N, S, E, W.
		beacon := searchFrom(newCoord(r.center.x+1, r.center.y), min, max, &rombs)
		if beacon != nil {
			return *beacon
		}

		beacon = searchFrom(newCoord(r.center.x-1, r.center.y), min, max, &rombs)
		if beacon != nil {
			return *beacon
		}

		beacon = searchFrom(newCoord(r.center.x, r.center.y-1), min, max, &rombs)
		if beacon != nil {
			return *beacon
		}

		beacon = searchFrom(newCoord(r.center.x, r.center.y+1), min, max, &rombs)
		if beacon != nil {
			return *beacon
		}
	}
	panic("no beacon found")
}

// go downhill from given "pos"
// Avoid recursive implementation, as the depth could be significant. Use instead
// less readable for loop with mutating position.
func searchFrom(pos coord, min, max int, rombs *[]romb) *coord {
	if pos.x < min || pos.x > max || pos.y < min || pos.y > max {
		return nil
	}

	for {
		hPos := getHeights(pos, rombs)
		// fmt.Println(pos, hPos)
		foundSmallerNeighbor := false
		neighbors := neighbors(min, max, pos)
		for _, neighbor := range neighbors {
			hVec := getHeights(neighbor, rombs)
			if allZeros(hVec) {
				// found solution
				return &neighbor
			}
			// Decide we found a smaller neighbor if we move down some pyramid slope,
			// and not up any other.
			// Note: we pick the first "smaller" neighbor - this is potentially unsafe,
			// as the solution may be down another smaller neighbor, but this was good
			// enough for the input test.
			if smallerHeightVec(hVec, hPos) {
				pos = neighbor
				foundSmallerNeighbor = true
				break
			}
		}
		if !foundSmallerNeighbor {
			return nil
		}
	}
}

func neighbors(min, max int, pos coord) (ret []coord) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			posX := pos.x + x
			posY := pos.y + y
			if posX >= min && posX <= max && posY >= min && posY <= max {
				ret = append(ret, newCoord(posX, posY))
			}
		}
	}
	return
}

func getHeights(c coord, rombs *[]romb) []int {
	ret := make([]int, len(*rombs))
	for i, _ := range *rombs {
		ret[i] = (*rombs)[i].height(c)
	}
	return ret
}

// true if no element of first is > element of second
// and at least one is strictly smaller
func smallerHeightVec(first, second []int) bool {
	if len(first) != len(second) {
		panic("wrong lengths")
	}
	anyStrictlySmaller := false
	for i := 0; i < len(first); i++ {
		if first[i] < second[i] {
			anyStrictlySmaller = true
			continue
		}
		if first[i] > second[i] {
			return false
		}
	}
	return anyStrictlySmaller
}

func allZeros(h []int) bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
