func wordBreak(s string, wordDict []string) bool {
    memo := make(map[int]bool) // s[i:] can be split up

	var dfs func(i int) bool 
	dfs = func(i int) bool {
		if i >= len(s) {
			return true
		}

		if val, exists := memo[i]; exists {
			return val
		}

		for _, word := range wordDict {
			if i + len(word) > len(s) || s[i: i+len(word)] != word {
				continue
			}

			if dfs(i + len(word)) {
				memo[i] = true
				return true
			}
		}

		memo[i] = false
		return false
	}

	return dfs(0)
}

// Input: s = "cat sin cars", wordDict = ["cat","in","cats","cars"]
//
// i >= len(s) -> true
// s[i:i+len(w1)] == w1 && dfs(i+len(w1)) || ... || s[i:i+len(wn)] == wn && dfs(i+len(wn))

