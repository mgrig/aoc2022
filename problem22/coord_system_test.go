package problem22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromOriginToThis(t *testing.T) {
	cs := newCoordSystem(newCoord(5, 7), 0)
	newPos := cs.fromOriginToThis(newCoord(8, 8))
	require.Equal(t, newCoord(3, 1), newPos)

	cs = newCoordSystem(newCoord(0, 0), 90)
	newPos = cs.fromOriginToThis(newCoord(2, 3))
	require.Equal(t, newCoord(3, -2), newPos)

	cs = newCoordSystem(newCoord(3, 1), 90)
	newPos = cs.fromOriginToThis(newCoord(2, 3))
	require.Equal(t, newCoord(2, 1), newPos)
}

func TestFromThisToOrigin(t *testing.T) {
	cs := newCoordSystem(newCoord(5, 7), 0)
	newPos := cs.fromThisToOrigin(newCoord(3, 1))
	require.Equal(t, newCoord(8, 8), newPos)

	cs = newCoordSystem(newCoord(0, 0), 90)
	newPos = cs.fromThisToOrigin(newCoord(3, -2))
	require.Equal(t, newCoord(2, 3), newPos)

	cs = newCoordSystem(newCoord(3, 1), 90)
	newPos = cs.fromThisToOrigin(newCoord(2, 1))
	require.Equal(t, newCoord(2, 3), newPos)
}
