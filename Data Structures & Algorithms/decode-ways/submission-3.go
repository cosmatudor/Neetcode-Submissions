func numDecodings(s string) int {
    dp := make(map[int]int)

	var dfs func(i int) int
	dfs = func(i int) int {
		// when we reached i = len(s), we have a solution
		if i == len(s) {
			return 1
		}

		if s[i] == '0' {
			return 0
		}

		if val, ok := dp[i]; ok {
			return val
		}

		// we can choose 1 or 2 digits
		dp[i] = dfs(i+1)
		if i < len(s)-1 && (s[i] == '1' || (s[i] == '2' && s[i+1] < '7')) {
			dp[i] += dfs(i+2)
		}

		return dp[i]
	}

	return dfs(0)
}
