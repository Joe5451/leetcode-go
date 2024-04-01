package leetcode

// time complexity is O(n), space complexity is O(1)
func majorityElement(nums []int) int {
	result := nums[0]
	count := 0
	for _, num := range nums {
		if count == 0 {
			result = num
		}

		// number of majority element gte (n/2 + 1), so finally its count would not be 0
		if num == result {
			count++
		} else {
			count--
		}
	}

	return result
}
