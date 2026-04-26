func generateParenthesis(n int) []string {
	res := []string{}
	currSol := ""

	var dfs func(int, int) 
	dfs = func(open, closed int) {
		if open == closed && open == n {
			res = append(res, currSol)
			return
		}

		if closed > open {
			return 
		}

		if open > n {
			return
		}

			currSol += "("
			dfs(open+1, closed)
			currSol = currSol[:len(currSol)-1]

			currSol += ")"
			dfs(open, closed+1)
			currSol = currSol[:len(currSol)-1]
	}

	dfs(0, 0)

	return res
}
