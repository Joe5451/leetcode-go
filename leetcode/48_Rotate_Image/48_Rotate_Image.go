package leetcode

// time complexity is O(n^2), space complexity is O(1)
func rotate(matrix [][]int)  {
	n := len(matrix)
	var tempNum, currentNum, rotatedRow, rotatedColumn int

	for i := 0; i < n/2; i++ {
		for j := i; j < n - 1 - i; j++ {
			currentRow := i
			currentColumn := j
			currentNum = matrix[currentRow][currentColumn]
			for count := 0; count < 4; count++ {
				rotatedRow = currentColumn
				rotatedColumn = n - 1 - currentRow
				tempNum = matrix[rotatedRow][rotatedColumn]
				matrix[rotatedRow][rotatedColumn] = currentNum

				currentNum = tempNum
				currentRow = rotatedRow
				currentColumn = rotatedColumn
			}
		}
	}
}


// time complexity is O(n^2), space complexity is O(n^2)
func rotate2(matrix [][]int)  {
	n := len(matrix)
	nums := make([]int, n*n)
	for i, row := range matrix {
		for j, num := range row {
			nums[n*j + n - 1 - i] = num
		}
	}

	for i, num := range nums {
		matrix[i/n][i%n] = num
	}
}
