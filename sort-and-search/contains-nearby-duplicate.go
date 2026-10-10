package sortandsearch

func containsNearbyDuplicate(nums []int, k int) bool {

	window := make(map[int]bool)
	windowSize := k

	for i := 0; i < len(nums); i++ {
		if window[nums[i]] {
			return true
		}

		window[nums[i]] = true

		if len(window) > windowSize {
			delete(window, nums[i-windowSize])
		}

	}
	return false
}
