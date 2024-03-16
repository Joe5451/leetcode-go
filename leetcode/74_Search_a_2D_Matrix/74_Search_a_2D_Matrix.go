package leetcode

// time complexity is O(log(m * n)), space complexity is O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m - 1][n - 1] {
		return false
	}

	// time complexity is O(log m)
	targetRow := findRow(matrix, target)

	// time complexity is O(log n)
	return binarySearch(targetRow, target)
}

func findRow(matrix [][]int, target int) []int {
	targetRowIndex := -1

	for targetRowIndex == -1 {
		mid := len(matrix) / 2

		if matrix[mid][0] <= target {
			if mid + 1 >= len(matrix) ||
				(mid + 1 < len(matrix) && matrix[mid + 1][0] > target) {
				targetRowIndex = mid
				break
			}
		} else {
			if mid - 1 < 0 {
				targetRowIndex = mid
				break
			} else if mid - 1 >= 0 && matrix[mid - 1][0] <= target {
				targetRowIndex = mid - 1
				break
			}
		}

		if matrix[mid][0] < target {
			matrix = matrix[mid + 1:]
		} else {
			matrix = matrix[:mid - 1]
		}
	}

	return matrix[targetRowIndex]
}

func binarySearch(row []int, target int) bool {
	if len(row) == 1 {
		return row[0] == target
	}

	if len(row) < 1 {
		return false
	}

	mid := len(row) / 2
	if row[mid] == target {
		return true
	} else if row[mid] < target {
		return binarySearch(row[mid + 1:], target)
	} else {
		return binarySearch(row[:mid], target)
	}
}
