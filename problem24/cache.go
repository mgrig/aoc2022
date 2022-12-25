package problem24

type key struct {
	step int
	r, c int
}

func newKey(step int, pos coord) key {
	return key{
		step: step % lcm,
		r:    pos.r,
		c:    pos.c,
	}
}

// ***

type cache struct {
	entries map[key]int // key > bestFromKeyToEnd
}

func newCache() *cache {
	return &cache{
		entries: make(map[key]int),
	}
}

func (c *cache) addEntry(k key, score int) {
	best, exists := c.entries[k]
	if !exists {
		// fmt.Printf("add new: key %v, value %d\n", k, score)
		c.entries[k] = score
	} else if score < best {
		// fmt.Printf("update: key %v, value %d\n", k, score)
		c.entries[k] = score
	}
}

func (c *cache) get(step int, pos coord) (exists bool, score int) {
	score, exists = c.entries[newKey(step, pos)]
	return
}
