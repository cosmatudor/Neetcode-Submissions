func change(amount int, coins []int) int {
	res := 0

	var dfs func(i, sum int) int 
	dfs = func(i, sum int) int {
		if sum > amount {
			return 0
		}

		if sum == amount {
			return 1
		}

		localRes := 0
		for j := i; j < len(coins); j++ {
			localRes += dfs(j, sum + coins[j])
		}

		return localRes
	}

	for i := 0; i < len(coins); i++ {
		res += dfs(i, coins[i])
	}

	return res
}