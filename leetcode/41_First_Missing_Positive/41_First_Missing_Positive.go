package leetcode

import "math"

// time complexity is O(n), space complexity is O(1)
func firstMissingPositive(nums []int) int {
	exceedMin := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] >= len(nums) && nums[i] < exceedMin {
			exceedMin = nums[i]
		}

		if nums[i] >= 0 && nums[i] < len(nums) && i != nums[i] && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			i--
		}
	}

	min := 1
	for _, num := range nums {
		if min == num {
			min++
		}
	}

	if min == exceedMin {
		min++
	}

	return min
}
