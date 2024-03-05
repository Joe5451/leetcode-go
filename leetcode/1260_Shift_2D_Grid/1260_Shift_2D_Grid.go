package leetcode

// time complexity is O(m*n), space complexity is O(m*n)
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	shiftedGrid := make([][]int, m)
	for rowIndex, _ := range shiftedGrid {
		shiftedGrid[rowIndex] = make([]int, n)
	}

	var shiftedRowIndex, shiftedColumnIndex int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			shiftedRowIndex = (i + (j + k) / n) % m
			shiftedColumnIndex = (j + k) % n
			shiftedGrid[shiftedRowIndex][shiftedColumnIndex] = grid[i][j]
		}
	}

	return shiftedGrid
}

// use 1D array
// time complexity is O(m*n), space complexity is O(m*n)
func shiftGrid2(grid [][]int, k int) [][]int {
	totalAmount := len(grid) * len(grid[0])
	k = k % totalAmount
	unshiftedArr := make([]int, 0)
	for _, row := range grid {
		for _, element := range row {
			unshiftedArr = append(unshiftedArr, element)
		}
	}

	shiftedArr := make([]int, 0)
	shiftedArr = append(unshiftedArr[totalAmount - k:], unshiftedArr[:totalAmount - k]...)
	pos := 0
	for rowIndex, row := range grid {
		for columnIndex, _ := range row {
			grid[rowIndex][columnIndex] = shiftedArr[pos]
			pos++
		}
	}

	return grid
}
