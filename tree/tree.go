package tree

import "fmt"

func Part2(lines []string) int {
	root, humn, _ := parseNodes(lines)
	root.operation = "="

	fmt.Println("old graph:\n", root.graphvizString())

	// transform the original tree into a new tree which has "humn" as the new root!
	/*
		Example: say in the original graph we model 3 + 5 resulting to the parent value = 8
		    aaaa:"+"
		   /       \
		bbbb:3   humn:5

		The new graph will look like this:
		    humn:"-"
		   /       \
		aaaa:8   bbbb:3

		On the new graph, just call getValue() on the new root node, which is "humn".
	*/
	newKnown := newKnownNodes()
	newHumn := rec(humn, newKnown)

	fmt.Println("new graph:\n", newHumn.graphvizString())

	return newHumn.getValue()
}

func rec(oldNode *node, newKnown *knownNodes) (newNode *node) {
	// fmt.Printf("rec: oldNode=%v\n", oldNode)
	switch oldNode.parent.operation {
	case "=":
		value := getSibling(oldNode).getValue()
		newNode = createValueNode(oldNode.name, value, newKnown)
	case "-":
		if oldNode.parent.left == oldNode {
			// node was left of parent
			rightNewNode := createValueNode(oldNode.parent.right.name, oldNode.parent.right.getValue(), newKnown)
			newNode = createOperationNode(oldNode.name, "+", rightNewNode, rec(oldNode.parent, newKnown), newKnown)
		} else {
			// node was right of parent
			leftNewNode := createValueNode(oldNode.parent.left.name, oldNode.parent.left.getValue(), newKnown)
			newNode = createOperationNode(oldNode.name, "-", leftNewNode, rec(oldNode.parent, newKnown), newKnown)
		}
	case "+":
		oldSibling := getSibling(oldNode)
		newSiblingNode := createValueNode(oldSibling.name, oldSibling.getValue(), newKnown)
		newNode = createOperationNode(oldNode.name, "-", rec(oldNode.parent, newKnown), newSiblingNode, newKnown)
	case "*":
		oldSibling := getSibling(oldNode)
		newSiblingNode := createValueNode(oldSibling.name, oldSibling.getValue(), newKnown)
		newNode = createOperationNode(oldNode.name, "/", rec(oldNode.parent, newKnown), newSiblingNode, newKnown)
	case "/":
		if oldNode.parent.left == oldNode {
			// node was left of parent
			rightNewNode := createValueNode(oldNode.parent.right.name, oldNode.parent.right.getValue(), newKnown)
			newNode = createOperationNode(oldNode.name, "*", rightNewNode, rec(oldNode.parent, newKnown), newKnown)
		} else {
			// node was right of parent
			leftNewNode := createValueNode(oldNode.parent.left.name, oldNode.parent.left.getValue(), newKnown)
			newNode = createOperationNode(oldNode.name, "/", leftNewNode, rec(oldNode.parent, newKnown), newKnown)
		}
	default:
		panic("oops")
	}
	return
}

func getSibling(n *node) *node {
	if n.parent.left == n {
		return n.parent.right
	}
	return n.parent.left
}

func createOperationNode(name string, operation string, left, right *node, known *knownNodes) *node {
	pNode := known.getOrCreateEmpty(name)
	pNode.typ = "operation"
	pNode.directValue = -1
	pNode.operation = operation
	pNode.left = known.getOrCreateEmpty(left.name)
	pNode.left.parent = pNode
	pNode.right = known.getOrCreateEmpty(right.name)
	pNode.right.parent = pNode
	return pNode
}

func createValueNode(name string, value int, known *knownNodes) *node {
	pNode := known.getOrCreateEmpty(name)
	pNode.typ = "value"
	pNode.directValue = value
	pNode.operation = ""
	pNode.left = nil
	pNode.right = nil
	return pNode
}
