package leetcode

// time complexity is O(n), space complexity is O(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	product := 1
	for i, j := 0, 0; j < len(nums); j++ {
		product *= nums[j]
		for product >= k && i < j {
			product /= nums[i]
			i++
		}

		if product < k {
			result += j - i + 1
		}
	}

	return result
}
