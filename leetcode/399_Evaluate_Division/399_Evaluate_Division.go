package leetcode

type Vertex struct {
	Name string
	Value float64
}

// time complexity is O((V + E) * q), space complexity is O(V+E)
// q - number of queries
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Vertex)
	for idx, equation := range equations {
		v1, v2 := equation[0], equation[1]
		graph[v1] = append(graph[v1], Vertex{v2, values[idx]})
		graph[v2] = append(graph[v2], Vertex{v1, 1.0 / values[idx]})
	}

	result := make([]float64, len(queries))
	for idx, query := range queries {
		start, end := query[0], query[1]
		visited := make(map[string]bool)
		result[idx], _ = dfs(start, end, graph, visited)
	}

	return result
}

func dfs(start, end string, graph map[string][]Vertex, visited map[string]bool) (float64, bool) {
	if _, ok := graph[start]; !ok {
		return -1.0, false
	}
	if _, ok := graph[end]; !ok {
		return -1.0, false
	}
	if start == end {
		return 1.0, true
	}
	for _, neighbor := range graph[start] {
		if visited[neighbor.Name] {
			continue
		}
		visited[neighbor.Name] = true
		if value, ok := dfs(neighbor.Name, end, graph, visited); ok {
			return neighbor.Value * value, true
		}
	}

	return -1.0, false
}
