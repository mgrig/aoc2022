package cycle

import (
	"aoc2022/common"
	"fmt"
)

// ****

type instructionExecution struct {
	instr           instruction
	completedCycles int
}

func newInstructionExecution(line string) *instructionExecution {
	return &instructionExecution{
		instr:           *parseInstruction(line),
		completedCycles: 0,
	}
}

func (ie *instructionExecution) tick(regX *int) (finished bool) {
	ie.completedCycles += 1

	if ie.completedCycles > ie.instr.nrCycles() {
		panic("wrong cycles number")
	}

	if ie.completedCycles == ie.instr.nrCycles() {
		ie.instr.apply(regX)
		return true
	}

	return false
}

func Sum1(lines []string) int {
	nextLineIndex := 0
	regX := 1
	var currentInstructionExecution *instructionExecution
	totalSS := 0

	cycle := 1
	for ; ; cycle++ {
		if cycle >= 20 && (cycle-20)%40 == 0 {
			ss := cycle * regX
			fmt.Printf("ss during cycle %d: %d\n", cycle, ss)
			totalSS += ss
		}

		if currentInstructionExecution == nil {
			// pick next instruction
			if nextLineIndex == len(lines) {
				break
			}
			line := lines[nextLineIndex]
			nextLineIndex++

			currentInstructionExecution = newInstructionExecution(line)
		}

		finished := currentInstructionExecution.tick(&regX)
		if finished {
			// fmt.Printf("regX = %d after instr line %d, after cycle %d\n", regX, nextLineIndex, cycle)
			currentInstructionExecution = nil
		}
	}

	fmt.Println("cycles", cycle)

	return totalSS
}

func Screen(lines []string) {
	pixels := ""

	nextLineIndex := 0
	regX := 1
	var currentInstructionExecution *instructionExecution

	cycle := 1
	for ; ; cycle++ {
		if currentInstructionExecution == nil {
			// pick next instruction
			if nextLineIndex == len(lines) {
				break
			}
			line := lines[nextLineIndex]
			nextLineIndex++

			currentInstructionExecution = newInstructionExecution(line)
		}

		drawnPos := (cycle - 1) % 40
		if common.IntAbs(drawnPos-regX) <= 1 {
			pixels += "#"
		} else {
			pixels += "."
		}
		if drawnPos == 39 {
			pixels += "\n"
		}
		if cycle <= 12 {
			fmt.Printf("during cycle %2d: drawPos=%2d, X=%2d %s\n", cycle, drawnPos, regX, pixels)
		}

		finished := currentInstructionExecution.tick(&regX)
		if finished {
			// fmt.Printf("regX = %d after instr line %d, after cycle %d\n", regX, nextLineIndex, cycle)
			currentInstructionExecution = nil
		}
	}

	fmt.Println(pixels)
}
