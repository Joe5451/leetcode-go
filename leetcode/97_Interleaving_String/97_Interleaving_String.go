package leetcode

// time complexity is O(m*n), space complexity is O(m*n)
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m + n != len(s3) {
		return false
	}

	dp := make([][]bool, m + 1)
	for i, _ := range dp {
		dp[i] = make([]bool, n + 1)
	}

	dp[0][0] = true
	for i := 0; i < len(s1); i++ {
		if s1[i] == s3[i] {
			dp[i + 1][0] = true
		} else {
			break
		}
	}

	for i := 0; i < len(s2); i++ {
		if s2[i] == s3[i] {
			dp[0][i + 1] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if dp[i - 1][j] && s1[i - 1] == s3[i + j - 1] {
				dp[i][j] = true
			} else if dp[i][j - 1] && s2[j - 1] == s3[i + j - 1] {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[m][n]
}
