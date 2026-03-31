func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	shouldContinue := false
    res := []int{}

	fmt.Println(n)
	fmt.Println(m)

	k := 0
	for {
		fmt.Println("iteraton ---", k)
		if k <= n/2 {
			for j := k; j <= n-k-1; j++ {
				res = append(res, matrix[k][j])
				fmt.Println(res)
			}
			shouldContinue = true
		}

		if k <= m/2 {
			for i := k+1; i <= m-k-1; i++ {
				res = append(res, matrix[i][n-k-1])
				fmt.Println(res)
			}
			shouldContinue = true
		}

		if k <= n/2 {
			for j := n-k-2; j >= k; j-- {
				res = append(res, matrix[m-k-1][j])
				fmt.Println(res)
			}
			shouldContinue = true
		}

		if k <= m/2 {
			for i := m-k-2; i >= k+1; i-- {
				res = append(res, matrix[i][k])
				fmt.Println(res)
			}
			shouldContinue = true
		}

		if shouldContinue {
			k++
			shouldContinue = false
		} else {
			break
		}
	}

	return res[:m*n]
}
