package leetcode

// time complexity is O(n), space complexity is O(n)
func maxSubarrayLength(nums []int, k int) int {
	frequencyMap := make(map[int]int)

	max := 0
	for i, j := 0, 0; j < len(nums); j++ {
		frequencyMap[nums[j]]++
		if frequencyMap[nums[j]] > k {
			for nums[i] != nums[j] {
				frequencyMap[nums[i]]--
				i++
			}
			frequencyMap[nums[i]]--
			i++
			continue
		}

		if j - i + 1 > max {
			max = j - i + 1
		}
	}

	return max
}
