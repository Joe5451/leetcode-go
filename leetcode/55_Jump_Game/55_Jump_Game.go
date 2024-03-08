package leetcode

// time complexity is O(n), space complexity is O(1)
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	i := len(nums) - 2
	currentTarget := len(nums) - 1

	for i >= 0 {
		for i >= 0 && nums[i] < currentTarget - i {
			i--
		}

		if i == 0 {
			return true
		}

		currentTarget = i
		i = i - 1
	}

	return false
}

// time complexity is O(n), space complexity is O(1)
func canJump2(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i + nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}
