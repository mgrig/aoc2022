package distress

import (
	"github.com/alecthomas/participle/v2"
)

func Sum1(lines []string) int {
	parser, err := participle.Build[Root](
		participle.Union[Element](Number{}, List{}),
	)
	if err != nil {
		panic(err)
	}

	sum1 := 0

	for i := 0; i < len(lines)/2; i++ {
		leftLine := lines[2*i]
		rightLine := lines[2*i+1]

		left, err := parser.ParseString("", leftLine)
		if err != nil {
			panic(err)
		}
		right, err := parser.ParseString("", rightLine)
		if err != nil {
			panic(err)
		}

		cmp := left.Elem.compareTo(right.Elem)
		if cmp == -1 {
			sum1 += (i + 1)
		}
	}
	return sum1
}
