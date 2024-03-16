package leetcode

// time complexity is O(n), space complexity is O(n)
func findMaxLength(nums []int) int {
	zeroCount := 0
	oneCount := 0
	max := 0
	diffToIndex := make(map[int]int)
	diffToIndex[0] = -1

	for i, num := range nums {
		if num == 1 {
			oneCount++
		} else {
			zeroCount++
		}

		diff := zeroCount - oneCount

		// the same difference means there is a subarray
		// between the difference first occured index and current index
		if startIndex, ok := diffToIndex[diff]; ok {
			if max < i - startIndex {
				max = i - startIndex
			}
		} else {
			diffToIndex[diff] = i
		}
	}

	return max
}
