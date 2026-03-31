func computeNextInterval(l, r int, nums []int) (int, int) {
	farthestPoint := r + 1
	for i := l; i <= r; i++ {
		farthestPoint = max(farthestPoint, i+nums[i])
	}

	return r + 1, farthestPoint
}

func jump(nums []int) int {
    r, l := 0, 0
	n := len(nums)
	cnt := 0

	for {
		cnt++
		l, r = computeNextInterval(l, r, nums)
		if r >= n - 1 {
			break
		}
	}

	return cnt
}
