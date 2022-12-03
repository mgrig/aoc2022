package rucksack

import "fmt"

func SumPriorities(lines []string) int {
	sum := 0
	for _, line := range lines {
		commonItem := commonItemType(line)
		sum += itemTypeToPrio(commonItem)
	}
	return sum
}

func SumPrioritiesGroup(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)/3; i++ {
		commonItem := commonItemInGroup(lines[3*i], lines[3*i+1], lines[3*i+2])
		sum += itemTypeToPrio(commonItem)
	}
	return sum
}

func commonItemInGroup(line1, line2, line3 string) string {
	inRucksack1 := make(map[string]bool)
	for i := 0; i < len(line1); i++ {
		inRucksack1[string(line1[i])] = true
	}

	in12 := make(map[string]bool)
	for i := 0; i < len(line2); i++ {
		item := string(line2[i])
		if inRucksack1[item] {
			in12[item] = true
		}
	}

	for i := 0; i < len(line3); i++ {
		item := string(line3[i])
		if in12[item] {
			return item
		}
	}

	panic("no common item in group")
}

func itemTypeToPrio(it string) int {
	run := int(it[0])
	if run >= 97 && run <= 122 {
		return run - 96
	}
	if run >= 65 && run <= 90 {
		return run - 65 + 27
	}
	panic("wrong item type " + it)
}

func commonItemType(line string) string {
	if len(line)%2 != 0 {
		panic(fmt.Sprintf("wrong line length: %d, line: %s", len(line), line))
	}
	indexComp2 := len(line) / 2

	inComp1 := make(map[string]bool)
	for i := 0; i < indexComp2; i++ {
		inComp1[string(line[i])] = true
	}
	// fmt.Println(line, " > ", inComp1)

	for i := indexComp2; i < len(line); i++ {
		item := string(line[i])
		_, ok := inComp1[item]
		if ok {
			// item from comp2 is present in comp1 as well. stop searching
			return item
		}
	}
	panic("no common item found, line: " + line)
}
