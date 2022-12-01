package main

import (
	"aoc2022/calorie"
	"aoc2022/common"
	"fmt"
)

func main() {
	// day 1
	lines := common.GetLinesFromFile("resources/01_calorie.txt", false)
	maxCalories := calorie.MostCalories(lines)
	top3 := calorie.TopThree(lines)
	fmt.Println("max calories:", maxCalories)
	fmt.Println("top 3:", top3)
}
