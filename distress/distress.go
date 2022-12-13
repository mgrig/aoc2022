package distress

import (
	"reflect"
	"sort"

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

func Sorted(lines []string) int {
	parser, err := participle.Build[Root](
		participle.Union[Element](Number{}, List{}),
	)
	if err != nil {
		panic(err)
	}

	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")

	all := make([]Element, len(lines))
	var sep1 Element
	var sep2 Element
	for i := 0; i < len(lines); i++ {
		root, err := parser.ParseString("", lines[i])
		if err != nil {
			panic(err)
		}
		all[i] = root.Elem

		if i == len(lines)-2 {
			sep1 = root.Elem
		}

		if i == len(lines)-1 {
			sep2 = root.Elem
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].compareTo(all[j]) == -1
	})

	ret := 1
	for i := 0; i < len(all); i++ {
		if reflect.DeepEqual(all[i], sep1) || reflect.DeepEqual(all[i], sep2) {
			ret *= (i + 1)
		}
	}

	return ret
}
