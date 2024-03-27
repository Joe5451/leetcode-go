package leetcode

import "math"

// Bottom-up DP
// time complexity is O(n), space complexity is O(n)
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i - 1][0]
			} else if j == len(triangle[i]) - 1 {
				dp[i][j] = triangle[i][j] + dp[i - 1][j - 1]
			} else {
				dp[i][j] = triangle[i][j] + min(dp[i - 1][j], dp[i - 1][j - 1])
			}
		}
	}

	min := math.MaxInt64
	for _, num := range dp[len(triangle) - 1] {
		if num < min {
			min = num
		}
	}

	return min
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

// Top-down DP
// time complexity is O(n), space complexity is O(n)
func minimumTotal2(triangle [][]int) int {
	visited := make(map[[2]int]int)
	return dfs(0, 0, triangle, visited)
}

func dfs(row, column int, triangle [][]int, visited map[[2]int]int) int {
	if result, ok := visited[[2]int{row, column}]; ok {
		return result
	}

	if row >= len(triangle) {
		return 0
	}

	result := triangle[row][column] + min(dfs(row + 1, column, triangle, visited), dfs(row + 1, column + 1, triangle, visited))
	visited[[2]int{row, column}] = result
	return result
}
