func partition(s string) [][]string {
	res := [][]string{}
	sol := []string{}

	isPalindrome := func(partition string) bool {
			if len(partition) == 1 {
				return true
			}

			for i := 0; i < len(partition)/2; i++ {
				if partition[i] != partition[len(partition)-i-1] {
					return false
				}
			}

		return true
	}

	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, sol...))
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s[start:end+1]) {
				sol = append(sol, s[start:end+1])
				dfs(end+1)
				sol = sol[:len(sol)-1]
			}
		}
	}

	dfs(0)
	return res
}
