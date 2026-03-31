func uniquePaths(m int, n int) int {
    memo := make(map[[2]int]int)
	memo[[2]int{m-1, n-1}] = 1

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}
		
		if i == m && j == n {
			return 1
		}

		if val, ok := memo[[2]int{i,j}]; ok {
			return val
		}

		return dfs(i+1, j) + dfs(i, j+1)
	}

	return dfs(0,0)
}
