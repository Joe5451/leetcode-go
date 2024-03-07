package leetcode

func removeDuplicates(nums []int) int {
	max := -101
	sortedIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			nums[sortedIndex] = nums[i]
			sortedIndex++
		}
	}

	return sortedIndex
}
