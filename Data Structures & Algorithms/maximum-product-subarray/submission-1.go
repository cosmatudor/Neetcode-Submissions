func maxProduct(nums []int) int {
	currMax, currMin := 1, 1
	res := nums[0]
	
	for _, x := range nums {
		tmp := currMax
		currMin = min(currMin, 1)
		currMin *= x

		currMax = max(currMax, 1)
		currMax *= x

		currMax = max(currMax, currMin)
		currMin = min(tmp * x, currMin)

		res = max(res, currMax)
	}

	return res
}
