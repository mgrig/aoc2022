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
