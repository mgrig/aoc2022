package monkey

import (
	"aoc2022/common"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type monkey struct {
	id          int
	items       []item
	op          operation
	t           test
	targetTrue  int
	targetFalse int
	nrInspects  int
}

func (m *monkey) throwItemTo(itemIndex int, target *monkey) {
	it := m.items[itemIndex]
	m.items = append(m.items[:itemIndex], m.items[itemIndex+1:]...)

	target.items = append(target.items, it)
}

func parseMonkey(lines []string) *monkey {
	if len(lines) != 6 {
		panic(fmt.Sprintf("wrong nr of lines for monkey, expected 6 was %d", len(lines)))
	}

	str := common.GetOneRegexGroup(regexp.MustCompile(`.*Monkey (\d+):`), lines[0])
	id, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	items := parseItems(lines[1])
	op := parseOperation(lines[2])
	t := parseTest(lines[3])
	str = common.GetOneRegexGroup(regexp.MustCompile(`.*If true: throw to monkey (\d+)`), lines[4])
	targetTrue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	str = common.GetOneRegexGroup(regexp.MustCompile(`.*If false: throw to monkey (\d+)`), lines[5])
	targetFalse, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return &monkey{
		id:          id,
		items:       items,
		op:          op,
		t:           t,
		targetTrue:  targetTrue,
		targetFalse: targetFalse,
		nrInspects:  0,
	}
}

func parseAlls(lines []string) []*monkey {
	monkeys := make([]*monkey, len(lines)/6)
	for m := 0; m < len(monkeys); m++ {
		monkeyLines := lines[m*6 : m*6+6]
		monkeys[m] = parseMonkey(monkeyLines)
	}
	return monkeys
}

func Business20Rounds(lines []string, rounds int, divideBy3 bool) int {
	monkeys := parseAlls(lines)

	debug := false
	for round := 1; round <= rounds; round++ {
		for m, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkeys[m].nrInspects++
				item := monkey.items[0]
				if debug {
					fmt.Printf("Monkey %d inspects item with worry level %d\n", monkey.id, item.worryLevel)
				}
				newWorryLevel := monkey.op.apply(item.worryLevel)
				if debug {
					fmt.Printf("  worry level after operation: %d\n", newWorryLevel)
				}

				if divideBy3 {
					newWorryLevel = newWorryLevel.divInt(3)
					// fmt.Printf("  worry level after /3: %d\n", newWorryLevel)
				}

				target := monkey.targetFalse
				if monkey.t.passes(newWorryLevel) {
					target = monkey.targetTrue
					if debug {
						fmt.Printf("  (pass) ")
					}
				} else {
					if debug {
						fmt.Printf("  (fail) ")
					}
				}
				monkey.items[0].worryLevel = newWorryLevel
				monkey.throwItemTo(0, monkeys[target])
				if debug {
					fmt.Printf("item thrown to monkey %d\n", target)
				}
			}
		}

		// show monkey items
		fmt.Printf("\nAfter round %d\n", round)
		for _, monkey := range monkeys {
			fmt.Printf("Monkey %d: %v\n", monkey.id, monkey.items)
		}
	}
	fmt.Println()

	times := make([]int, len(monkeys))
	for m, monkey := range monkeys {
		fmt.Printf("Monkey %d: %d times\n", monkey.id, monkey.nrInspects)
		times[m] = monkey.nrInspects
	}

	sort.Sort(sort.Reverse(sort.IntSlice(times)))
	fmt.Println(times)

	return times[0] * times[1]
}
