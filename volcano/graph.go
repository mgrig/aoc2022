package volcano

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type graph struct {
	nodes map[string]*node
}

func (g *graph) getOrCreate(name string) *node {
	pNode, exists := g.nodes[name]
	if !exists {
		pNode = &node{
			name: name,
			// flow added on the returned node!
			edgesTo: make([]*node, 0),
		}
		g.nodes[name] = pNode
	}
	return pNode
}

func (g graph) getExisting(name string) *node {
	pNode, exists := g.nodes[name]
	if !exists {
		panic("missing expected node for name " + name)
	}
	return pNode
}

func (g graph) getTotalNonZeroValves() int {
	count := 0
	for _, v := range g.nodes {
		if v.flow > 0 {
			count += 1
		}
	}
	return count
}

var re *regexp.Regexp = regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func parseGraph(lines []string) *graph {
	ret := graph{
		nodes: make(map[string]*node, 0),
	}

	for _, line := range lines {
		tokens := re.FindStringSubmatch(line)
		if len(tokens) != 4 {
			panic(fmt.Sprintf("wrong nr tokens (expected 4, was %d) on line: %s", len(tokens), line))
		}

		name := tokens[1]
		flow, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(err)
		}
		targets := strings.Split(tokens[3], ", ")

		// fmt.Println(name, flow, targets)
		node := ret.getOrCreate(name)
		node.flow = flow
		for _, target := range targets {
			node.edgesTo = append(node.edgesTo, ret.getOrCreate(target))
		}
	}

	return &ret
}
