package leetcode

// time complexity is O(n), space complexity is O(1)
func timeRequiredToBuy(tickets []int, k int) int {
	result := 0
	for i, num := range tickets {
		if i <= k {
			result += min(num, tickets[k])
		} else {
			result += min(num, tickets[k] - 1)
		}
	}

	return result
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
