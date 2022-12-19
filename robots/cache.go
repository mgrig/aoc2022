package robots

import (
	"math"
)

type cacheKey struct {
	step              int
	robotsOre         int
	robotsClay        int
	robotsObsidian    int
	robotsGeode       int
	resourcesOre      int
	resourcesClay     int
	resourcesObsidian int
	resourcesGeode    int
}

func newCacheKey(step int, s state) cacheKey {
	return cacheKey{
		step:              step,
		robotsOre:         s.availableRobots["ore"],
		robotsClay:        s.availableRobots["clay"],
		robotsObsidian:    s.availableRobots["obsidian"],
		robotsGeode:       s.availableRobots["geode"],
		resourcesOre:      s.availableResources["ore"],
		resourcesClay:     s.availableResources["clay"],
		resourcesObsidian: s.availableResources["obsidian"],
		resourcesGeode:    s.availableResources["geode"],
	}
}

// ***

type cache struct {
	entries    map[cacheKey]int // cache key > score
	minStep    int
	hitsByStep map[int]int
}

func newCache() cache {
	return cache{
		entries:    make(map[cacheKey]int),
		minStep:    math.MaxInt,
		hitsByStep: make(map[int]int),
	}
}

func (c *cache) addEntry(step int, s state, score int) {
	if step > 20 {
		return
	}
	key := newCacheKey(step, s)
	best, exists := c.entries[key]
	if !exists {
		c.entries[key] = score
		if step < c.minStep {
			c.minStep = step
			// fmt.Println("cache minStep", step)
			// fmt.Println("hits by step", c.hitsByStep)
		}
	} else if score > best {
		c.entries[key] = score
	}
}

func (c *cache) get(step int, s state) (found bool, score int) {
	score, found = c.entries[newCacheKey(step, s)]
	if found {
		c.hitsByStep[step]++
	}
	return
}
