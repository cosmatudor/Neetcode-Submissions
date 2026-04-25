func canJump(nums []int) bool {
    if len(nums) == 1 {
		return true
	}

	dst := len(nums)-1
	for src := dst-1; src >= 0; src-- {
		if dst - src <= nums[src] {
			dst = src
		}
	}

	return dst == 0
}
