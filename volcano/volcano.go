package volcano

import (
	"fmt"
)

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

func Part2(lines []string) int {
	g := parseGraph(lines)

	// for _, n := range g.nodes {
	// if n.flow > 0 {
	// 	fmt.Printf("  %s [label = \"%s:%d\", color = \"green\"]\n", n.name, n.name, n.flow)
	// } else {
	// 	fmt.Printf("  %s [label = \"%s:%d\"]\n", n.name, n.name, n.flow)
	// }
	// for _, next := range n.edgesTo {
	// 	if n.name < next.name {
	// 		fmt.Printf("  %s -- %s\n", n.name, next.name)
	// 	}
	// }
	// }

	hist := newHistory()
	histElefant := newHistory()
	rec2("AA", "AA", hist, histElefant, g, 1, 26)

	fmt.Println("best score:", bestScore)

	return -1
}

func rec2(name, nameElefant string, preHist history, preHistElefant history, g *graph, step, maxSteps int) {
	if len(preHist.actions) != len(preHistElefant.actions) {
		panic("wrong lengths")
	}
	preHist.sanityCheck()
	preHistElefant.sanityCheck()

	// fmt.Printf("step:%d, name:%s, nameE: %s\n", step, name, nameElefant)
	if step > maxSteps {
		return
	}

	remainingValvesToOpen := g.getTotalNonZeroValves() - preHist.countOpenValves() - preHistElefant.countOpenValves()
	// if remainingValvesToOpen == 0 {
	// 	// solution?
	// 	currentFlown := preHist.getTotalFlown(g, true) + preHistElefant.getTotalFlown(g, true)
	// 	if currentFlown > bestScore {
	// 		bestScore = currentFlown
	// 		// fmt.Printf("step=%d, currentFlown=%d,\n\thist:  %s\n\thistE: %s\n", step, currentFlown, preHist, preHistElefant)
	// 	}
	// 	return
	// }

	// check if visited the same node before, with the same nr of valves open
	a := preHist.getLastAction()
	if a != nil && !a.isNop() {
		visited, historyUntil := preHist.lastVisitTo(name)
		if visited {
			if preHist.countOpenValves() == historyUntil.countOpenValves() {
				return
			}
		}
	}

	aElefant := preHistElefant.getLastAction()
	if aElefant != nil && !aElefant.isNop() {
		visitedElefant, historyUntilElefant := preHistElefant.lastVisitTo(nameElefant)
		if visitedElefant {
			if preHistElefant.countOpenValves() == historyUntilElefant.countOpenValves() {
				return
			}
		}
	}

	// fmt.Printf("step:%d, name:%s, nameE: %s\n", step, name, nameElefant)
	node := g.getExisting(name)
	nodeElefant := g.getExisting(nameElefant)

	if step == maxSteps {
		currentFlown := preHist.getTotalFlown(g, true) + preHistElefant.getTotalFlown(g, true)
		if currentFlown > bestScore {
			bestScore = currentFlown
			fmt.Printf("step=%d, currentFlown=%d,\n\thist:  %s\n\thistE: %s\n", step, currentFlown, preHist, preHistElefant)
		}
		return
	}

	possibleActions := make([]recArgs, 0)
	if remainingValvesToOpen > 0 {
		if node.flow == 0 || preHist.isValveOpen(name) || preHistElefant.isValveOpen(name) {
			// no point trying to open valve = 0 or valve already opened > just move along
		} else {
			newAction := newOpenAction(name)
			postHist := preHist.addAction(newAction)
			possibleActions = append(possibleActions, recArgs{a: newAction, name: name, postHist: postHist})
			// rec2(name, postHist, g, step+1, maxSteps)
		}

		nextNodes := g.nodes[name].edgesTo
		for _, nextNode := range nextNodes {
			newAction := newMoveAction(name, nextNode.name)
			postHist := preHist.addAction(newAction)
			possibleActions = append(possibleActions, recArgs{a: newAction, name: nextNode.name, postHist: postHist})
			// rec2(nextNode.name, postHist, g, step+1, maxSteps)
		}
	}

	if remainingValvesToOpen == 0 { // only elefant is allowed to pick the last valve
		newAction := newNop()
		postHist := preHist.addAction(newAction)
		possibleActions = append(possibleActions, recArgs{a: newAction, name: name, postHist: postHist})
	}

	possibleActionsElefant := make([]recArgs, 0)
	if remainingValvesToOpen > 0 {
		if nodeElefant.flow == 0 || preHist.isValveOpen(nameElefant) || preHistElefant.isValveOpen(nameElefant) {
			// no point trying to open valve = 0 or valve already opened > just move along
		} else {
			newActionElefant := newOpenAction(nameElefant)
			postHistElefant := preHistElefant.addAction(newActionElefant)
			possibleActionsElefant = append(possibleActionsElefant, recArgs{a: newActionElefant, name: nameElefant, postHist: postHistElefant})
		}

		nextNodesElefant := g.nodes[nameElefant].edgesTo
		for _, nextNodeElefant := range nextNodesElefant {
			newActionElefant := newMoveAction(nameElefant, nextNodeElefant.name)
			postHistElefant := preHistElefant.addAction(newActionElefant)
			possibleActionsElefant = append(possibleActionsElefant, recArgs{a: newActionElefant, name: nextNodeElefant.name, postHist: postHistElefant})
		}
	}

	if remainingValvesToOpen <= 1 {
		newActionElefant := newNop()
		postHistElefant := preHistElefant.addAction(newActionElefant)
		possibleActionsElefant = append(possibleActionsElefant, recArgs{a: newActionElefant, name: nameElefant, postHist: postHistElefant})
	}

	//// target hists
	//targetHist := []action{
	//	newMoveAction("AA", "II"),
	//	newMoveAction("II", "JJ"),
	//	newOpenAction("JJ"),
	//	newMoveAction("JJ", "II"),
	//	newMoveAction("II", "AA"),
	//	newMoveAction("AA", "BB"),
	//	newOpenAction("BB"),
	//	newMoveAction("BB", "CC"),
	//	newOpenAction("CC"),
	//	newNop(),
	//	newNop(),
	//	// newNop(),
	//}
	//
	//targetHistElefant := []action{
	//	newMoveAction("AA", "DD"),
	//	newOpenAction("DD"),
	//	newMoveAction("DD", "EE"),
	//	newMoveAction("EE", "FF"),
	//	newMoveAction("FF", "GG"),
	//	newMoveAction("GG", "HH"),
	//	newOpenAction("HH"),
	//	newMoveAction("HH", "GG"),
	//	newMoveAction("GG", "FF"),
	//	newMoveAction("FF", "EE"),
	//	newOpenAction("EE"),
	//	// newNop(),
	//}

	// combine actions
	//rand.Shuffle(len(possibleActions), func(i, j int) {
	//	possibleActions[i], possibleActions[j] = possibleActions[j], possibleActions[i]
	//})
	//rand.Shuffle(len(possibleActionsElefant), func(i, j int) {
	//	possibleActionsElefant[i], possibleActionsElefant[j] = possibleActionsElefant[j], possibleActionsElefant[i]
	//})
	for _, a := range possibleActions {
		for _, aElefant := range possibleActionsElefant {
			if a.a.isOpenAction() && a.a == aElefant.a {
				continue
			}
			if remainingValvesToOpen > 0 && a.a.isNop() && aElefant.a.isNop() {
				continue
			}

			//if equalActions(targetHist, a.postHist.actions) && equalActions(targetHistElefant, aElefant.postHist.actions) {
			//	// totalFlown := history{actions: targetHist}.getTotalFlown(g, false) + history{actions: targetHistElefant}.getTotalFlown(g, false)
			//	totalFlown := a.postHist.getTotalFlown(g, false) + aElefant.postHist.getTotalFlown(g, false)
			//	fmt.Printf("%d: %s\t%s > total: %d\n", step, a.a, aElefant.a, totalFlown)
			//	rec2(a.name, aElefant.name, a.postHist, aElefant.postHist, g, step+1, maxSteps)
			//} else {
			rec2(a.name, aElefant.name, a.postHist, aElefant.postHist, g, step+1, maxSteps)
			//}
		}
	}
	// fmt.Printf("  out of actions %s %s\n", name, nameElefant)
}

type recArgs struct {
	a        action
	name     string
	postHist history
}
