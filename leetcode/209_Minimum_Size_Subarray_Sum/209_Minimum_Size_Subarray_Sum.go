package leetcode

// Sliding Window
// time complexity is O(n), space complexity is O(1)
// total count of internal loop would not exceed n
func minSubArrayLen(target int, nums []int) int {
	start := 0
	minLen := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			currentLen := i - start + 1
			if minLen == 0 || currentLen < minLen {
				minLen = currentLen
			}

			sum -= nums[start]
			start++
		}
	}

	return minLen
}
