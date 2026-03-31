func lengthOfLIS(nums []int) int {
    memo := make(map[int]int)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == len(nums) {
			return 0
		}

		if val, exists := memo[i]; exists {
			return val
		}

		LIS := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= nums[i] {
				continue
			}

			LIS = max(LIS, 1+dfs(j))
		}

		memo[i] = LIS
		return LIS
	}

	maxLIS := 1
    for i := 0; i < len(nums); i++ {
        maxLIS = max(maxLIS, dfs(i))
    }

    return maxLIS
}
