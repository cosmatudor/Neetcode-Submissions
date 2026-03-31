func missingNumber(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}

	return nums[len(nums)-1]*(nums[len(nums)-1]+1)/2 - sum
}
