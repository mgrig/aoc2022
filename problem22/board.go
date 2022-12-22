package problem22

type board struct {
	matrix [][]int // 0 - off, 1 - open, 2 - wall
}

func (b *board) nrRows() int {
	return len(b.matrix)
}

func (b *board) nrCols() int {
	return len(b.matrix[0])
}

func parseBoard(lines []string) *board {
	rows := len(lines)
	mat := make([][]int, rows)

	cols := 0
	for _, line := range lines {
		if len(line) > cols {
			cols = len(line)
		}
	}

	for r, line := range lines {
		mat[r] = make([]int, cols)
		for c, run := range line {
			// 32 > 0 (off)
			// 46 > 1 (open)
			// 35 > 2 (wall)
			switch run {
			case 46:
				mat[r][c] = 1
			case 35:
				mat[r][c] = 2
			}
		}
	}

	return &board{
		matrix: mat,
	}
}

func (b *board) String() string {
	ret := ""
	for r := 0; r < len(b.matrix); r++ {
		for c := 0; c < len(b.matrix[r]); c++ {
			value := b.matrix[r][c]
			switch value {
			case 0:
				ret += " "
			case 1:
				ret += "."
			case 2:
				ret += "#"
			default:
				panic("wrong value in matrix")
			}
		}
		ret += "\n"
	}
	return ret
}

func (b *board) moveWithWrap(p position) position {
	var newR, newC int
	switch p.facing {
	case 0: // right
		newR = p.r
		newC = p.c + 1

		if newC >= b.nrCols() || b.matrix[newR][newC] == 0 { // wrap around
			for c := 0; c < newC; c++ {
				if b.matrix[newR][c] != 0 {
					newC = c
					break
				}
			}
		}
	case 2: // left
		newR = p.r
		newC = p.c - 1

		if newC < 0 || b.matrix[newR][newC] == 0 { // wrap around
			for c := b.nrCols() - 1; c >= 0; c-- {
				if b.matrix[newR][c] != 0 {
					newC = c
					break
				}
			}
		}
	case 1: // down
		newR = p.r + 1
		newC = p.c

		if newR >= b.nrRows() || b.matrix[newR][newC] == 0 { // wrap around
			for r := 0; r < newR; r++ {
				if b.matrix[r][newC] != 0 {
					newR = r
					break
				}
			}
		}
	case 3: // up
		newR = p.r - 1
		newC = p.c

		if newR < 0 || b.matrix[newR][newC] == 0 { // wrap around
			for r := b.nrRows() - 1; r >= 0; r-- {
				if b.matrix[r][newC] != 0 {
					newR = r
					break
				}
			}
		}
	default:
		panic("wrong facing")
	}
	return newPosition(newR, newC, p.facing)
}
