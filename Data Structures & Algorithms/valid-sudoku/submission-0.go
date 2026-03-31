func isValidSudoku(board [][]byte) bool {
	set := make(map[byte]struct{})

	y_dir := [9]int{0, -1, -1, 0, 1, 1, 1, 0, -1}
	x_dir := [9]int{0, 0, 1, 1, 1, 0, -1, -1, -1}

	for i := 0; i < 9; i++ {
		set = map[byte]struct{}{}
		cnt := 0
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				set[board[i][j]] = struct{}{}
				cnt++
			}
		}

		if len(set) != cnt {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		set = map[byte]struct{}{}
		cnt := 0
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				set[board[j][i]] = struct{}{}
				cnt++
			}
		}

		if len(set) != cnt {
			return false
		}
	}

	for i := 1; i < 9; i += 3 {
		for j := 1; j < 9; j+= 3 {
			set = map[byte]struct{}{}
			cnt := 0
			for k := 0; k < 9; k++ {
				if board[i+x_dir[k]][j+y_dir[k]] != '.' {
					set[board[i+x_dir[k]][j+y_dir[k]]] = struct{}{}
					cnt++
				}
			}
			if len(set) != cnt {
				return false
			}
		}
	}

	return true
}
