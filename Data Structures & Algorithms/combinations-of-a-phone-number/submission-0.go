func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	
	digitToLetters := map[rune][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}

	set := make(map[string]struct{})
	res := []string{}

	var dfs func(idxs []int)
	dfs = func(idxs []int) {
		for i, idx := range idxs {
			if idx >= len(digitToLetters[rune(digits[i])]) {
				return
			}
		}

		sol := ""
		for i, idx := range idxs {
			sol += string(digitToLetters[rune(digits[i])][idx])
		}
		set[sol] = struct{}{}

		for i := range idxs {
			temp := append([]int{}, idxs...)
			temp[i]++
			dfs(temp)
		}
	}

	idxs := make([]int, len(digits))
	dfs(idxs)

	for val, _ := range set {
		res = append(res, val)
	}

	return res
}
