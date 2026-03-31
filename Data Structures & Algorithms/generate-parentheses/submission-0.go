func generateParenthesis(n int) []string {
	res := []string{}
	sol := ""

	var dfs func(int, int) 
	dfs = func(open, closed int) {
		if open == n && closed == n {
			res = append(res, sol) 
			return
		}

		if open < n {
			sol += "("
			dfs(open+1, closed)
			sol = sol[:len(sol)-1]
		}

		if closed < open {
			sol += ")"
			dfs(open, closed+1)
			sol = sol[:len(sol)-1]
		}
	}

	dfs(0, 0)
	return res
}
