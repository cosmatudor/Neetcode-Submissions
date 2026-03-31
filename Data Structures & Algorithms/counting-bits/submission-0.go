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

func countBits(n int) []int {
	arr := []int{}
	for i := 0; i <= n; i++ {
		arr = append(arr, hammingWeight(i))
	}

	return arr
}
