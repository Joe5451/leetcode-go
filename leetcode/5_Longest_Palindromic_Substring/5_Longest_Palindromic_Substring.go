package leetcode

// time complexity is O(n), space complexity is O(1)
func longestPalindrome(s string) string {
	maxLen, start, end := 1, 0, 0
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		for right + 1 < len(s) && s[left] == s[right + 1] {
			right++
		}

		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		if right - left - 1 > maxLen {
			maxLen, start, end = right - left - 1, left + 1, right - 1
		}
	}

	return s[start:end + 1]
}

// DP
// time complexity is O(n^2), space complexity is O(n^2)
func longestPalindrome2(s string) string {
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	max, start, end := 1, 0, 0
	for i := len(s) - 2; i >= 0; i-- {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] && (i + 1 == j || dp[i + 1][j -1]) {
				dp[i][j] = true
				if j - i + 1 > max {
					max = j - i + 1
					start = i
					end = j
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return s[start:end + 1]
}
