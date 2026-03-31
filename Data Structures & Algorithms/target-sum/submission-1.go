func findTargetSumWays(nums []int, target int) int {
	// memo := make(map[[2]int]int)

	var dfs func(i, sum int) int 
	dfs = func(i, sum int) int {
		if i == len(nums) - 1 {
			if sum + nums[i] == target || sum - nums[i] == target {
				return 1
			}
			return 0
		}

		res := dfs(i + 1, sum + nums[i]) + dfs(i + 1, sum - nums[i])
		return res
	}

	return dfs(0, 0)
}
