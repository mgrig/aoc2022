package problem24

type path struct {
	history map[key]bool
}

func newPath() path {
	return path{
		history: make(map[key]bool),
	}
}

func (p path) clone() path {
	newPath := newPath()
	for k, v := range p.history {
		newPath.history[k] = v
	}
	return newPath
}

func (p path) contains(k key) bool {
	_, exists := p.history[k]
	return exists
}
