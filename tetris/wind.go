package tetris

type wind struct {
	line  string
	index int
}

func newWind(line string) wind {
	return wind{
		line:  line,
		index: 0,
	}
}

// returns "<" or ">"
func (w *wind) getNext() string {
	ret := string(w.line[w.index%len(w.line)])
	w.index += 1
	return ret
}

func (w wind) getUpcomingIndex() int {
	return w.index
}
