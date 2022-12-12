package monkey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type item struct {
	worryLevel WorryLevel
}

var reItems = regexp.MustCompile(`Starting items: (.+)`)

func parseItems(line string) []item {
	line = strings.Trim(line, " ")

	strItems := reItems.FindStringSubmatch(line)[1]
	tokens := strings.Split(strItems, ", ")
	ret := make([]item, len(tokens))
	for i, token := range tokens {
		value, err := strconv.Atoi(token)
		// value, err := strconv.ParseUint(token, 10, 64)
		if err != nil {
			panic(err)
		}
		ret[i] = item{
			// worryLevel: newIntWorryLevel(value),
			// worryLevel: newBigWorryLevel(value),
			// worryLevel: newU64WorryLevel(uint64(value)),
			worryLevel: newSmartWL(value),
		}
	}
	return ret
}

func (i item) String() string {
	return fmt.Sprintf("%v", i.worryLevel)
}
