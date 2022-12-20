package ringbuffer

import (
	"aoc2022/common"
	"fmt"
)

func Part1(lines []string) int {
	N := len(lines)
	originalOrder := make([]*node, N)
	var first, pNode *node
	for i, line := range lines {
		value := common.StringToInt(line)
		pNode = newNode(value)
		if i == 0 {
			first = pNode
		} else {
			first.addLeft(pNode)
		}
		originalOrder[i] = pNode
	}

	for _, v := range originalOrder {
		v.move(v.value % (N - 1))
	}

	// fmt.Println(first.Show())

	pNode = first
	index := 0
	for {
		if pNode.value == 0 {
			break
		}
		index++
		pNode = pNode.right
	}

	n1 := first.getByIndex((1000 + index) % N).value
	n2 := first.getByIndex((2000 + index) % N).value
	n3 := first.getByIndex((3000 + index) % N).value

	fmt.Println(n1, n2, n3)

	return n1 + n2 + n3
}

func Part2(lines []string, multiplier int) int {
	N := len(lines)
	originalOrder := make([]*node, N)
	var first, pNode *node
	for i, line := range lines {
		value := common.StringToInt(line) * multiplier
		pNode = newNode(value)
		if i == 0 {
			first = pNode
		} else {
			first.addLeft(pNode)
		}
		originalOrder[i] = pNode
	}

	for i := 0; i < 10; i++ {
		for _, v := range originalOrder {
			v.move(v.value % (N - 1))
		}
	}

	// fmt.Println(first.Show())

	pNode = first
	index := 0
	for {
		if pNode.value == 0 {
			break
		}
		index++
		pNode = pNode.right
	}

	n1 := first.getByIndex((1000 + index) % N).value
	n2 := first.getByIndex((2000 + index) % N).value
	n3 := first.getByIndex((3000 + index) % N).value

	fmt.Println(n1, n2, n3)

	return n1 + n2 + n3
}
