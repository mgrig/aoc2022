package ranges

import (
	"strconv"
	"strings"
)

type elfRange struct {
	min, max int
}

func newElfRange(min, max int) *elfRange {
	if min > max {
		panic("min > max")
	}
	return &elfRange{
		min: min,
		max: max,
	}
}

// input "2-4"
func elfRangeFromString(str string) *elfRange {
	tokens := strings.Split(str, "-")
	min, _ := strconv.Atoi(tokens[0])
	max, _ := strconv.Atoi(tokens[1])
	return newElfRange(min, max)
}

func (er *elfRange) contains(other *elfRange) bool {
	return er.min <= other.min && er.max >= other.max
}

func CommonRanges(lines []string) int {
	sum := 0
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		elf1 := elfRangeFromString(tokens[0])
		elf2 := elfRangeFromString(tokens[1])
		if elf1.contains(elf2) || elf2.contains(elf1) {
			sum += 1
		}
	}
	return sum
}

func CommonRangesPartial(lines []string) int {
	sum := 0
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		elf1 := elfRangeFromString(tokens[0])
		elf2 := elfRangeFromString(tokens[1])
		if partialOverlap(elf1, elf2) {
			sum += 1
		}
	}
	return sum
}

func partialOverlap(elf1, elf2 *elfRange) bool {
	return !(elf1.max < elf2.min || elf1.min > elf2.max)
}
