func partition(s string) [][]string {
	res := [][]string{}
	sol := []string{}

	isPalindrome := func(partition []string) bool {
		for _, part := range partition {
			if len(part) == 1 {
				continue
			}

			for i := 0; i < len(part)/2; i++ {
				if part[i] != part[len(part)-i-1] {
					return false
				}
			}
		}

		return true
	}

	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			if isPalindrome(sol) {
				res = append(res, append([]string{}, sol...))
			}
		}

		for end := start+1; end <= len(s); end++ {
			sol = append(sol, s[start:end])
			dfs(end)
			sol = sol[:len(sol)-1]
		}
	}

	dfs(0)
	return res
}
