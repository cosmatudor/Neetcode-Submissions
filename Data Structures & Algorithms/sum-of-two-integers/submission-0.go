func getSum(a int, b int) int {
	res := 0
	reminder := 0
	i := 0
	for a != 0 || b != 0 {
		x := a & 1
		y := b & 1

		bit := (x + y + reminder) % 2
		if x + y + reminder >= 2 {
			reminder = 1
		} else {
			reminder = 0
		}

		fmt.Println(a, b)

		res |= (bit & 1) << i

		a >>= 1
		b >>= 1
		i++
	}

	res |= (reminder & 1) << i

	return res
}
