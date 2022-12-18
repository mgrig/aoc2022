package volcano

import (
	"fmt"
	"strings"
)

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

func (h history) getLastAction() action {
	if len(h.actions) > 0 {
		return h.actions[len(h.actions)-1]
	}
	return nil
}

func (h history) sanityCheck() {
	if len(h.actions) > 0 {
		currentName := "AA"
		for _, a := range h.actions {
			if !a.canApplyTo(currentName) {
				panic("failed sanity check")
			}
			moveAction, ok := a.(move)
			if ok {
				currentName = moveAction.to
			}
		}
	}
}

func (h history) addAction(a action) history {
	if !a.canApplyTo(h.getCurrentName()) {
		panic(fmt.Sprintf("incompatible action %s on name %s", a, h.getCurrentName()))
	}
	newHist := history{
		actions: make([]action, len(h.actions)),
	}
	copy(newHist.actions, h.actions)
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

func (h history) getTotalFlown(g *graph, maxCompletedSteps int) (totalFlown, completedSteps int) {
	totalFlown = 0
	flow := 0
	steps := 0
	for _, a := range h.actions {
		if steps == maxCompletedSteps {
			return totalFlown, steps
		}
		if moveAction, ok := a.(move); ok { // track current name
			moveDuration := g.getEdge(moveAction.from, moveAction.to).weight
			if steps+moveDuration > maxCompletedSteps {
				totalFlown += flow * (maxCompletedSteps - steps)
				steps = maxCompletedSteps
				return totalFlown, steps
			} else {
				totalFlown += flow * moveDuration
				steps += moveDuration
			}
		} else if openAction, ok := a.(open); ok {
			totalFlown += flow // open is 1 step
			flow += g.getExistingNode(openAction.name).flow
			steps++
		} else {
			// must be nop
			totalFlown += flow
			steps++
		}
	}
	return totalFlown, steps
}

func (h history) lastVisitTo(name string) (visited bool, historyUntil history) {
	if len(h.actions) <= 1 {
		return false, newHistory()
	}

	lastIndex := -1
	leftNameValve := false
	leftFirstValve := false
	for i := 0; i < len(h.actions)-1; i++ {
		if h.actions[i].isOpen(name) {
			lastIndex = i
			continue
		}
		if h.actions[i].isMoveTo(name) {
			lastIndex = i
			continue
		}
		if i == 0 && h.actions[i].isMoveFrom(name) {
			leftFirstValve = true
			lastIndex = 0
		}
		if lastIndex != -1 {
			leftNameValve = true
		}
	}

	if lastIndex != -1 && leftNameValve {
		return true, history{actions: h.actions[0 : lastIndex+1]}
	}

	if leftFirstValve {
		return true, newHistory()
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

func (h history) getCurrentName() string {
	ret := "AA"
	for _, a := range h.actions {
		moveAction, ok := a.(move)
		if ok {
			ret = moveAction.to
		}
	}
	return ret
}
