func findTargetSumWays(nums []int, target int) int {
	memo := make(map[[2]int]int)

	var dfs func(i, sum int) int 
	dfs = func(i, sum int) int {
		if i == len(nums) - 1 {
			if sum + nums[i] == target || sum - nums[i] == target {
				if nums[i] == 0 {
					return 2
				}
				return 1
			}
			return 0
		}

		res := 0
		if negVal, ok1 := memo[[2]int{i+1, sum-nums[i]}]; ok1 {
			res += negVal
		} else {
			res += dfs(i + 1, sum - nums[i])
		}

		if posVal, ok2 := memo[[2]int{i+1, sum+nums[i]}]; ok2 {
			res += posVal
		} else {
			res += dfs(i + 1, sum + nums[i])
		}

		memo[[2]int{i, sum}] = res
		return res
	}

	return dfs(0, 0)
}
