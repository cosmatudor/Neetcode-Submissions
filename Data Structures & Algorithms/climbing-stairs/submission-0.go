func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

    cache := make([]int, n+1) 

	cache[1] = 1
	cache[2] = 2

	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}
