func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    capture := func() {
        q := [][]int{}
        for r := 0; r < rows; r++ {
            for c := 0; c < cols; c++ {
                if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
                    if board[r][c] == 'O' {
                        q = append(q, []int{r, c})
                    }
                }
            }
        }
        for len(q) > 0 {
            // Dequeue
            r, c := q[0][0], q[0][1]
            q = q[1:]
            if board[r][c] == 'O' {
                board[r][c] = 'T'
                for _, dir := range directions {
                    nr, nc := r+dir[0], c+dir[1]
                    if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
                        q = append(q, []int{nr, nc})
                    }
                }
            }
        }
    }

    capture()

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if board[r][c] == 'O' {
                board[r][c] = 'X'
            } else if board[r][c] == 'T' {
                board[r][c] = 'O'
            }
        }
    }
}