package tetris

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStateEquals(t *testing.T) {
	state1 := newState(1, 2, newBottom(7).skyline())
	state2 := newState(1, 2, newBottom(7).skyline())

	require.True(t, reflect.DeepEqual(state1, state2))
}
