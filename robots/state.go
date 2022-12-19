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
// - if too many resources of a type (that cannot be spent by step 24) then throw away excess
func (s *state) simplify(step int, blue *blueprint) {
	remainingSteps := 24 - step + 1 // includes "step" and 24

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
