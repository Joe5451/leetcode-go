package leetcode

// time complexity is O(n), space complexity is O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	i, j := size, size
	for k := 0; k < size; k++ {
		if newInterval[0] <= intervals[k][1] && i == size {
			i = k
		}

		if newInterval[1] <= intervals[k][1] && j == size {
			j = k
		}

		if i != size && j != size {
			break
		}
	}

	if i != size && newInterval[0] > intervals[i][0] {
		newInterval[0] = intervals[i][0]
	}

	if j != size && newInterval[1] >= intervals[j][0] {
		newInterval[1] = intervals[j][1]
		j++
	}

	return append(append(append([][]int{}, intervals[:i]...), newInterval), intervals[j:]...)
}
