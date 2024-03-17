package leetcode

// time complexity is O(n), space complexity is O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	insertStartIndex := -1
	insertEndIndex := -1

	for i, interval := range intervals {
		if newInterval[0] <= interval[1] && insertStartIndex == -1 {
			insertStartIndex = i
		}

		if newInterval[1] <= interval[1] && insertEndIndex == -1 {
			insertEndIndex = i
		}
	}

	if insertStartIndex != -1 {
		newInterval[0] = min(newInterval[0], intervals[insertStartIndex][0])
	} else {
		insertStartIndex = len(intervals)
	}

	if insertEndIndex != -1 {
		if newInterval[1] < intervals[insertEndIndex][0] {
			insertEndIndex -= 1
		} else {
			newInterval[1] = max(newInterval[1], intervals[insertEndIndex][1])
		}
	} else {
		insertEndIndex = len(intervals) - 1
	}

	prefix := intervals[:insertStartIndex]
	suffix := intervals[insertEndIndex + 1:]
	return append(append(append([][]int{}, prefix...), newInterval), suffix...)
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}

	return num1
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}


func insert2(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	i := size
	j := size
	for k := 0; k < size; k++ {
		if i == size && intervals[k][0] > newInterval[0] {
			i = k
		}

		if j == size && intervals[k][1] > newInterval[1] {
			j = k
		}

		if i != size && j != size {
			break
		}
	}

	if i >= 1 && intervals[i - 1][1] >= newInterval[0] {
		newInterval[0] = intervals[i-1][0]
		i--
	}

	 if j < size && intervals[j][0] <= newInterval[1] {
		newInterval[1] = intervals[j][1]
		j++
	}

	return append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
}
