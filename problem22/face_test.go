package problem22

import (
	"aoc2022/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformations(t *testing.T) {
	lines := common.GetLinesFromFile("../resources/22_board_test.txt", true, false)
	board := parseBoard(lines[:len(lines)-1])
	require.Equal(t, 12, board.nrRows())
	require.Equal(t, 16, board.nrCols())

	N := 4
	// define faces
	f1 := newFace(1, N, newCoord(0, 8), nil)
	f2 := newFace(1, N, newCoord(4, 0), nil)
	f3 := newFace(1, N, newCoord(4, 4), nil)
	f4 := newFace(1, N, newCoord(4, 8), map[int]transformation{0: newTransformation(7, 12, 90, 8, 12, 0)})
	f5 := newFace(1, N, newCoord(8, 8), map[int]transformation{1: newTransformation(12, 8, 0, 7, 3, 180)})
	f6 := newFace(1, N, newCoord(8, 12), nil)
	fmt.Println(f1, f2, f3, f4, f5, f6)

	// A -> B example
	A := newCoord(5, 12)
	B := f4.borders[0].transform(A)
	require.Equal(t, newCoord(8, 14), B)

	// C -> D example
	C := newCoord(12, 10)
	D := f5.borders[1].transform(C)
	require.Equal(t, newCoord(7, 1), D)
}
