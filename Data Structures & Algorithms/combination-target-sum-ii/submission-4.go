func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sol := []int{}
	sum := 0

	sort.Ints(candidates)

	var dfs func(i int)
	dfs = func(i int) {
		if sum == target {
			temp := make([]int, len(sol))
			copy(temp, sol)
			res = append(res, temp)
			return
		}
		if i == len(candidates) {
			return
		}
		if sum > target {
			return
		}

		sol = append(sol, candidates[i])
		sum += candidates[i]
		dfs(i+1)

		sol = sol[:len(sol)-1]
		sum -= candidates[i]

		for i < len(candidates) - 1 && candidates[i] == candidates[i+1] {
			i++
		}

		dfs(i+1)
	}

	dfs(0)

	return res
}
