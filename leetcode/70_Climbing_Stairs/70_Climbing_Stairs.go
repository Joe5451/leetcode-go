package leetcode

// time complexity is O(n), space complexity is O(n)
func climbStairs(n int) int {
	count := 1
	visited := make(map[int]int)
	visited[0] = 1
	return dfs(n, count, visited)
}

func dfs(step, count int, visited map[int]int) int {
	if result, ok := visited[step]; ok {
		return result
	}

	result := count * dfs(step - 1, count, visited)

	if step >= 2 {
		result += count * dfs(step - 2, count, visited)
	}

	visited[step] = result
	return result
}
