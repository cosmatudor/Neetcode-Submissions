func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	res := 0
	memo := make(map[[2]int]int)

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
			if val, exists := memo[[2]int{j, sum + coins[j]}]; exists {
				localRes += val
				continue
			}
			localRes += dfs(j, sum + coins[j])
		}

		memo[[2]int{i, sum}] = localRes
		return localRes
	}

	for i := 0; i < len(coins); i++ {
		res += dfs(i, coins[i])
	}

	return res
}