package leetcode

import "sort"

// time complexity is O(n^2), space complexity is O(1)
func threeSum(nums []int) [][]int {
	size := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i <= size - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := size - 1

		for left < right {
			target := 0 - nums[i]
			if target == nums[left] + nums[right] {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left - 1] {
					left++
				}

				for left < right && nums[right] == nums[right + 1] {
					right--
				}
			} else if nums[left] + nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}

	return ans
}
