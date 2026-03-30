func permute(nums []int) [][]int {
	set := make(map[int]struct{})
	res := [][]int{}
	currSol := []int{}

	var dfs func()
	dfs = func() {
		if len(set) == len(nums) {
			temp := make([]int, len(currSol))
			copy(temp, currSol)
			res = append(res, temp)
			return
		}

		for j := 0; j < len(nums); j++ {
			if _, ok := set[j]; ok {
				continue
			}
			
			currSol = append(currSol, nums[j])
			set[j] = struct{}{}

			dfs()

			currSol = currSol[:len(currSol)-1]
			delete(set, j)
		}
	}

	dfs()
	return res
}
