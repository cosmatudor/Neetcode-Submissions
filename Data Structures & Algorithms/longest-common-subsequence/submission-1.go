func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
    
	memo := make(map[[2]int]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m || j == n {
			return 0
		}

		if val, ok := memo[[2]int{i,j}]; ok {
			return val
		}

		var localRes int
		if text1[i] == text2[j] {
			localRes = 1 + dfs(i+1, j+1)
		} else {
			localRes = max(dfs(i+1, j), dfs(i, j+1))
		}

		memo[[2]int{i,j}] = localRes
		return localRes
	}

	return dfs(0, 0)
}
