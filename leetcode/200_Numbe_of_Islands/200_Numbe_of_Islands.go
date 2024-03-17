package leetcode

// time complexity is O(m*n + 4*m*n) -> O(m*n), space complexity is O(m*n)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if grid[row][column] == '1' {
				count++
				dfs(grid, row, column)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row, column int) {
	if row < 0 || row >= len(grid) ||
		column < 0 || column >= len(grid[0]) ||
		grid[row][column] != '1' {
		return
	}

	grid[row][column] = '2'
	dfs(grid, row - 1, column)
	dfs(grid, row + 1, column)
	dfs(grid, row, column - 1)
	dfs(grid, row, column + 1)
}
