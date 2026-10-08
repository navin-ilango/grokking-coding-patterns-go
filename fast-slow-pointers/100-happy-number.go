package main

func isHappy(num int) bool {

	ptrSlow := num
	ptrFast := digitsSquareSum(num)

	for ptrFast != 1 && ptrSlow != ptrFast {
		ptrSlow = digitsSquareSum(ptrSlow)
		ptrFast = digitsSquareSum(digitsSquareSum(ptrFast))
	}
	if ptrFast == 1 {
		return true
	}
	return false
}

func digitsSquareSum(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}
