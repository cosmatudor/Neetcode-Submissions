func maxProfit(prices []int) int {
	memo := make(map[[2]int]int)

    var dfs func(i, j int) int 
	dfs = func(i, j int) int {
		if j >= len(prices) {
			return 0
		}

		if val, ok := memo[[2]int{i,j}]; ok {
			return val
		}

		res := max(prices[j] - prices[i] + dfs(j+2, j+3), dfs(i, j+1))
		memo[[2]int{i,j}] = res

		return res
	}

	res := 0
	for i := 0; i < len(prices); i++ {
		res = max(res, dfs(i, i+1))
	}

	return res
}



// 1 2 3 4

// 12 4      		13(4)
// 				14 