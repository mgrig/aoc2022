package ringbuffer

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

type node struct {
	value       int
	left, right *node
}

func newNode(value int) *node {
	pNode := &node{
		value: value,
	}
	pNode.left = pNode
	pNode.right = pNode
	return pNode
}

func (n *node) addRight(other *node) {
	wasRight := n.right

	other.right = wasRight
	other.left = n

	n.right = other
	wasRight.left = other
}

func (n *node) addLeft(other *node) {
	wasLeft := n.left

	other.right = n
	other.left = wasLeft

	n.left = other
	wasLeft.right = other
}

func (n *node) extract() {
	l := n.left
	r := n.right

	l.right = r
	r.left = l
}

func (first *node) getByIndex(index int) *node {
	if index == 0 {
		return first
	}
	pNode := first
	for i := 0; i < common.IntAbs(index); i++ {
		pNode = pNode.right
	}
	return pNode
}

// steps > 0 means move right
// steps < 0 means move left
func (n *node) move(steps int) {
	if steps == 0 {
		return
	}
	if steps > 0 {
		n.moveRight(steps)
	} else {
		n.moveLeft(-steps)
	}
}

func (n *node) moveRight(steps int) {
	if steps < 0 {
		panic("steps must be >= 0")
	}
	pNode := n.getByIndex(steps)
	n.extract()
	pNode.addRight(n)
}

func (n *node) moveLeft(steps int) {
	if steps < 0 {
		panic("steps must be >= 0")
	}
	pNode := n
	for i := 0; i < steps; i++ {
		pNode = pNode.left
	}
	n.extract()
	pNode.addLeft(n)
}

func (n *node) String() string {
	return fmt.Sprintf("%p: %p < %d > %p", n, n.left, n.value, n.right)
}

func (n *node) Show() string {
	strs := make([]string, 0)

	pNode := n
	for {
		strs = append(strs, fmt.Sprint(pNode.value))
		pNode = pNode.right
		if pNode == n {
			break
		}
	}

	return strings.Join(strs, ", ")
}
