package leetcode

// this solution is better
// time complexity is O(n), space complexity is O(1)
func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || i == 1 || nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

// time complexity is O(n), space complexity is O(1)
func removeDuplicates2(nums []int) int {
	currentVal := -10001
	currentIndex := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == currentVal && count < 2 {
			nums[currentIndex] = nums[i]
			count++
			currentIndex++
		}

		if nums[i] > currentVal {
			currentVal = nums[i]
			count = 1
			nums[currentIndex] = nums[i]
			currentIndex++
		}
	}

	return currentIndex
}
