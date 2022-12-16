package volcano

import "fmt"

// var bestHistory *history
var bestScore int = 0

func Part1(lines []string) int {
	g := parseGraph(lines)

	// x := newHistory()
	// bestHistory = &x

	hist := newHistory()
	rec("AA", hist, g, 1, 30)

	// bestFlown := bestHistory.getTotalFlown(g, true)
	// fmt.Printf("best flown: %d, history: %s\n", bestFlown, bestHistory)
	fmt.Println("best score:", bestScore)

	return -1
}

func rec(name string, preHist history, g *graph, step, maxSteps int) {
	if step > maxSteps {
		return
	}

	// check if visited the same node before, with the same nr of valves open
	visited, historyUntil := preHist.lastVisitTo(name)
	if visited {
		if preHist.countOpenValves() == historyUntil.countOpenValves() {
			return
		} else {
			// fmt.Println("preHist  ", preHist)
			// fmt.Println("histUntil", historyUntil)
		}
	}

	node := g.getExisting(name)

	if step == maxSteps {
		currentFlown := preHist.getTotalFlown(g, true)
		// bestFlown := bestHistory.getTotalFlown(g, true)
		// if bestScore != bestFlown {
		// 	fmt.Printf("step=%d, currentFlown=%d, bestFlown_pre=%d, bestScore=%d, bestFlown=%d, hist: %s\n", step, currentFlown, bestFlown, bestScore, bestHistory.getTotalFlown(g, true), bestHistory)
		// 	panic("WTF")
		// }
		if currentFlown > bestScore {
			bestScore = currentFlown
			// bestHistory.actions = preHist.actions
			fmt.Printf("step=%d, currentFlown=%d, hist: %s\n", step, currentFlown, preHist)
		}
		return
	}

	if node.flow == 0 || preHist.isValveOpen(name) {
		// no point trying to open valve = 0 or valve already opened > just move along
	} else {
		newAction := newOpenAction(name)
		postHist := preHist.addAction(newAction)
		rec(name, postHist, g, step+1, maxSteps)
	}

	nextNodes := g.nodes[name].edgesTo
	for _, nextNode := range nextNodes {
		newAction := newMoveAction(name, nextNode.name)
		postHist := preHist.addAction(newAction)
		rec(nextNode.name, postHist, g, step+1, maxSteps)
	}
}
