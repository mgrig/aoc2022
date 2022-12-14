package hill

import "math"

func parseMap(lines []string) (karte [][]int, start, end coord) { //TODO add start, end
	nrR := len(lines)
	karte = make([][]int, nrR)
	for r, line := range lines {
		karte[r] = make([]int, len(line))
		for c, value := range line {
			if value == 'S' {
				karte[r][c] = 0
				start = newCoord(r, c)
			} else if value == 'E' {
				karte[r][c] = int('z') - int('a')
				end = newCoord(r, c)
			} else {
				karte[r][c] = int(value) - int('a')
			}
		}
	}
	return
}

func NrSteps(lines []string) int {
	karte, start, end := parseMap(lines)
	nrR := len(karte)
	nrC := len(karte[0])

	path := newPath()
	searchMap := newSearchMap(nrR, nrC)
	moveTo(start, path, searchMap, karte, end)

	return searchMap.mapInfo[end.r][end.c].bestSteps()
}

func moveTo(cell coord, previousPath path, searchMap *searchMap, karte [][]int, end coord) {
	path := previousPath.withNewCoord(cell)
	wasBetter := searchMap.newPath(path)
	if !wasBetter {
		return
	}

	if cell == end {
		return
	}

	nrR := len(searchMap.mapInfo)
	nrC := len(searchMap.mapInfo[0])
	neighbors := cell.neighbors(nrR, nrC)

	for _, neighbor := range neighbors {
		// reject neighbor if path too steep
		if karte[neighbor.r][neighbor.c]-karte[cell.r][cell.c] > 1 {
			continue
		}

		moveTo(neighbor, path, searchMap, karte, end)
	}
}

func NrStepsDown(lines []string) int {
	karte, _, end := parseMap(lines)
	nrR := len(karte)
	nrC := len(karte[0])

	path := newPath()
	searchMap := newSearchMap(nrR, nrC)
	shortest := math.MaxInt
	moveToDown(end, path, searchMap, karte, &shortest)

	return shortest
}

func moveToDown(cell coord, previousPath path, searchMap *searchMap, karte [][]int, shortest *int) {
	path := previousPath.withNewCoord(cell)
	wasBetter := searchMap.newPathDown(path)
	if !wasBetter {
		return
	}

	if karte[cell.r][cell.c] == 0 {
		if path.nrSteps() < *shortest {
			*shortest = path.nrSteps()
		}
		return
	}

	nrR := len(searchMap.mapInfo)
	nrC := len(searchMap.mapInfo[0])
	neighbors := cell.neighbors(nrR, nrC)

	for _, neighbor := range neighbors {
		// reject neighbor if drop too steep
		if karte[cell.r][cell.c]-karte[neighbor.r][neighbor.c] > 1 {
			continue
		}

		moveToDown(neighbor, path, searchMap, karte, shortest)
	}
}
