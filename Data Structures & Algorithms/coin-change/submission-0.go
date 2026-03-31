func coinChange(coins []int, amount int) int {
    sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	i := 0
	res := 0
	for amount > 0 && i < len(coins) {
		for amount - coins[i] >= 0 {
			amount -= coins[i]
			res++
		}

		i++
	}

	if amount == 0 {
		return res
	}
	return -1
}
