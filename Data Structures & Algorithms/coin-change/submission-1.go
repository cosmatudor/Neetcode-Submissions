func coinChange(coins []int, amount int) int {
    dp := make(map[int]int)

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount < 0 {
			return math.MaxInt32
		}

		if amount == 0 {
			return 0
		}

		if val, ok := dp[amount]; ok {
			return val
		}

		localMin := math.MaxInt32
		for _, coin := range coins {
			localRes := dfs(amount - coin)
			if localRes < localMin {
				localMin = localRes
			}
		}

		if localMin == math.MaxInt32 {
			dp[amount] = math.MaxInt32
		} else {
			dp[amount] = 1 + localMin
		}

		return dp[amount]
	}

	res := dfs(amount)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
