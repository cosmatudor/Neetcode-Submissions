func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if len(word1[i:]) == 0 && len(word2[j:]) == 0 {
			return 0
		} 
		if len(word1[i:]) == 0 {
			return len(word2[j:])
		}
		if len(word2[j:]) == 0 {
			return len(word1[i:])
		}

		if i < m && j < n && word1[i] == word2[j] {
			return dfs(i+1, j+1)
		} 

		return 1 + min(min(dfs(i, j+1), dfs(i+1, j)), dfs(i+1, j+1))
	}
	
	return dfs(0, 0)
}
