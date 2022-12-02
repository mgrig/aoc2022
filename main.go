package main

import (
	"aoc2022/common"
	"aoc2022/rps"
	"fmt"
)

func main() {
	// // day 1
	// lines := common.GetLinesFromFile("resources/01_calorie.txt", false)
	// maxCalories := calorie.MostCalories(lines)
	// top3 := calorie.TopThree(lines)
	// fmt.Println("max calories:", maxCalories)
	// fmt.Println("top 3:", top3)

	// day 2
	lines := common.GetLinesFromFile("resources/02_rps.txt", true)
	totalScore := rps.TotalScore(lines)
	fmt.Println("total score:", totalScore)
	totalScoreByOutcome := rps.TotalScoreByOutcome(lines)
	fmt.Println("total score by outcome:", totalScoreByOutcome)
}
