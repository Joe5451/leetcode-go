package leetcode

// time complexity is O(logn + n) -> O(logn), space complexity is O(1)
func search(nums []int, target int) int {
	pivot := 0
	// time complexity: O(n)
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {
			pivot = i + 1
		}
	}

	// time complexity: O(logn)
	if target <= nums[len(nums) - 1] {
		return binarySearch(nums, pivot, len(nums) - 1, target)
	} else {
		return binarySearch(nums, 0, pivot - 1, target)
	}
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
