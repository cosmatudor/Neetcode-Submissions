func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	top, bot := 0, rows - 1
	row := -1
	for top <= bot {
		row = (top + bot) / 2
		if target < matrix[row][0] {
			bot--
		} else if target > matrix[row][cols-1] {
			top++
		} else {
			break
		}
	}

	if row == -1 {
		return false
	}

	left, right := 0, cols-1
	for left <= right {
		mid := (left + right) / 2
		if target < matrix[row][mid] {
			right = mid - 1
		} else if target > matrix[row][mid] {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
