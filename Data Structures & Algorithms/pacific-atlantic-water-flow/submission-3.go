func pacificAtlantic(heights [][]int) [][]int {
	xs := [4]int{0, 1, 0, -1}
	ys := [4]int{-1, 0, 1, 0}

	rows, cols := len(heights), len(heights[0])
	touchPacific := make([][]int, rows)
	for i := 0; i < rows; i++ {
		touchPacific[i] = make([]int, cols)
	}
	touchAtlantic := make([][]int, rows)
	for i := 0; i < rows; i++ {
		touchAtlantic[i] = make([]int, cols)
	}

	bfs := func(visited [][]int, queue [][]int) {
		for len(queue) != 0 {
			top := queue[0]
			queue = queue[1:]
			visited[top[0]][top[1]] = 1

			for i := 0; i < 4; i++ {
				nextI := top[0] + xs[i]
				nextJ := top[1] + ys[i]

				if nextI < 0 || nextI >= rows || nextJ < 0 || nextJ >= cols {
					continue
				}

				if visited[nextI][nextJ] == 1 {
					continue
				}

				if heights[top[0]][top[1]] > heights[nextI][nextJ] {
					continue
				}

				queue = append(queue, []int{nextI, nextJ})
			}
		}
	}

	pacificQueue, atlanticQueue := [][]int{}, [][]int{}
	for i := 0; i < rows; i++ {
		pacificQueue = append(pacificQueue, []int{i, 0})
		atlanticQueue = append(atlanticQueue, []int{i, cols-1})
	}
	for j := 1; j < cols; j++ {
		pacificQueue = append(pacificQueue, []int{0, j})
		atlanticQueue = append(atlanticQueue, []int{rows-1, cols-1-j})
	}

	bfs(touchPacific, pacificQueue)
	bfs(touchAtlantic, atlanticQueue)

	computeRes := func(visitedAtlantic [][]int, visitedPacific [][]int) [][]int {
		ans := [][]int{}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if visitedAtlantic[i][j] == 1 && visitedPacific[i][j] == 1 {
					ans = append(ans, []int{i,j})
				}
			}
		}

		return ans
	}

	return computeRes(touchAtlantic, touchPacific)
}
