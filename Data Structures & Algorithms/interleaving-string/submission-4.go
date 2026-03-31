func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)

	if m+n != len(s3) {
		return false
	}

	memo := make(map[[2]int]bool)

	var dfs func(i, j int, curr string) bool
	dfs = func(i, j int, curr string) bool {
		if i == m && j == n {
			return true
		}

		if val, ok := memo[[2]int{i, j}]; ok {
			return val
		}

		res := false
		if i < m && s1[i] == s3[i+j] {
			res = res || dfs(i+1, j, curr + string(s1[i]))
		}
		if j < n && s2[j] == s3[i+j] {
			res = res || dfs(i, j+1, curr + string(s2[j]))
		}

		memo[[2]int{i, j}] = res
		return res
	}

	return dfs(0, 0, "")
}