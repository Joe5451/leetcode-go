package leetcode

// time complexity: O(n^2), space complexity is O(1)
func minFallingPathSum(grid [][]int) int {
	if len(grid) <= 1 {
		return grid[0][0]
	}
	previousFirstMinIndex, previousSecondMinIndex := -1, -1
	for i, j := 0, 0; j < len(grid[i]); j++ {
		if previousFirstMinIndex == -1 || grid[i][j] <= grid[i][previousFirstMinIndex] {
			previousFirstMinIndex, previousSecondMinIndex = j, previousFirstMinIndex
		} else if previousSecondMinIndex == -1 || grid[i][j] < grid[i][previousSecondMinIndex] {
			previousSecondMinIndex = j
		}
	}

	for i := 1; i < len(grid); i++ {
		currentFirstMinIndex, currentSecondMinIndex := -1, -1
		for j := 0; j < len(grid); j++ {
			if j != previousFirstMinIndex {
				grid[i][j] = grid[i][j] + grid[i - 1][previousFirstMinIndex]
			} else {
				grid[i][j] = grid[i][j] + grid[i - 1][previousSecondMinIndex]
			}

			if currentFirstMinIndex == -1 || grid[i][j] <= grid[i][currentFirstMinIndex] {
				currentFirstMinIndex, currentSecondMinIndex = j, currentFirstMinIndex
			} else if currentSecondMinIndex == -1 || grid[i][j] < grid[i][currentSecondMinIndex] {
				currentSecondMinIndex = j
			}
		}
		previousFirstMinIndex, previousSecondMinIndex = currentFirstMinIndex, currentSecondMinIndex
	}

	return grid[len(grid) - 1][previousFirstMinIndex]
}
