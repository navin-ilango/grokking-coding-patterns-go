package binarysearch

func firstBadVersion(n int) int {

	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid // The first bad version is mid or to the left
		} else {
			left = mid + 1 // The first bad version is to the right
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return true // Placeholder implementation; replace with actual logic to check if a version is bad
}
