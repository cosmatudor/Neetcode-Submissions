func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	sol := ""

	var dfs func(i int)
	dfs = func(i int) {
		if len(sol) == len(digits) {
			res = append(res, sol)
			return
		}

		for _, x := range digitToLetters[digits[i]] {
			sol += string(x)
			dfs(i+1)
			sol = sol[:len(sol)-1]
		}
	}

	dfs(0)

	return res
}
