package problem25

func Part1(lines []string) string {
	sum := 0
	for _, line := range lines {
		sum += newSnafu(line).toDec()
	}

	return fromDec(sum).value
}
