package tetris

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBottomClean(t *testing.T) {
	b := newBottom(7)
	b.addCoord(newCoord(0, 1))

	require.True(t, b.contains(newCoord(0, 1)))
	require.False(t, b.contains(newCoord(0, 0)))
}
