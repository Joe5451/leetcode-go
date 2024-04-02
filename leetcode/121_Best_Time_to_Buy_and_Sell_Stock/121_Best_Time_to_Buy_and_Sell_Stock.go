package leetcode

// time complexity is O(n), space complexity is O(1)
func maxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price > min && price - min > maxProfit {
			maxProfit = price - min
		}

		if price < min {
			min = price
		}
	}

	return maxProfit
}
