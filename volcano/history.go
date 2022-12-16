package volcano

import "strings"

type history struct {
	actions []action
}

func newHistory() history {
	return history{
		actions: make([]action, 0),
	}
}

func (h history) String() string {
	strs := make([]string, len(h.actions))
	for i, a := range h.actions {
		strs[i] = a.String()
	}
	return strings.Join(strs, ", ")
}

func (h history) addAction(a action) history {
	newHist := h
	newHist.actions = append(newHist.actions, a)
	return newHist
}

func (h history) isValveOpen(name string) bool {
	for _, a := range h.actions {
		if a.isOpen(name) {
			return true
		}
	}
	return false
}

func (h history) getTotalFlown(g *graph, withLastStep bool) int {
	totalFlown := 0
	flow := 0
	for _, a := range h.actions {
		totalFlown += flow
		openAction, ok := a.(open)
		if ok {
			flow += g.getExisting(openAction.name).flow
		}
	}
	if withLastStep {
		totalFlown += flow
	}
	return totalFlown
}

func (h history) lastVisitTo(name string) (visited bool, historyUntil history) {
	if len(h.actions) <= 1 {
		return false, newHistory()
	}

	lastIndex := -1
	leftNameValve := false
	for i := 0; i < len(h.actions)-1; i++ {
		if h.actions[i].isOpen(name) {
			lastIndex = i
			continue
		}
		if h.actions[i].isMoveTo(name) {
			lastIndex = i
			continue
		}
		if lastIndex != -1 {
			leftNameValve = true
		}
	}

	if lastIndex != -1 && leftNameValve {
		return true, history{actions: h.actions[0 : lastIndex+1]}
	}

	return false, newHistory()
}

func (h history) countOpenValves() int {
	count := 0
	for _, a := range h.actions {
		if a.isOpenAction() {
			count += 1
		}
	}
	return count
}
