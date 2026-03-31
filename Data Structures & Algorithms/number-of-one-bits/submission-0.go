func hammingWeight(n int) int {
	cnt := 0
	mask := 1
	for i := 0; i < 32; i++ {
		if n & mask != 0 {
			cnt++
		}
		mask = mask << 1
	}
	return cnt
}
