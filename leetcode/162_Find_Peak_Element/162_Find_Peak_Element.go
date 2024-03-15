package leetcode

// time complexity is O(log n), space complexity is O(1)
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > nums[mid + 1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// time complexity is O(log n), space complexity is O(log n)
func findPeakElement2(nums []int) int {
	return partition(nums, 0, len(nums) - 1)
}

func partition(nums []int, low, high int) int {
	if low == high {
		return low
	}

	if high - low == 1 {
		if nums[high] > nums[low] {
			return high
		} else {
			return low
		}
	}

	mid := (low + high) / 2
	if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1]  {
		return mid
	}

	leftPeekIndex := partition(nums, low, mid)
	rightPeekIndex := partition(nums, mid + 1, high)
	if nums[leftPeekIndex] > nums[rightPeekIndex] {
		return leftPeekIndex
	} else {
		return rightPeekIndex
	}
}
