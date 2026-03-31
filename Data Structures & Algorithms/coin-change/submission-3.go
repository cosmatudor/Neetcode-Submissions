func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if val, ok := dp[amount]; ok {
			return val
		}

		res := math.MaxInt32
		for _, coin := range coins {
			if amount - coin >= 0 {
				res = min(res, 1 + dfs(amount - coin))
			}
		}

		dp[amount] = res

		return dp[amount]
	}

	res := dfs(amount)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
