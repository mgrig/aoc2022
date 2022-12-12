package monkey

import (
	"strconv"
	"strings"
)

type operation interface {
	apply(old WorryLevel) WorryLevel
}

var _ operation = addScalar{}
var _ operation = multScalar{}
var _ operation = square{}

func parseOperation(line string) operation {
	line = strings.Trim(line, " ")

	var op operation

	if !strings.HasPrefix(line, "Operation: new = old ") {
		panic("wrong operation line " + line)
	}
	line = line[21:]

	if strings.HasPrefix(line, "+ ") {
		line = line[2:]
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		op = addScalar{
			value: value,
		}
	} else if strings.HasPrefix(line, "* ") {
		line = line[2:]
		if line == "old" {
			op = square{}
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			op = multScalar{
				value: value,
			}
		}
	} else {
		panic("wrong math op " + line)
	}

	return op
}

// ****

type addScalar struct {
	value int
}

func (as addScalar) apply(old WorryLevel) WorryLevel {
	return old.addInt(as.value)
}

// ****

type multScalar struct {
	value int
}

func (ms multScalar) apply(old WorryLevel) WorryLevel {
	return old.multInt(ms.value)
}

// ****

type square struct{}

func (s square) apply(old WorryLevel) WorryLevel {
	return old.square()
}
