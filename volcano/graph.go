package volcano

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type graph struct {
	nodes map[string]*node
	edges map[string]map[string]*edge
}

func (g *graph) getOrCreateNode(name string) *node {
	pNode, exists := g.nodes[name]
	if !exists {
		pNode = &node{
			name: name,
			// flow added on the returned node!
		}
		g.nodes[name] = pNode
	}
	return pNode
}

func (g graph) getExistingNode(name string) *node {
	pNode, exists := g.nodes[name]
	if !exists {
		panic("missing expected node for name " + name)
	}
	return pNode
}

func (g *graph) createEdge(from, to string, weight int) {
	if from == to {
		panic("self-edge not allowed")
	}

	if from > to { // save with from < to
		from, to = to, from
	}

	pEdge := &edge{
		from:   g.getOrCreateNode(from),
		to:     g.getOrCreateNode(to),
		weight: weight,
	}

	_, exists := g.edges[from]
	if !exists {
		g.edges[from] = make(map[string]*edge)
	}
	g.edges[from][to] = pEdge
}

// result can be nil
func (g graph) getEdge(from, to string) *edge {
	if from == to {
		panic("self-edge not allowed")
	}

	if from > to { // save with from < to
		from, to = to, from
	}

	to2Edge, exists := g.edges[from]
	if !exists {
		return nil
	}
	return to2Edge[to]
}

func (g *graph) replaceOrCreateEdge(from, to string, weight int) {
	existingEdge := g.getEdge(from, to)
	if existingEdge != nil {
		if existingEdge.weight > weight {
			existingEdge.weight = weight
		}
		return
	}

	g.createEdge(from, to, weight)
}

func (g *graph) removeEdge(from, to string) {
	if from == to {
		panic("self-edge not allowed")
	}

	if from > to { // save with from < to
		from, to = to, from
	}

	delete(g.edges[from], to)
}

var re *regexp.Regexp = regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func parseGraph(lines []string) *graph {
	g := graph{
		nodes: make(map[string]*node, 0),
		edges: make(map[string]map[string]*edge, 0),
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
		node := g.getOrCreateNode(name)
		node.flow = flow

		for _, target := range targets {
			g.createEdge(node.name, target, 1)
		}
	}

	return &g
}

func (g graph) String() string {
	strs := make([]string, 0)
	for _, pNode := range g.nodes {
		strs = append(strs, pNode.String())
	}
	ret := fmt.Sprintf("Nodes: %s\n", strings.Join(strs, " "))

	strs = make([]string, 0)
	for _, to2Edge := range g.edges {
		for _, pEdge := range to2Edge {
			strs = append(strs, pEdge.String())
		}
	}
	ret += fmt.Sprintf("Edges: %s\n", strings.Join(strs, " "))

	return ret
}

func (g graph) GraphvizString() string {
	ret := "graph G {\n"

	for _, pNode := range g.nodes {
		colorPart := ""
		if pNode.flow > 0 {
			colorPart = `, color = "red"`
		}
		ret += fmt.Sprintf("  %s [label = \"%s:%d\"%s]\n", pNode.name, pNode.name, pNode.flow, colorPart)
	}

	ret += fmt.Sprintln()
	for _, to2Edge := range g.edges {
		for _, pEdge := range to2Edge {
			ret += fmt.Sprintf("  %s -- %s [label = %d]\n", pEdge.from.name, pEdge.to.name, pEdge.weight)
		}
	}

	ret += "}\n"

	return ret
}

func (g graph) getNeighbors(name string) []*node {
	ret := make([]*node, 0)

	for _, to2Edge := range g.edges {
		for _, pEdge := range to2Edge {
			if pEdge.from.name == name {
				ret = append(ret, pEdge.to)
			} else if pEdge.to.name == name {
				ret = append(ret, pEdge.from)
			}
		}
	}

	return ret
}

func (g *graph) findNodeWithNoFlowAndFewestNeighbors() *node {
	minNeighbors := math.MaxInt
	var minNode *node
	for _, pNode := range g.nodes {
		if pNode.flow > 0 || pNode.name == "AA" {
			continue
		}
		nrNeighs := len(g.getNeighbors(pNode.name))
		if nrNeighs < minNeighbors {
			minNeighbors = nrNeighs
			minNode = pNode
		}
	}
	return minNode // could be nil
}

func (g *graph) compress() {
	for {
		pNode := g.findNodeWithNoFlowAndFewestNeighbors()
		if pNode == nil {
			// no node remaining with flow 0
			return
		}
		// fmt.Println("minNode", pNode)

		// replace edges
		neighs := g.getNeighbors(pNode.name)
		for _, pNeigh1 := range neighs {
			for _, pNeigh2 := range neighs {
				if pNeigh1.name >= pNeigh2.name {
					continue
				}
				e1 := g.getEdge(pNode.name, pNeigh1.name)
				e2 := g.getEdge(pNode.name, pNeigh2.name)
				g.replaceOrCreateEdge(pNeigh1.name, pNeigh2.name, e1.weight+e2.weight)
			}
		}

		for _, pNeigh := range neighs {
			g.removeEdge(pNode.name, pNeigh.name)
		}

		delete(g.nodes, pNode.name)
	}
}
