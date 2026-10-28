package solution

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j && i < len(height) && j > 0 {
		area := (j - i) * min(height[i], height[j])
		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
