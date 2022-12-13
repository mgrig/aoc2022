package distress

import (
	"aoc2022/common"
	"fmt"
	"strings"
)

type Root struct {
	Elem Element `@@`
}

type Element interface {
	String() string
	compareTo(other Element) int // -1 if this < other; 0 if this == other; +1 if this > other
}

var _ Element = Number{}
var _ Element = List{}

// ***

type Number struct {
	Value int `@Int`
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}
func (n Number) compareTo(other Element) int {
	otherInt, ok := other.(Number)
	if ok {
		return common.IntSgn(n.Value - otherInt.Value)
	}

	thisAsList := newList(n)
	return thisAsList.compareTo(other)
}

// ***

type List struct {
	Elements []Element `"["(@@(","@@)*)?"]"`
}

func newList(elements ...Element) List {
	return List{
		Elements: elements,
	}
}

func (l List) String() string {
	elementsStr := make([]string, len(l.Elements))
	for i, elem := range l.Elements {
		elementsStr[i] = fmt.Sprint(elem)
	}
	return fmt.Sprintf("[%s]", strings.Join(elementsStr, ", "))
}

func (l List) compareTo(other Element) int {
	otherInt, ok := other.(Number)
	if ok {
		other = newList(otherInt)
	}

	otherList := other.(List)
	for i := 0; ; i++ {
		if i == len(l.Elements) {
			if i == len(otherList.Elements) {
				return 0
			} else {
				return -1 // this list ended before other
			}
		}
		if i == len(otherList.Elements) {
			return 1 // other list ended before this
		}

		thisElem := l.Elements[i]
		otherElem := otherList.Elements[i]
		elemCompare := thisElem.compareTo(otherElem)
		if elemCompare != 0 {
			return elemCompare
		}
		// elements are equal, continue comparison
	}
}
