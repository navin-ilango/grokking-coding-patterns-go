package slidingwindow

func maxProfit(prices []int) int {
	buyPointer := 0
	sellPointer := 1
	maxProfit := 0

	for sellPointer < len(prices) {
		currentProfit := prices[sellPointer] - prices[buyPointer]
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		if prices[sellPointer] < prices[buyPointer] {
			buyPointer = sellPointer
		}
		sellPointer++
	}
	return maxProfit
}
