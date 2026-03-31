func maxProduct(nums []int) int {
	currMax, currMin := 1, 1
	res := nums[0]
	
	for _, x := range nums {
		currMin *= x

		currMax = max(currMax, 1)
		currMax *= x

		currMax := max(currMax, currMin)
		res = max(res, currMax)
	}

	return res
}
