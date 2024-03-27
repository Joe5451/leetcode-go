package leetcode

// time complexity is O(m*n), space complexity is O(1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 0 || (i == 0 && j == 0) {
				continue
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j - 1]
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i - 1][j]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i - 1][j] + obstacleGrid[i][j - 1]
			}
		}
	}

	return obstacleGrid[m - 1][n - 1]
}
