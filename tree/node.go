package tree

import (
	"aoc2022/common"
	"fmt"
	"regexp"
)

type node struct {
	name        string
	cachedValue *int
	typ         string
	directValue int
	operation   string
	left, right *node
	parent      *node
}

var reValue *regexp.Regexp = regexp.MustCompile(`(.*): (-?\d+)`)
var reOperation *regexp.Regexp = regexp.MustCompile(`(.*): (.*) ([\+\-\*/]) (.*)`)

func (n *node) String() string {
	return fmt.Sprintf("(%p: %s, op %v)", n, n.name, n.name)
}

func (n *node) getValue() int {
	if n.cachedValue != nil {
		return *n.cachedValue
	}

	computedValue := n.computeValueWithoutCache()
	n.cachedValue = &computedValue

	return computedValue
}

func (n node) computeValueWithoutCache() int {
	switch n.typ {
	case "value":
		return n.directValue
	case "operation":
		leftValue := n.left.getValue()
		rightValue := n.right.getValue()
		switch n.operation {
		case "+":
			return leftValue + rightValue
		case "-":
			return leftValue - rightValue
		case "*":
			return leftValue * rightValue
		case "/":
			return leftValue / rightValue
		default:
			panic("unknown operation " + n.operation)
		}
	default:
		panic("unknown type " + n.typ)
	}
}

// ***
type knownNodes map[string]*node

func newKnownNodes() *knownNodes {
	var kn knownNodes = make(map[string]*node)
	return &kn
}

func (kn *knownNodes) getOrCreateEmpty(name string) *node {
	pNode, ok := (*kn)[name]
	if !ok {
		pNode = &node{
			name: name,
		}
		(*kn)[name] = pNode
	}
	return pNode
}

// ***

func parseNodes(lines []string) (root *node, humn *node, known *knownNodes) {
	known = newKnownNodes()
	for _, line := range lines {
		var pNode *node
		tokens := reOperation.FindStringSubmatch(line)
		if len(tokens) == 5 {
			// parse as operation node
			pNode = known.getOrCreateEmpty(tokens[1])
			// name already set
			pNode.cachedValue = nil
			pNode.typ = "operation"
			pNode.directValue = -1
			pNode.operation = tokens[3]
			pNode.left = known.getOrCreateEmpty(tokens[2])
			pNode.left.parent = pNode
			pNode.right = known.getOrCreateEmpty(tokens[4])
			pNode.right.parent = pNode
		} else {
			tokens = reValue.FindStringSubmatch(line)
			if len(tokens) != 3 {
				panic("wrong line " + line)
			}
			// parse as value node
			pNode = known.getOrCreateEmpty(tokens[1])
			// name already set
			pNode.cachedValue = nil
			pNode.typ = "value"
			pNode.directValue = common.StringToInt(tokens[2])
			pNode.operation = ""
			pNode.left = nil
			pNode.right = nil
		}
		if pNode.name == "root" {
			root = pNode
		}
		if pNode.name == "humn" {
			humn = pNode
		}
	}
	return
}

func (root *node) graphvizString() string {
	ret := "graph G {\n"

	ret += recGraphviz(root)

	ret += "}"
	return ret
}

func recGraphviz(n *node) (ret string) {
	style := ""
	if n.name == "root" {
		style = `, color = "blue"`
	}
	if n.name == "humn" {
		style = `, color = "red"`
	}

	if n.typ == "value" {
		ret = fmt.Sprintf("  %s [label = %d%s]\n", n.name, n.directValue, style)
	} else {
		ret = fmt.Sprintf("  %s [label = \"%s\"%s]\n", n.name, n.operation, style)
		ret += fmt.Sprintf("    %s -- %s\n", n.name, n.left.name)
		ret += fmt.Sprintf("    %s -- %s\n", n.name, n.right.name)
		ret += fmt.Sprintf(`    {rank = same; %s -- %s [style=invis]}`, n.left.name, n.right.name)
		ret += fmt.Sprintln()
		ret += fmt.Sprintln()

		ret += recGraphviz(n.left)
		ret += recGraphviz(n.right)
	}
	return
}
