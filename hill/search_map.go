package hill

type cellInfo struct {
	bestPathToHere path
}

func (ci cellInfo) bestSteps() int {
	return ci.bestPathToHere.nrSteps()
}

// ***

type searchMap struct {
	mapInfo [][]*cellInfo
}

func newSearchMap(nrR, nrC int) *searchMap {
	ret := searchMap{}
	ret.mapInfo = make([][]*cellInfo, nrR)
	for r := 0; r < nrR; r++ {
		ret.mapInfo[r] = make([]*cellInfo, nrC)
	}

	return &ret
}

func (sm *searchMap) newPath(p path) (wasBetter bool) {
	if len(p.coords) == 0 {
		return
	}

	co := p.coords[len(p.coords)-1]
	if sm.mapInfo[co.r][co.c] == nil || p.nrSteps() < sm.mapInfo[co.r][co.c].bestSteps() {
		sm.mapInfo[co.r][co.c] = &cellInfo{
			bestPathToHere: p,
		}
		return true
	}
	return false
}

func (sm *searchMap) newPathDown(p path) (wasBetter bool) {
	if len(p.coords) == 0 {
		return
	}

	co := p.coords[len(p.coords)-1]
	if sm.mapInfo[co.r][co.c] == nil || p.nrSteps() < sm.mapInfo[co.r][co.c].bestSteps() {
		sm.mapInfo[co.r][co.c] = &cellInfo{
			bestPathToHere: p,
		}
		return true
	}
	return false
}
