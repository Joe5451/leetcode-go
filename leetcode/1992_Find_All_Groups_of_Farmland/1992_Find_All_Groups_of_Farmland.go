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

func dfs(land [][]int, i, j int) (int, int) {
	r2, c2 := i, j
	for r2 + 1 < len(land) && land[r2 + 1][j] == 1 {
		r2++
	}
	for c2 + 1 < len(land[0]) && land[i][c2 + 1] == 1 {
		c2++
	}

	temp := j
	for i <= r2 {
		j = temp
		for j <= c2 {
			land[i][j] = 0
			j++
		}
		i++
	}

	return r2, c2
}
