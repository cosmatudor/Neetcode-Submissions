func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	currSol := []int{}
	res := [][]int{}

	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if sum == target {
			temp := make([]int, len(currSol))
			copy(temp, currSol)
			res = append(res, temp)
			return
		}

		if sum > target || i >= len(candidates) {
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			currSol = append(currSol, candidates[j])
			dfs(j+1, sum+candidates[j])
			currSol = currSol[:len(currSol)-1]
		}
	}

	dfs(0, 0)

	return res
}