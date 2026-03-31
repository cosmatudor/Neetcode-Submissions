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

		currSol = append(currSol, candidates[i])
		dfs(i+1, sum+candidates[i])
		currSol = currSol[:len(currSol)-1]
		
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
            i++
        }
		dfs(i+1, sum)
	}

	dfs(0, 0)

	return res
}