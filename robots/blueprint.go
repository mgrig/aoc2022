package robots

import (
	"aoc2022/common"
	"regexp"
)

type blueprint struct {
	id                                  int
	ore_oreCost                         int
	clay_oreCost                        int
	obsidian_oreCost, obsidian_clayCost int
	geode_oreCost, geode_obsidianCost   int

	maxCostByRobotType map[string]int
}

var re *regexp.Regexp = regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

func parseBlueprint(line string) blueprint {
	tokens := re.FindStringSubmatch(line)
	if len(tokens) != 8 {
		panic("wrong line " + line)
	}

	blue := blueprint{
		id:                 common.StringToInt(tokens[1]),
		ore_oreCost:        common.StringToInt(tokens[2]),
		clay_oreCost:       common.StringToInt(tokens[3]),
		obsidian_oreCost:   common.StringToInt(tokens[4]),
		obsidian_clayCost:  common.StringToInt(tokens[5]),
		geode_oreCost:      common.StringToInt(tokens[6]),
		geode_obsidianCost: common.StringToInt(tokens[7]),
	}

	maxCosts := make(map[string]int)
	maxCosts["obsidian"] = blue.geode_obsidianCost
	maxCosts["clay"] = blue.obsidian_clayCost
	maxCosts["ore"] = common.IntMax(blue.ore_oreCost, blue.clay_oreCost, blue.obsidian_oreCost, blue.geode_oreCost)

	blue.maxCostByRobotType = maxCosts

	return blue
}
