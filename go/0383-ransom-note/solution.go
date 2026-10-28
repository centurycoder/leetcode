package solution

func canConstruct(ransomNote string, magazine string) bool {

	m := map[rune]int{}
	for _, c := range magazine {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	for _, c := range ransomNote {
		if _, ok := m[c]; ok {
			if m[c] > 0 {
				m[c]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
