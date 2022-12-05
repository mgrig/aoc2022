package stacks

import (
	"regexp"
	"strconv"
)

type stack struct {
	crates []string
}

func (s *stack) push(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *stack) pop() string {
	if len(s.crates) == 0 {
		panic("pop from empty")
	}
	last := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return last
}

func (s stack) getLast() string {
	if len(s.crates) == 0 {
		panic("getLast from empty")
	}
	return s.crates[len(s.crates)-1]
}

func move(n int, from, to *stack, moveGrouped bool) {
	if n > len(from.crates) {
		panic("n too large")
	}

	if moveGrouped {
		offset := len(from.crates) - n
		for i := 0; i < n; i++ {
			to.push(from.crates[offset+i])
		}
		for i := 0; i < n; i++ {
			from.pop()
		}
	} else {
		for i := 0; i < n; i++ {
			crate := from.pop()
			to.push(crate)
		}
	}
}

func TopCrates(lines []string, moveGrouped bool) string {
	// get index of empty line
	indexSeparator := -1
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			indexSeparator = i
			break
		}
	}
	if indexSeparator == -1 {
		panic("index separator not found")
	}
	// fmt.Println("index separator:", indexSeparator)

	// get nr stacks
	nrStacks := (len(lines[indexSeparator-1]) + 1) / 4
	// fmt.Printf("line:>%s< len: %d\n", lines[indexSeparator-1], len(lines[indexSeparator-1]))
	// fmt.Println("nr stacks:", nrStacks)

	// init stacks
	stacks := make([]stack, nrStacks)
	for i := 0; i < nrStacks; i++ {
		stacks[i] = stack{
			crates: make([]string, 0),
		}
	}

	// parse lines before index for start map
	for i := indexSeparator - 2; i >= 0; i-- {
		line := lines[i]
		// fmt.Printf("%d: >%s<\n", i, line)

		for j := 0; j < nrStacks; j++ {
			crate := string(line[4*j+1])
			if crate != " " {
				stacks[j].push(crate)
			}
		}
	}
	// for i := 0; i < nrStacks; i++ {
	// 	fmt.Printf("stack %d: %v\n", i, stacks[i])
	// }

	// for lines after index:
	// - parse each movement
	// - apply movement on stacks state
	pattern := `move (?P<nr>\d+) from (?P<from>\d+) to (?P<to>\d+)` // move 1 from 2 to 1
	re := regexp.MustCompile(pattern)
	for i := indexSeparator + 1; i < len(lines); i++ {
		tokens := re.FindStringSubmatch(lines[i])
		nr, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[2])
		to, _ := strconv.Atoi(tokens[3])
		move(nr, &stacks[from-1], &stacks[to-1], moveGrouped)
	}

	tops := ""
	for s := 0; s < nrStacks; s++ {
		tops += stacks[s].getLast()
	}

	return tops
}
