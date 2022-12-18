package volcano

import (
	"aoc2022/common"
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	a1 := newMoveAction("AA", "BB")
	a2 := newOpenAction("BB")

	h := newHistory()
	printHistoryPointers(&h)

	h1 := h.addAction(a1)
	printHistoryPointers(&h1)

	h2 := h1.addAction(a2)
	printHistoryPointers(&h2)
}

func printHistoryPointers(h *history) {
	if len(h.actions) > 0 {
		fmt.Printf("h: %p,\th.actions: %p,\th.actions[0] %p\n", h, &(h.actions), &(h.actions[0]))
	} else {
		fmt.Printf("h: %p,\th.actions: %p\n", h, &(h.actions))
	}
}

func TestManualHistory(t *testing.T) {
	lines := common.GetLinesFromFile("../resources/16_volcano_test.txt", true, true)
	g := parseGraph(lines)
	g.compress()

	actions := []action{
		newMoveAction("AA", "DD"),
		newOpenAction("DD"),
		newMoveAction("DD", "CC"),
		newMoveAction("CC", "BB"),
		newOpenAction("BB"),
		newMoveAction("BB", "AA"),
		newMoveAction("AA", "JJ"),
		newOpenAction("JJ"),
		newMoveAction("JJ", "AA"),
		newMoveAction("AA", "DD"),
		newMoveAction("DD", "EE"),
		newMoveAction("EE", "HH"),
		newOpenAction("HH"),
		newMoveAction("HH", "EE"),
		newOpenAction("EE"),
		newMoveAction("EE", "DD"),
		newMoveAction("DD", "CC"),
		newOpenAction("CC"),
		newNop(),
		newNop(),
		newNop(),
		newNop(),
		newNop(),
		newNop(),

		newNop(),
		newNop(),
	}

	h := newHistory()
	for _, a := range actions {
		h = h.addAction(a)
	}
	fmt.Println(h.getTotalFlown(g, 30))
}
