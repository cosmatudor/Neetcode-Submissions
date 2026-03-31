func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if val, ok := dp[amount]; ok {
			return val
		}

		localMin := math.MaxInt32
		for _, coin := range coins {
			if amount - coin >= 0 {
				localRes := dfs(amount - coin)
				if localRes < localMin {
					localMin = localRes
				}
			}
		}

		dp[amount] = 1 + localMin

		return dp[amount]
	}

	res := dfs(amount)
	if res == amount + 1 {
		return -1
	}
	return res
}
