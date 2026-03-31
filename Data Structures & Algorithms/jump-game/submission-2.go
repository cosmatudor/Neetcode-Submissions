func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	n := len(nums)
    goal := n - 1
	it := n - 2

	for it >= 0 {
		for it >= 0 && nums[it] == 0 {
			it--
		}

		if it >= 0 && nums[it] >= goal - it {
			goal = it
			if goal == 0 {
				return true
			}
		}

		it--
	}

	return false

}
