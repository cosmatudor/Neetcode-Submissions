func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }

    m := len(board)
    n := len(board[0])
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    xs := [4]int{0, 1, 0, -1}
    ys := [4]int{-1, 0, 1, 0}

    var dfs func(int, int, int) bool
    dfs = func(i, j, word_i int) bool {
        if board[i][j] != word[word_i] {
            return false
        }

        if word_i+1 == len(word) {
            return true
        }

        visited[i][j] = true

        for k := 0; k < 4; k++ {
            ni, nj := i+xs[k], j+ys[k]
            if ni < 0 || ni >= m || nj < 0 || nj >= n {
                continue
            }
            if visited[ni][nj] {
                continue
            }
            if dfs(ni, nj, word_i+1) {
                return true
            }
        }

        visited[i][j] = false
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}