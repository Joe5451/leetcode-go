package leetcode

// time complexity is O(n), space complexity is O(1)
func countStudents(students []int, sandwiches []int) int {
	preferCount := [2]int{}
	for _, prefer := range students {
		preferCount[prefer]++
	}

	for i, sandwich := range sandwiches {
		if preferCount[sandwich] <= 0 {
			return len(sandwiches) - i
		}

		preferCount[sandwich]--
	}
	return 0
}
