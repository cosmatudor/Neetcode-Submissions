func setZeroes(matrix [][]int) {
    ROWS, COLS := len(matrix), len(matrix[0])
    rows := make([]bool, ROWS)
    cols := make([]bool, COLS)

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if matrix[r][c] == 0 {
                rows[r] = true
                cols[c] = true
            }
        }
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if rows[r] || cols[c] {
                matrix[r][c] = 0
            }
        }
    }
}