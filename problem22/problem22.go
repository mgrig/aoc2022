package problem22

import (
	"aoc2022/common"
	"fmt"
	"regexp"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`(\d+)(.*)`)

func Part1(lines []string) int {
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
			pos = pos.moveAhead(steps, board)
		}
	}
	fmt.Printf("r=%d, c=%d, facing=%d\n", pos.r, pos.c, pos.facing)

	pass := 1000*(pos.r+1) + 4*(pos.c+1) + pos.facing

	return pass
}
