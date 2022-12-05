package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetLinesFromFile(fileName string, skipEmpty bool, trim bool) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if trim {
			line = strings.Trim(line, " ")
		}
		if skipEmpty && len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func ToIntegerValues(lines []string) []int {
	values := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(strings.Trim(line, " "))
		if err != nil {
			panic(err)
		}

		values[i] = value
	}

	return values
}
