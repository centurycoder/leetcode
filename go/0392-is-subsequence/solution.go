package solution

func isSubsequence(s string, t string) bool {

	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
			continue
		} else {
			// s[i] ! = t[j], move ahead j
			for j < len(t) && t[j] != s[i] {
				j++
			}
		}

	}
	if i == len(s) {
		return true
	}

	return false
}
