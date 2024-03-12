package leetcode

// time complexity is O(n), space complexity is O(1)
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	sellSum, buySum := 0, 0

	for i := 0; i < len(prices) - 1; i++ {
		if prices[i] < prices[i + 1] {
			sellSum += prices[i + 1]
			buySum += prices[i]
		}
	}

	return sellSum - buySum
}
