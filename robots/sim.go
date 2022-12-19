package robots

import "fmt"

var bestScore int

func Part1(lines []string) int {
	initState := newState()
	initState.availableRobots["ore"] = 1

	totalScore := 0
	for _, line := range lines {
		bestScore = 0
		blue := parseBlueprint(line)
		c := newCache()
		geodes := rec(1, initState, &blue, &c)
		fmt.Println("blueprint ", blue.id, "geodes", geodes)
		totalScore += blue.id * geodes
	}

	return totalScore
}

func rec(step int, preState state, blue *blueprint, c *cache) int {
	// fmt.Println("start step", step)
	// fmt.Println(step, preState.decisionHistory)

	// any shortcuts?

	// - not enough steps remaining to beat the current best solution?

	// simplify state

	found, score := c.get(step, preState)
	if found {
		// fmt.Println("cache hit in step", step)
		return score
	}

	// list possible actions
	// - actions/decisions are only about if and what to build
	// - even if can build a cheap robot, may decide to wait to build a more expensive one
	// - if can build any robot, must build one
	canBuild := canBuildRobots(preState, blue)
	// fmt.Println("  canBuild:", canBuild)

	// each available robot collects its resource
	postCollectionState := preState.clone()
	for _, resType := range resource_types {
		postCollectionState.availableResources[resType] += postCollectionState.availableRobots[resType]
	}
	// fmt.Println("  post collection res:", postCollectionState.availableResources)

	if step == 24 {
		currentGeodes := postCollectionState.availableResources["geode"]
		if currentGeodes > bestScore {
			bestScore = currentGeodes
			fmt.Println("best score:", bestScore, "state", postCollectionState.decisionHistory)
		}
		return currentGeodes
	}

	// for each action
	// - apply action > postState
	// - check termination
	// - call rec() with postState
	bestFromThisState := 0
	for i := len(resource_types) - 1; i >= 0; i-- { // build higher grade robot first (if possible)
		resType := resource_types[i]
		if canBuild[resType] {
			// maybe we already have enough robots of this type?
			if resType != "geode" && preState.availableRobots[resType] == blue.maxCostByRobotType[resType] {
				continue
			}

			postState := postCollectionState.clone()
			postState.availableRobots[resType]++
			cost := costRobot(resType, blue)
			for k, v := range cost {
				postState.availableResources[k] -= v
			}

			// fmt.Println("  decision: build robot ", resType)
			// fmt.Println("  postState:", postState)
			postState.decisionHistory = append(postState.decisionHistory, fmt.Sprintf("(build %s)", resType))
			postState.simplify(step, blue)

			finalGeodes := rec(step+1, postState, blue, c)
			if finalGeodes > bestFromThisState {
				bestFromThisState = finalGeodes
			}
		}
	}

	if makesSenseToWait(preState, canBuild) {
		// fmt.Println("  decision: wait")
		// fmt.Println("  postState:", postCollectionState)
		postState := postCollectionState.clone()
		postState.decisionHistory = append(postState.decisionHistory, "(wait)")
		postState.simplify(step, blue)

		finalGeodes := rec(step+1, postState, blue, c)
		if finalGeodes > bestFromThisState {
			bestFromThisState = finalGeodes
		}
	}

	c.addEntry(step, preState, bestFromThisState)
	return bestFromThisState
}

func costRobot(resType string, blue *blueprint) map[string]int {
	switch resType {
	case "ore":
		return map[string]int{
			"ore": blue.ore_oreCost,
		}
	case "clay":
		return map[string]int{
			"ore": blue.clay_oreCost,
		}
	case "obsidian":
		return map[string]int{
			"ore":  blue.obsidian_oreCost,
			"clay": blue.obsidian_clayCost,
		}
	case "geode":
		return map[string]int{
			"ore":      blue.geode_oreCost,
			"obsidian": blue.geode_obsidianCost,
		}
	default:
		panic("wrong resType " + resType)
	}
}

func makesSenseToWait(s state, canBuild map[string]bool) bool {
	if s.availableRobots["clay"] == 0 &&
		s.availableRobots["obsidian"] == 0 &&
		s.availableRobots["geode"] == 0 {
		if canBuild["ore"] && canBuild["clay"] {
			return false
		}
		return true
	}

	// we produce ore and clay > may wait to build an obsidian robot
	if s.availableRobots["obsidian"] == 0 &&
		s.availableRobots["geode"] == 0 {
		return true
	}

	// we produce ore, clay and obsidian > may wait to build a geode robot
	if s.availableRobots["geode"] == 0 {
		return true
	}

	// we produce every type > may wait to get enough resources to build a particular robot
	// UNLESS we can already build any robot, in which case we must not wait
	if canBuildAnyRobot(canBuild) {
		return false
	} else {
		return true
	}
}

func canBuildAnyRobot(canBuild map[string]bool) bool {
	for _, resType := range resource_types {
		if !canBuild[resType] {
			return false
		}
	}
	return true
}

func canBuildRobots(preState state, blue *blueprint) map[string]bool {
	canBuild := make(map[string]bool)

	canBuild["ore"] = preState.availableResources["ore"] >= blue.ore_oreCost
	canBuild["clay"] = preState.availableResources["ore"] >= blue.clay_oreCost
	canBuild["obsidian"] = preState.availableResources["ore"] >= blue.obsidian_oreCost &&
		preState.availableResources["clay"] >= blue.obsidian_clayCost
	canBuild["geode"] = preState.availableResources["ore"] >= blue.geode_oreCost &&
		preState.availableResources["obsidian"] >= blue.geode_obsidianCost

	return canBuild
}
