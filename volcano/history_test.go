package volcano

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistoryLastVisited(t *testing.T) {
	history := newHistory()
	history = history.addAction(newMoveAction("a", "b"))
	history = history.addAction(newMoveAction("b", "c"))
	history = history.addAction(newMoveAction("c", "d"))
	history = history.addAction(newMoveAction("d", "b"))

	visited, historyUntil := history.lastVisitTo("b")

	require.True(t, visited)
	fmt.Println(historyUntil)
}
