package solution

func twoSum(numbers []int, target int) []int {
	i := 0
	j := 1
	for i < j && j < len(numbers) {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			j++
		} else {
			i++
			j = i + 1
		}

	}
	return nil
}
