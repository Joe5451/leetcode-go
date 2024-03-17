package leetcode

import "sort"

// time complexity is O(nlogn + n) -> O(nlogn), space complexity is O(n)
func merge(intervals [][]int) [][]int {
	// time complexity: O(nlogn)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// time complexity: O(n)
	results := [][]int{intervals[0]}
	for i, j := 0, 1; j < len(intervals); j++ {
		if results[i][1] >= intervals[j][0] {
			results[i][1] = max(results[i][1], intervals[j][1])
		} else {
			results = append(results, intervals[j])
			i++
		}
	}

	return results
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
