package rps

import (
	"strings"
)

var shapeScore = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

func shapeFromInputChar(input string) string {
	if input == "A" || input == "X" {
		return "R"
	}

	if input == "B" || input == "Y" {
		return "P"
	}

	if input == "C" || input == "Z" {
		return "S"
	}
	panic("wrong input: " + input)
}

// input "A Y" -> 8 (2 for Paper + 6 for win)
func roundScore(line string) int {
	tokens := strings.Split(line, " ")
	if len(tokens) != 2 {
		panic("bzzt: wrong line: " + line)
	}
	opponentShape := shapeFromInputChar(tokens[0])
	myShape := shapeFromInputChar(tokens[1])

	var roundScore int
	if opponentShape == myShape {
		// draw
		roundScore = 3
	} else if (opponentShape == "R" && myShape == "S") ||
		(opponentShape == "S" && myShape == "P") ||
		(opponentShape == "P" && myShape == "R") {
		// opponent wins
		roundScore = 0
	} else {
		// I win
		roundScore = 6
	}

	return roundScore + shapeScore[myShape]
}

func TotalScore(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		totalScore += roundScore(line)
	}
	return totalScore
}

func roundScoreByOutcome(line string) int {
	tokens := strings.Split(line, " ")
	if len(tokens) != 2 {
		panic("bzzt: wrong line: " + line)
	}
	opponentShape := shapeFromInputChar(tokens[0])
	var myShape string
	var roundScore int

	if tokens[1] == "X" {
		// need to lose
		roundScore = 0
		switch opponentShape {
		case "R":
			myShape = "S"
		case "S":
			myShape = "P"
		case "P":
			myShape = "R"
		}
	} else if tokens[1] == "Y" {
		// need to draw
		roundScore = 3
		myShape = opponentShape
	} else if tokens[1] == "Z" {
		// need to win
		roundScore = 6
		switch opponentShape {
		case "R":
			myShape = "P"
		case "S":
			myShape = "R"
		case "P":
			myShape = "S"
		}
	} else {
		panic("bzzt")
	}

	return roundScore + shapeScore[myShape]
}

func TotalScoreByOutcome(lines []string) int {
	totalScore := 0
	for _, line := range lines {
		totalScore += roundScoreByOutcome(line)
	}
	return totalScore
}
