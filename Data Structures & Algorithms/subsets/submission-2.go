func subsets(nums []int) [][]int {
	res := [][]int{}
	sol := []int{}

	var dfs func(i int) 
	dfs = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(sol))
			copy(temp, sol)
			res = append(res, temp)
			return
		}

		sol = append(sol, nums[i])
		dfs(i+1)
		sol = sol[:len(sol)-1]
		dfs(i+1)
	}
	
	dfs(0)
	return res
}


	// 	1     -
	//    2 -   2 -
	//   3 - 