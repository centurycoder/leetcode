// Package solution contains the Go solution for LeetCode 15.
package solution

import (
	"slices"
)

// threeSum returns the unique triplets in nums that sum to zero.
func threeSum(nums []int) [][]int {

	result := [][]int{}
	// sort array for quick
	slices.Sort(nums)

	var i, j, k int
	len := len(nums)
	for i = 0; i < len-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j = i + 1
		k = len - 1
		for j < len-1 && k > i {
			if j >= k {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				continue
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return result
}
