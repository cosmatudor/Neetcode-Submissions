func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	// shouldContinue := false
    res := []int{}

	k := 0
	for k < (m+1)/2 && k < (n+1)/2 {
		if k < (n+1)/2 && k < m {
			for j := k; j <= n-k-1; j++ {
				res = append(res, matrix[k][j])
			}
		}

		if k < (m+1)/2 && k < n {
			for i := k+1; i <= m-k-1; i++ {
				res = append(res, matrix[i][n-k-1])
			}
		}

		if k < (n+1)/2 && k < m && m-k-1 != k {
			for j := n-k-2; j >= k; j-- {
				res = append(res, matrix[m-k-1][j])
			}
		}

		if k < (m+1)/2 && k < n && n-k-1 != k {
			for i := m-k-2; i >= k+1; i-- {
				res = append(res, matrix[i][k])
			}
		}

		k++
	}

	return res
}
