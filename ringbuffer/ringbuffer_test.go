package ringbuffer

import (
	"aoc2022/common"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func parseLine(line string) *node {
	tokens := strings.Split(line, ", ")

	var first, pNode *node
	for i, token := range tokens {
		value := common.StringToInt(token)
		pNode = newNode(value)
		if i == 0 {
			first = pNode
		} else {
			first.addLeft(pNode)
		}
	}
	return first
}

func TestMoveRight(t *testing.T) {
	ring := parseLine("4, 5, 6, 1, 7, 8, 9")
	expected := parseLine("4, 5, 6, 7, 1, 8, 9")
	ring.getByIndex(3).move(1)
	fmt.Println(ring.Show())

	require.Equal(t, expected, ring)
}

func TestMoveLeft(t *testing.T) {
	ring := parseLine("4, -2, 5, 6, 7, 8, 9")
	expected := parseLine("4, 5, 6, 7, 8, -2, 9")
	ring.getByIndex(1).move(-2)
	fmt.Println(ring.Show())

	require.Equal(t, expected, ring)
}
