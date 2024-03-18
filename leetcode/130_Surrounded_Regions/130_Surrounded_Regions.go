package leetcode

// time complexity is O(m*n), space complexity is O(m*n)
func solve(board [][]byte)  {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m - 1 || j == 0 || j == n - 1 {
			   dfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'R' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, column int) {
	if row < 0 || row >= len(board) ||
		column < 0 || column >= len(board[0]) ||
		board[row][column] == 'X' || board[row][column] == 'R' {
		return
	}

	board[row][column] = 'R'
	dfs(board, row + 1, column)
	dfs(board, row - 1, column)
	dfs(board, row, column + 1)
	dfs(board, row, column - 1)
}
