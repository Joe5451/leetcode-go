package leetcode

// time complexity is O(m*n), space complexity is O(m*n)
func maximalSquare(matrix [][]byte) int {
	maxWidth := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				left, up, diagonal := 0, 0, 0
				if j > 0 {
					up = dp[i][j - 1]
				}
				if i > 0 {
					left = dp[i - 1][j]
				}
				if i > 0 && j > 0 {
					diagonal = dp[i - 1][j - 1]
				}
				squareWidth := min(left, up, diagonal) + 1
				dp[i][j] = squareWidth
			}

			maxWidth = max(maxWidth, dp[i][j])
		}
	}

	return maxWidth * maxWidth
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
