package treetop

import "fmt"

// ****
type forest struct {
	matrix [][]int
}

func newForest(r, c int) *forest {
	mat := make([][]int, r)
	for ir := 0; ir < r; ir++ {
		mat[ir] = make([]int, c)
	}

	return &forest{
		matrix: mat,
	}
}

func (f forest) String() string {
	ret := ""
	for r := 0; r < len(f.matrix); r++ {
		row := f.matrix[r]
		line := ""
		for c := 0; c < len(row); c++ {
			line += fmt.Sprintf("%d", row[c])
		}
		line += fmt.Sprintln()
		ret += line
	}

	return ret
}

func parseForest(lines []string) *forest {
	forest := newForest(len(lines), len(lines[0]))

	for r, line := range lines {
		for c, run := range line {
			forest.matrix[r][c] = int(run) - int('0')
		}
	}
	// fmt.Println(forest)
	return forest
}

func CountVisible(lines []string) int {
	forest := parseForest(lines)

	visibleTrees := make(map[int]bool)
	nrR := len(lines)
	nrC := len(lines[0])

	// visible from left
	for r := 0; r < nrR; r++ {
		visibleLevel := -1
		for c := 0; c < nrC; c++ {
			if forest.matrix[r][c] > visibleLevel {
				visibleTrees[rcToOneValue(r, c, nrC)] = true
				visibleLevel = forest.matrix[r][c]
			}
		}
	}

	// visible from right
	for r := 0; r < nrR; r++ {
		visibleLevel := -1
		for c := nrC - 1; c >= 0; c-- {
			if forest.matrix[r][c] > visibleLevel {
				visibleTrees[rcToOneValue(r, c, nrC)] = true
				visibleLevel = forest.matrix[r][c]
			}
		}
	}

	// visible from up
	for c := 0; c < nrC; c++ {
		visibleLevel := -1
		for r := 0; r < nrR; r++ {
			if forest.matrix[r][c] > visibleLevel {
				visibleTrees[rcToOneValue(r, c, nrC)] = true
				visibleLevel = forest.matrix[r][c]
			}
		}
	}

	// visible from down
	for c := 0; c < nrC; c++ {
		visibleLevel := -1
		for r := nrR - 1; r >= 0; r-- {
			if forest.matrix[r][c] > visibleLevel {
				visibleTrees[rcToOneValue(r, c, nrC)] = true
				visibleLevel = forest.matrix[r][c]
			}
		}
	}

	return len(visibleTrees)
}

func rcToOneValue(r, c, nrC int) int {
	return r*nrC + c
}

func BestScenicScore(lines []string) int {
	forest := parseForest(lines)

	nrR := len(lines)
	nrC := len(lines[0])

	bestScore := 0
	for r := 0; r < nrR; r++ {
		for c := 0; c < nrC; c++ {
			score := getScore(r, c, forest)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func getScore(r, c int, forest *forest) int {
	nrR := len(forest.matrix)
	nrC := len(forest.matrix[0])

	// look up
	up := 0
	if r > 0 {
		ir := r - 1
		for ir >= 0 {
			up += 1
			if forest.matrix[ir][c] >= forest.matrix[r][c] {
				break
			}
			ir--
		}
	}

	if r == 3 && c == 2 {
		fmt.Println("debug")
	}

	// look down
	down := 0
	if r < nrR {
		ir := r + 1
		for ir < nrR {
			down += 1
			if forest.matrix[ir][c] >= forest.matrix[r][c] {
				break
			}
			ir++
		}
	}

	// look left
	left := 0
	if c > 0 {
		ic := c - 1
		for ic >= 0 {
			left += 1
			if forest.matrix[r][ic] >= forest.matrix[r][c] {
				break
			}
			ic--
		}
	}

	// look right
	right := 0
	if c < nrC {
		ic := c + 1
		for ic < nrC {
			right += 1
			if forest.matrix[r][ic] >= forest.matrix[r][c] {
				break
			}
			ic++
		}
	}

	return left * right * up * down
}
