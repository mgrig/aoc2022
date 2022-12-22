package problem22

import (
	"aoc2022/common"
	"fmt"
	"regexp"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`(\d+)(.*)`)

func Part2(lines []string) int {
	board := parseBoard(lines[:len(lines)-1]) // excludes last
	cmds := lines[len(lines)-1]
	// fmt.Println(board.String())
	// fmt.Println("cmds:", cmds)

	// get start position
	var pos position
	for c := 0; c < board.nrCols(); c++ {
		if board.matrix[0][c] == 1 {
			pos = newPosition(0, c, 0)
			break
		}
	}

	// faces := createTestFaces()
	faces := createRealFaces()

	for len(cmds) > 0 {
		if strings.HasPrefix(cmds, "R") {
			pos = pos.turnRight()
			cmds = cmds[1:]
		} else if strings.HasPrefix(cmds, "L") {
			pos = pos.turnLeft()
			cmds = cmds[1:]
		} else {
			tokens := re.FindStringSubmatch(cmds)
			if len(tokens) != 3 {
				panic("wrong next cmd in " + cmds)
			}
			cmds = tokens[2]
			steps := common.StringToInt(tokens[1])
			pos = pos.moveAhead(steps, board, faces)
		}
	}
	fmt.Printf("r=%d, c=%d, facing=%d\n", pos.r, pos.c, pos.facing)

	pass := 1000*(pos.r+1) + 4*(pos.c+1) + pos.facing

	return pass
}

// (vscode formatter is an a$$ and messes up the image below. I hope it's still undersandable.)
//
//	    ┌───┬───┐
//		│ 1 │ 2 │
//		├───┼───┘
//		│ 3 │
//
// ┌───┼───┤
// │ 5 │ 4 │
// ├───┼───┘
// │ 6 │
// └───┘
func createRealFaces() map[int]face {
	faces := make(map[int]face)

	// NOT ALL border wrappings are defined! I only did the ones used in my data sample.
	tx_13_62 := newTransformation(-1, 50, 90, 150, 0, 0)
	tx_20_40 := newTransformation(49, 149, 180, 100, 100, 0)
	tx_30_21 := newTransformation(50, 100, 0, 49, 100, 90)
	tx_52_12 := newTransformation(149, -1, 180, 0, 50, 0)
	tx_53_32 := newTransformation(99, 0, 90, 50, 50, 0)
	tx_60_41 := newTransformation(150, 50, 0, 149, 50, 90)
	tx_61_23 := newTransformation(200, 0, 0, 0, 100, 0)

	faces[1] = newFace(1, 50, newCoord(0, 50), map[int]transformation{
		2: tx_52_12.reverse(),
		3: tx_13_62,
	})

	faces[2] = newFace(2, 50, newCoord(0, 100), map[int]transformation{
		0: tx_20_40,
		1: tx_30_21.reverse(),
		3: tx_61_23.reverse(),
	})

	faces[3] = newFace(3, 50, newCoord(50, 50), map[int]transformation{
		0: tx_30_21,
		2: tx_53_32.reverse(),
	})

	faces[4] = newFace(4, 50, newCoord(100, 50), map[int]transformation{
		0: tx_20_40.reverse(),
		1: tx_60_41.reverse(),
	})

	faces[5] = newFace(5, 50, newCoord(100, 0), map[int]transformation{
		2: tx_52_12,
		3: tx_53_32,
	})

	faces[6] = newFace(6, 50, newCoord(150, 0), map[int]transformation{
		0: tx_60_41,
		1: tx_61_23,
		2: tx_13_62.reverse(),
	})

	return faces
}

func createTestFaces() map[int]face {
	faces := make(map[int]face)

	faces[1] = newFace(1, 4, newCoord(0, 8), map[int]transformation{
		0: newTransformation(3, 12, 90, 8, 15, -90),
		2: newTransformation(3, 7, 180, 4, 7, -90),
		3: newTransformation(-1, 8, 90, 4, 3, -90),
	})

	faces[2] = newFace(2, 4, newCoord(4, 0), make(map[int]transformation))
	faces[3] = newFace(3, 4, newCoord(4, 4), map[int]transformation{
		3: newTransformation(3, 4, 90, 0, 8, 0),
	})

	faces[4] = newFace(4, 4, newCoord(4, 8), map[int]transformation{
		0: newTransformation(7, 12, 90, 8, 12, 0),
	})

	faces[5] = newFace(5, 4, newCoord(8, 8), map[int]transformation{
		1: newTransformation(12, 8, 0, 7, 3, 180),
	})

	faces[6] = newFace(6, 4, newCoord(8, 12), make(map[int]transformation))

	return faces
}

func getFace(posAbs coord, faces map[int]face) face {
	for faceId := 1; faceId <= 6; faceId++ {
		if faces[faceId].contains(posAbs) {
			return faces[faceId]
		}
	}
	panic(fmt.Sprintf("no face found for posAbs=%s", posAbs))
}
