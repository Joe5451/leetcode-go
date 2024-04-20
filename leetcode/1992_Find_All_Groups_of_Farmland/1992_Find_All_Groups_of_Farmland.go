package leetcode

// time complexity is O(m * n), space complexity is O(m * n)
func findFarmland(land [][]int) [][]int {
	groups := [][]int{}
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 1 {
				r2, c2 := dfs(land, i, j)
				groups = append(groups, []int{i, j, r2, c2})
				j = c2
			}
		}
	}
	return groups
}

func dfs(land [][]int, row, column int) (int, int) {
	r2, c2 := row, column
	for i := row; i < len(land) && land[i][column] == 1; i++ {
		for j := column; j < len(land[0]) && land[i][j] == 1; j++ {
			land[i][j] = 0
			r2, c2 = i, j
		}
	}
	return r2, c2
}
