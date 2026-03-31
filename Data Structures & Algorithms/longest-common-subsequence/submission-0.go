func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
    dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			curr := 0
			if text1[i-1] == text2[j-1] {
				curr = 1
			}

			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + curr
		}
	}

	return dp[m][n]
}
