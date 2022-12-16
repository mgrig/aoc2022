package volcano

import "fmt"

func Part2(lines []string) int {
	g := parseGraph(lines)
	fmt.Println("original:\n", g.GraphvizString())

	g.compress()
	fmt.Println("compressed:\n", g.GraphvizString())

	return -1
}
