package leetcode

// time complexity is O(m*n), space complexity is O(m*n)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i - 1][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i - 1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])
		}
	}

	return dp[m - 1][n - 1]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}