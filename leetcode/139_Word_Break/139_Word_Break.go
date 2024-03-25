package leetcode

// time complexity is O(n^2), space complexity is O(n)
func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}

	return dfs(s, dictMap)
}

func dfs(s string, dictMap map[string]bool) bool {
	if result, ok := dictMap[s]; ok {
		return result
	}

	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		prefix := s[:i+1]
		suffix := s[i+1:]
		if dictMap[prefix] {
			if dfs(suffix, dictMap) {
				dictMap[s] = true
				return true
			}
		} else {
			continue
		}
	}

	dictMap[s] = false
	return false
}
