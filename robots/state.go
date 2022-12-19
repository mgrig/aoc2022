package robots

type state struct {
	availableRobots    map[string]int // ore, clay, obsidian, geode
	availableResources map[string]int
	decisionHistory    []string
}

func newState() state {
	return state{
		availableRobots:    make(map[string]int),
		availableResources: make(map[string]int),
		decisionHistory:    make([]string, 0),
	}
}

func (s state) clone() state {
	ret := newState()
	for _, rt := range resource_types {
		ret.availableRobots[rt] = s.availableRobots[rt]
		ret.availableResources[rt] = s.availableResources[rt]
	}
	ret.decisionHistory = make([]string, len(s.decisionHistory))
	copy(ret.decisionHistory, s.decisionHistory)
	return ret
}

// Simplify state to make them look more like each other, for caching.
// - if too many resources of a type (that cannot be spent by step maxSteps) then throw away excess
func (s *state) simplify(step int, blue *blueprint) {
	remainingSteps := maxSteps - step + 1 // includes "step" and maxSteps

	maxSpendableOre := remainingSteps * blue.maxCostByRobotType["ore"]
	if s.availableResources["ore"] > maxSpendableOre {
		s.availableResources["ore"] = maxSpendableOre
	}

	maxSpendableClay := remainingSteps * blue.maxCostByRobotType["clay"]
	if s.availableResources["clay"] > maxSpendableClay {
		s.availableResources["clay"] = maxSpendableClay
	}

	maxSpendableObsidian := remainingSteps * blue.maxCostByRobotType["obsidian"]
	if s.availableResources["obsidian"] > maxSpendableObsidian {
		s.availableResources["obsidian"] = maxSpendableObsidian
	}
}

// not enough steps remaining to beat the current best solution?
func (s state) canStillBeatBest(step int) bool {
	if step == maxSteps {
		return true
	}
	remainingSteps := maxSteps - step + 1
	maxGeodesTillEnd := s.availableResources["geode"] +
		remainingSteps*s.availableRobots["geode"] +
		remainingSteps*(remainingSteps+1)/2
	return maxGeodesTillEnd > bestScore
}
