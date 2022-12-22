package problem22

import (
	"aoc2022/common"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveWithWrap(t *testing.T) {
	lines := common.GetLinesFromFile("../resources/22_board_test.txt", true, false)
	board := parseBoard(lines[:len(lines)-1])
	require.Equal(t, 12, board.nrRows())
	require.Equal(t, 16, board.nrCols())

	// right
	newPos := board.moveWithWrap(newPosition(0, 8, 0))
	require.Equal(t, newPosition(0, 9, 0), newPos)

	newPos = board.moveWithWrap(newPosition(0, 11, 0)) // should wrap
	require.Equal(t, newPosition(0, 8, 0), newPos)

	newPos = board.moveWithWrap(newPosition(11, 15, 0)) // off board, should wrap
	require.Equal(t, newPosition(11, 8, 0), newPos)

	// left
	newPos = board.moveWithWrap(newPosition(0, 8, 2)) // wrap
	require.Equal(t, newPosition(0, 11, 2), newPos)

	newPos = board.moveWithWrap(newPosition(0, 11, 2))
	require.Equal(t, newPosition(0, 10, 2), newPos)

	newPos = board.moveWithWrap(newPosition(4, 0, 2))
	require.Equal(t, newPosition(4, 11, 2), newPos)

	// down
	newPos = board.moveWithWrap(newPosition(0, 8, 1))
	require.Equal(t, newPosition(1, 8, 1), newPos)

	newPos = board.moveWithWrap(newPosition(11, 8, 1))
	require.Equal(t, newPosition(0, 8, 1), newPos)

	newPos = board.moveWithWrap(newPosition(7, 0, 1))
	require.Equal(t, newPosition(4, 0, 1), newPos)

	// up
	newPos = board.moveWithWrap(newPosition(0, 8, 3))
	require.Equal(t, newPosition(11, 8, 3), newPos)

	newPos = board.moveWithWrap(newPosition(1, 8, 3))
	require.Equal(t, newPosition(0, 8, 3), newPos)

	newPos = board.moveWithWrap(newPosition(4, 0, 3))
	require.Equal(t, newPosition(7, 0, 3), newPos)
}
