package leetcode

import "sort"

// time complexity is O(nlogn + n) -> O(nlogn), space complexity is O(1)
func findMinArrowShots(points [][]int) int {
	// time complexity: O(nlogn)
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrowPosition := points[0][1]
	count := 1
	// time complexity: O(n)
	for i := 0; i < len(points) - 1; i++ {
		if arrowPosition >= points[i + 1][0] && arrowPosition <= points[i + 1][1] {
			continue
		} else {
			arrowPosition = points[i + 1][1]
			count++
		}
	}

	return count
}
