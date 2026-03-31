func maxSubArray(nums []int) int {
    maxSum := nums[0]
	currSum := 0

	for _, x := range nums {
		// choose whether add previous subarray or not
		// in this context of max sum, this is done by checking if 
		// the sum of that subarray is > 0 or not
		currSum = max(currSum, 0)

		// then add the current element
		currSum += x

		// then, for every step, check if that sum is the biggest one
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
