func maxSubArray(nums []int) int {
	res := -10001
	sum := 0
	for _, x := range nums {
		sum = max(0, sum)
		sum += x
		res = max(res, sum)
	}

	return res
}
