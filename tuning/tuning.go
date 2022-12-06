package tuning

func StartOfPacket(line string, uniqueLen int) int {
	runes := []rune(line)
	for i := uniqueLen - 1; i < len(runes); i++ {
		if unique(runes[i-uniqueLen+1 : i+1]) {
			// if runes[i-3] != runes[i-2] && runes[i-3] != runes[i-1] && runes[i-3] != runes[i] &&
			// 	runes[i-2] != runes[i-1] && runes[i-2] != runes[i] &&
			// 	runes[i-1] != runes[i] {
			return i + 1
		}
	}
	panic("sop not found")
}

func unique(runes []rune) bool {
	set := make(map[rune]bool, len(runes))
	for _, r := range runes {
		_, exists := set[r]
		if exists {
			return false
		}
		set[r] = true
	}
	return true
}
