package leetcode

// time complexity: O(V), space complexity: O(V + E)
func validPath(n int, edges [][]int, source int, destination int) bool {
	// time complexity: O(V), space complexity: O(V + E)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	// time complexity: O(V + E), space complexity: O(V)
	visited := make(map[int]bool)
	queue := []int{source}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == destination {
			return true
		} else if !visited[current] {
			visited[current] = true
			queue = append(queue, graph[current]...)
		}
	}
	return false
}
