package leetcode

// time complexity is O(n), space complexity is O(n)
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = -1
	result := 0
	if i - 1 < 0 || grid[i - 1][j] == 0 {
		result++
	}
	if i + 1 >= len(grid) || grid[i + 1][j] == 0 {
		result++
	}
	if j - 1 < 0 || grid[i][j - 1] == 0 {
		result++
	}
	if j + 1 >= len(grid[0]) || grid[i][j + 1] == 0 {
		result++
	}

	return result + dfs(grid, i - 1, j) + dfs(grid, i + 1, j) + dfs(grid, i, j - 1) + dfs(grid, i, j + 1)
}
