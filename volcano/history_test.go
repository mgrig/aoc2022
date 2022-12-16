package volcano

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistoryLastVisited(t *testing.T) {
	history := newHistory()
	history = history.addAction(newMoveAction("a", "b"))
	history = history.addAction(newMoveAction("b", "a"))

	visited, historyUntil := history.lastVisitTo("a")

	require.True(t, visited)
	fmt.Println(historyUntil)
}
