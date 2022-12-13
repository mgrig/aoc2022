package main

import (
	"aoc2022/common"
	"aoc2022/distress"
	"fmt"
)

func main() {
	// // day 1
	// lines := common.GetLinesFromFile("resources/01_calorie.txt", false)
	// maxCalories := calorie.MostCalories(lines)
	// top3 := calorie.TopThree(lines)
	// fmt.Println("max calories:", maxCalories)
	// fmt.Println("top 3:", top3)

	// // day 2
	// lines := common.GetLinesFromFile("resources/02_rps.txt", true)
	// totalScore := rps.TotalScore(lines)
	// fmt.Println("total score:", totalScore)
	// totalScoreByOutcome := rps.TotalScoreByOutcome(lines)
	// fmt.Println("total score by outcome:", totalScoreByOutcome)

	// // day 3
	// lines := common.GetLinesFromFile("resources/03_rucksack.txt", true)
	// sumPrio := rucksack.SumPriorities(lines)
	// fmt.Println("sum prio:", sumPrio)
	// sumPrioGroup := rucksack.SumPrioritiesGroup(lines)
	// fmt.Println("sum prio group:", sumPrioGroup)

	// // day 4
	// lines := common.GetLinesFromFile("resources/04_ranges.txt", true)
	// commonRanges := ranges.CommonRanges(lines)
	// fmt.Println("common ranges: ", commonRanges)
	// commonRangesPartial := ranges.CommonRangesPartial(lines)
	// fmt.Println("common ranges partial: ", commonRangesPartial)

	// // day 5
	// lines := common.GetLinesFromFile("resources/05_stacks_test.txt", false, false)
	// topCrates := stacks.TopCrates(lines, false)
	// fmt.Println("top crates:", topCrates)
	// topCratesGrouped := stacks.TopCrates(lines, true)
	// fmt.Println("top crates moveGrouped:", topCratesGrouped)

	// // day 6
	// line := common.GetLinesFromFile("resources/06_tuning.txt", true, true)[0]
	// startOfPacket := tuning.StartOfPacket(line, 4)
	// fmt.Println("start packet:", startOfPacket)
	// startOfMessage := tuning.StartOfPacket(line, 14)
	// fmt.Println("start message:", startOfMessage)

	// // day 7
	// lines := common.GetLinesFromFile("resources/07_spaceleft.txt", true, true)
	// spaceleft.ParseAndShowTree(lines)
	// totalBelow100k := spaceleft.TotalSizeDirsBelow100k(lines)
	// fmt.Println("total below 100k:", totalBelow100k)
	// sizeDeleteDir := spaceleft.DeleteOneDir(lines)
	// fmt.Println("size delete dir:", sizeDeleteDir)

	// // day 8
	// lines := common.GetLinesFromFile("resources/08_treetop.txt", true, true)
	// countVisible := treetop.CountVisible(lines)
	// fmt.Println("visible trees:", countVisible)
	// bestScore := treetop.BestScenicScore(lines)
	// fmt.Println("best score:", bestScore)

	// // day 9
	// lines := common.GetLinesFromFile("resources/09_rope.txt", true, true)
	// countTail := rope.CountTail(lines)
	// fmt.Println("count tail positions:", countTail)
	// countTail10 := rope.CountTail10(lines)
	// fmt.Println("count tail positions 10 knots:", countTail10)

	// // day 10
	// lines := common.GetLinesFromFile("resources/10_cycle.txt", true, true)
	// sum1 := cycle.Sum1(lines)
	// fmt.Println("sum1:", sum1)
	// cycle.Screen(lines)

	// // day 11
	// lines := common.GetLinesFromFile("resources/11_monkey.txt", true, true)
	// // fmt.Println(monkey.Business20Rounds(lines, 20, true))
	// fmt.Println(monkey.Business20Rounds(lines, 10000, false))

	// // day 12
	// lines := common.GetLinesFromFile("resources/12_hill.txt", true, true)
	// nrSteps := hill.NrSteps(lines)
	// fmt.Println("nr steps up:", nrSteps)
	// nrStepsDown := hill.NrStepsDown(lines)
	// fmt.Println("nr steps down:", nrStepsDown)

	// day 13
	lines := common.GetLinesFromFile("resources/13_distress_test.txt", true, true)
	sum1 := distress.Sum1(lines)
	fmt.Println("sum1", sum1)
}
