package leetcode

// time complexity is O(n), space complexity is O(n)
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	strSet := make(map[byte]bool)
	max := 1
	strSet[s[0]] = true
	for i, j := 0, 0; j < len(s); j++ {
		if i == j {
			continue
		}

		if strSet[s[j]] {
			for s[i] != s[j] {
				strSet[s[i]] = false
				i++
			}

			i++
		} else {
			length := j - i + 1
			if length > max {
				max = length
			}

			strSet[s[j]] = true
		}
	}

	return max
}
