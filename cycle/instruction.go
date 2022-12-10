package cycle

import (
	"strconv"
	"strings"
)

type instruction interface {
	nrCycles() int
	apply(regX *int)
}

var _ instruction = noop{}
var _ instruction = addx{}

func parseInstruction(line string) *instruction {
	var ret instruction

	if line == "noop" {
		ret = noop{}
	} else if strings.HasPrefix(line, "addx") {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic("wrong line " + line)
		}
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		ret = addx{
			value: value,
		}
	}

	return &ret
}

// ***

type noop struct {
}

func (n noop) nrCycles() int {
	return 1
}

func (n noop) apply(regX *int) {
	return
}

// ***

type addx struct {
	value int
}

func (ax addx) nrCycles() int {
	return 2
}

func (ax addx) apply(regX *int) {
	*regX += ax.value
}
