// Package solution contains the Go solution for LeetCode 128.
package solution

// longestConsecutive returns the length of the longest consecutive sequence.
func longestConsecutive(nums []int) int {
	// Intentionally left empty per request. The return keeps this scaffold compilable.

	lookup := map[int]struct{}{}

	// first round, just make up the look up table
	for _, num := range nums {
		lookup[num] = struct{}{}
	}

	maxCount := 0
	count := 0

	for num := range lookup {
		// if the previos number not in sequence, then we are the start of the sequence， check remaining numbers
		if _, ok := lookup[num-1]; !ok {
			count = 1
			for {
				if _, ok := lookup[num+count]; ok {
					count++
				} else {
					break
				}
			}
			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}
