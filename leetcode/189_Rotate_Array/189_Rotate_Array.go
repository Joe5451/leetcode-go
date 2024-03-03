package leetcode

// time complexity is O(n), space complexity is O(1)
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, len(nums) - 1)
}

func reverse(nums []int, startIndex int, endIndex int) {
	for startIndex < endIndex {
		nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
		startIndex++
		endIndex--
	}
}
