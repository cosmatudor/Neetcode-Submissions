func maxProfit(prices []int) int {
    var dfs func(i, j int) int 
	dfs = func(i, j int) int {
		if j >= len(prices) {
			return 0
		}

		return max(prices[j] - prices[i] + dfs(j+2, j+3), dfs(i, j+1))
	}

	return dfs(0, 1)
}



// 1 2 3 4

// 12 4      		13(4)
// 				14 