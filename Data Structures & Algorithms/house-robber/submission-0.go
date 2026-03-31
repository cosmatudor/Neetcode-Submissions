func rob(nums []int) int {
	if len(nums) <= 2 {
		return nums[len(nums)-1]
	}

    for i := 2; i < len(nums); i++ {
		nums[i] += nums[i-2]
	}

	return max(nums[len(nums)-1], nums[len(nums)-2])
}
