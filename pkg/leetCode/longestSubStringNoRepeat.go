package leetCode

func LongestSubstringWithNoRepeatChars(s string) int {
	m := make(map[rune]int)
	current, longest := 0, 0

	for _, r := range s {
		//fmt.Printf("%c|", r)
		if _, found := m[r]; found {
			if current > longest {
				longest = current
			}
			m = make(map[rune]int)
			current = 0
		}

		m[r] = 1
		current++
	}

	return longest
}
