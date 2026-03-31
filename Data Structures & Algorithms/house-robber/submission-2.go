func rob(nums []int) int {
	rob1, rob2 := 0, 0 

	for _, x := range nums {
		temp := max(rob1 + x, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}
