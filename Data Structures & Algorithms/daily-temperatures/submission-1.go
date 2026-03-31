func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	s := [][2]int{}
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if len(s) == 0 {
			res[i] = 0
		} else {
			temp := temperatures[i]
			for len(s) != 0 && temp >= s[len(s)-1][0] {
				s = s[:len(s)-1]
			}
			if len(s) == 0 {
				res[i] = 0
			} else {
				res[i] = s[len(s)-1][1] - i
			}
		}
		s = append(s, [2]int{temperatures[i], i})
	}

	return res
}
