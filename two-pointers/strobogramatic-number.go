package main

func isStrobogrammatic(num string) bool {

	left, right := 0, len(num)-1
	strobogrammaticPairs := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	for left <= right {
		leftchar := num[left]
		rightchar := num[right]

		if leftVal, exists := strobogrammaticPairs[leftchar]; !exists || leftVal != rightchar {
			return false
		}
		left++
		right--
	}

	return true
}
