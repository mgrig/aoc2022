package main

import (
	"aoc2022/common"
	"aoc2022/problem25"
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

	// // day 13
	// lines := common.GetLinesFromFile("resources/13_distress.txt", true, true)
	// sum1 := distress.Sum1(lines)
	// fmt.Println("sum1", sum1)
	// sorted := distress.Sorted(lines)
	// fmt.Println("sorted", sorted)

	// // day 14
	// lines := common.GetLinesFromFile("resources/14_sand.txt", true, true)
	// part1 := sand.Part1(lines)
	// fmt.Println("part1", part1)
	// part2 := sand.Part2(lines)
	// fmt.Println("part2", part2)

	// // day 15
	// lines := common.GetLinesFromFile("resources/15_sensors.txt", true, true)
	// // nrExcludedPositions := sensors.NrExcludedPositionsOnRow(lines, 10)
	// // nrExcludedPositions := sensors.NrExcludedPositionsOnRow(lines, 2000000)
	// // fmt.Println("part1", nrExcludedPositions)
	// t0 := time.Now()
	// tuning := sensors.TuningFreq(lines, 0, 4000000)
	// t1 := time.Now()
	// // tuning := sensors.TuningFreq(lines, 0, 20)
	// fmt.Printf("(%d ms) tuning freq: %d\n", int(t1.Sub(t0).Milliseconds()), tuning)

	// // day 16
	// lines := common.GetLinesFromFile("resources/16_volcano.txt", true, true)
	// // part1 := volcano.Part1(lines)
	// // fmt.Println("part1", part1)
	// //rand.Seed(time.Now().UnixNano())
	// part2 := volcano.Part2(lines)
	// fmt.Println("part2", part2)

	// // day 17
	// line := common.GetLinesFromFile("resources/17_tetris.txt", true, true)[0]
	// part1 := tetris.Part1(line, 2022)
	// fmt.Println("part1", part1)
	// part2 := tetris.Part1(line, 1000000000000)
	// fmt.Println("part2", part2)

	// // day 18
	// lines := common.GetLinesFromFile("resources/18_cubes.txt", true, true)
	// // part1 in git repo. Modified structs for part 2
	// exteriorSides := cubes.ExteriorSides(lines)
	// fmt.Println("part2", exteriorSides)

	// // day 19
	// lines := common.GetLinesFromFile("resources/19_robots_test.txt", true, true)
	// part1 := robots.Part1(lines)
	// fmt.Println("part1", part1)
	// // part2 := robots.Part2(lines)
	// // fmt.Println("part2", part2)

	// // day 20
	// lines := common.GetLinesFromFile("resources/20_ringbuffer.txt", true, true)
	// part1 := ringbuffer.Part1(lines)
	// fmt.Println("part1:", part1)
	// t0 := time.Now()
	// part2 := ringbuffer.Part2(lines, 811589153)
	// t1 := time.Now()
	// fmt.Printf("(%d ms) part2: %d\n", int(t1.Sub(t0).Milliseconds()), part2)

	// // day 21
	// lines := common.GetLinesFromFile("resources/21_tree_test.txt", true, true)
	// part2 := tree.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 22
	// lines := common.GetLinesFromFile("resources/22_board.txt", true, false)
	// part2 := problem22.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 23
	// lines := common.GetLinesFromFile("resources/23.txt", true, true)
	// part2 := problem23.Part2(lines)
	// fmt.Println("part2:", part2)

	// // day 24
	// lines := common.GetLinesFromFile("resources/24.txt", true, true)
	// part1 := problem24.Part1(lines)
	// fmt.Println("part1:", part1)

	// day 25
	lines := common.GetLinesFromFile("resources/25.txt", true, true)
	part1 := problem25.Part1(lines)
	fmt.Println("part1:", part1)
}
