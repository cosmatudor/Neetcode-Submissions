func pacificAtlantic(heights [][]int) [][]int {
    ans := [][]int{}

	rows, cols := len(heights), len(heights[0])
	var visited [][]int

	clearVisited := func() {
		visited = make([][]int, rows)
		for i := 0; i < rows; i++ {
			visited[i] = make([]int, cols)
		}
	}

	var dfs func(int, int ,int)
	dfs = func(i, j, prev int) {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return
		}

		if visited[i][j] == 1 {
			return
		}

		if prev < heights[i][j] {
			return
		}

		visited[i][j] = 1
		dfs(i+1, j, heights[i][j])
		dfs(i-1, j, heights[i][j])
		dfs(i, j+1, heights[i][j])
		dfs(i, j-1, heights[i][j])
	}

	checkVisited := func() bool {
		var isPacific, isAtlantic bool
		for i := 0; i < rows; i++ {
			isPacific = isPacific || (visited[i][0] == 1)
		}
		for j := 0; j < cols; j++ {
			isPacific = isPacific || (visited[0][j] == 1)
		}

		for i := 0; i < rows; i++ {
			isAtlantic = isAtlantic || (visited[i][cols-1] == 1)
		}
		for j := 0; j < cols; j++ {
			isAtlantic = isAtlantic || (visited[rows-1][j] == 1)
		}

		return isPacific && isAtlantic
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			clearVisited()
			dfs(i, j, 1001)
			if checkVisited() {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}
