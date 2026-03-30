func permute(nums []int) [][]int {
	set := make(map[int]struct{})
	res := [][]int{}
	currSol := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if _, ok := set[i]; ok {
			return
		}

		currSol = append(currSol, nums[i])
		set[i] = struct{}{}

		if len(set) == len(nums) {
			temp := make([]int, len(currSol))
			copy(temp, currSol)
			res = append(res, temp)
		} else {
			for j := 0; j < len(nums); j++ {
				if _, ok := set[j]; !ok {
					dfs(j)
				}
			}
		}

		currSol = currSol[:len(currSol)-1]
		delete(set, i)
	}

	for i := 0; i < len(nums); i++ {
		dfs(i)
	}
	return res
}