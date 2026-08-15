func islandsAndTreasure(grid [][]int) {
    m, n := len(grid), len(grid[0])
    queue := [][]int{}

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                queue = append(queue, []int{i, j})
            }
        }
    }

    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        for _, d := range dirs {
            ni, nj := curr[0]+d[0], curr[1]+d[1]
            if ni < 0 || ni >= m || nj < 0 || nj >= n {
                continue
            }
            if grid[ni][nj] != math.MaxInt32 {
                continue
            }
            grid[ni][nj] = grid[curr[0]][curr[1]] + 1
            queue = append(queue, []int{ni, nj})
        }
    }
}