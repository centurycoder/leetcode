// Package solution contains the Go solution for LeetCode 49.
package solution

// groupAnagrams groups strings that contain the same lowercase letters with
// the same frequencies. The order of the groups is unspecified.
func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}

	for _, s := range strs {
		key := [26]int{}
		for _, c := range s {
			key[c-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	result := [][]string{}
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
