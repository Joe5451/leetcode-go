package leetcode

// time complexity is O(n), space complexity is O(1)
func countSubarrays(nums []int, minK int, maxK int) int64 {
	// Number of subarrays with gte minK and lte maxK
	result := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] < minK || nums[right] > maxK {
			left = right + 1
		} else {
			result += right - left + 1
		}
	}

	// Number of subarrays with gte minK and lt maxK
	gteMinkLtMaxk := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] < minK || nums[right] >= maxK {
			left = right + 1
		} else {
			gteMinkLtMaxk += right - left + 1
		}
	}

	// Number of subarrays with gt minK and lte maxK
	gtMinkLteMaxk := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] <= minK || nums[right] > maxK {
			left = right + 1
		} else {
			gtMinkLteMaxk += right - left + 1
		}
	}

	// Number of subarrays with gt minK and lt maxK
	gtMinkLtMaxk := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] <= minK || nums[right] >= maxK {
			left = right + 1
		} else {
			gtMinkLtMaxk += right - left + 1
		}
	}

	result = result - (gteMinkLtMaxk + gtMinkLteMaxk - gtMinkLtMaxk)
	return int64(result)
}
