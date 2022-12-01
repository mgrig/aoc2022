package calorie

import (
	"sort"
	"strconv"
	"strings"
)

func MostCalories(lines []string) int {
	maxSum := 0
	sum := 0
	for _, line := range lines {
		trimmed := strings.Trim(line, " ")
		if trimmed == "" {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
			// fmt.Println(maxSum)
			continue
		}
		value, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}

		sum += value
	}

	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}

type topThree struct {
	values []int
}

func (top *topThree) Add(value int) {
	top.values = append(top.values, value)
	sort.Ints(top.values)
	if len(top.values) > 3 {
		top.values = top.values[1:]
	}
}

func (top *topThree) GetSum() int {
	return top.values[0] + top.values[1] + top.values[2]
}

func TopThree(lines []string) int {
	var t3 topThree
	sum := 0
	for _, line := range lines {
		trimmed := strings.Trim(line, " ")
		if trimmed == "" {
			t3.Add(sum)
			sum = 0
			// fmt.Println(maxSum)
			continue
		}
		value, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}

		sum += value
	}

	if sum > 0 {
		t3.Add(sum)
	}

	return t3.GetSum()
}
