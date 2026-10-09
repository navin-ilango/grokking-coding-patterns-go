package cyclicsort

func findMissingNumber(nums []int) int {

	i := 0
	size := len(nums)

	for i < size {
		correctIndex := nums[i]
		if correctIndex < size && nums[i] != nums[correctIndex] {
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
			continue
		}
		i++
	}

	for i, x := range nums {
		if i != x {
			return i
		}
	}
	return size
}
