func rob(nums []int) int {
    n := len(nums)

	if n == 1 {
		return nums[0]
	}

	baseRob := func(start, end int) int {
		if end - start == 1 {
			return nums[start]
		}

		dp := make([]int, end+1)
		dp[start] = nums[start]
		dp[start+1] = max(nums[start], nums[start+1])

		for i := start+2; i < end; i++ {
			dp[i] = max(dp[i-2] + nums[i], dp[i-1])
		}

		return dp[end-1]
	}

	return max(baseRob(0, n-1), baseRob(1, n))
}
