package leetcode

// time complexity is O(n), space complexity is O(1)
func maxFrequencyElements(nums []int) int {
	frequencies := map[int]int{}
	maxFrequency := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		frequencies[nums[i]]++
		frequency := frequencies[nums[i]]

		if frequency > maxFrequency {
			maxFrequency = frequency
			count = frequency
		} else if frequency == maxFrequency {
			count += frequency
		}
	}

	return count
}
