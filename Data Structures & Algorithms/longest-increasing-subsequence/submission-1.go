func lengthOfLIS(nums []int) int {
	n := len(nums)
    dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n ; j++ {
			if nums[i] >= nums[j] {
				continue
			}

			dp[i] = max(dp[i], 1 + dp[j])
		}
	}

	res := math.MinInt32
	for _, lenght := range dp {
		res = max(res, lenght)
	}
	return res
}
