func missingNumber(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}

	return (len(nums)+1)*len(nums)/2 - sum
}
