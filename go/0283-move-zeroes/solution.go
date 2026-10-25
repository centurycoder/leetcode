// Package solution contains the Go solution for LeetCode 283.
package solution

// moveZeroes moves all zeros to the end of nums while preserving the order of
// non-zero values. It modifies nums in place using O(1) extra space.
func moveZeroes(nums []int) {
	write := 0
	for _, value := range nums {
		if value == 0 {
			continue
		}
		nums[write] = value
		write++
	}
	for write < len(nums) {
		nums[write] = 0
		write++
	}
}
