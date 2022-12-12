package monkey

import (
	"strconv"
	"strings"
)

type test interface {
	passes(value WorryLevel) bool
}

var _ test = divisibleBy{}

func parseTest(line string) test {
	line = strings.Trim(line, " ")

	var t test

	if !strings.HasPrefix(line, "Test: divisible by ") {
		panic("wrong test line " + line)
	}
	line = line[19:]

	value, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	t = divisibleBy{
		value: value,
	}

	return t
}

// ****

type divisibleBy struct {
	value int
}

func (db divisibleBy) passes(value WorryLevel) bool {
	return value.divisibleBy(db.value)
}
