package tetris

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShapeMove(t *testing.T) {
	s := newMinus(newCoord(3, 3))
	b := newBottom(7)
	ok, _ := moveRight(s, &b, 7)
	require.False(t, ok)
}
