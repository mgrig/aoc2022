package tree

func Part1(lines []string) int {
	root, _ := parseNodes(lines)
	return root.getValue()
}
