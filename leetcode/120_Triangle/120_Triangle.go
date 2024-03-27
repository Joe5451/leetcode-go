package leetcode

import "math"

// Bottom-up DP
// time complexity is O(n), space complexity is O(n)
func minimumTotal(triangle [][]int) int {
	visited := make([][]int, len(triangle))
	visited[0] = make([]int, 1)
	visited[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		visited[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				visited[i][j] = triangle[i][j] + visited[i - 1][0]
			} else if j == len(triangle[i]) - 1 {
				visited[i][j] = triangle[i][j] + visited[i - 1][j - 1]
			} else {
				visited[i][j] = triangle[i][j] + min(visited[i - 1][j], visited[i - 1][j - 1])
			}
		}
	}

	min := math.MaxInt64
	for _, num := range visited[len(triangle) - 1] {
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
