func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := [][]int{} // 1. index, 2. value
	for i := 0; i < len(temperatures); i++ {
		for len(stack) != 0 && stack[len(stack)-1][1] < temperatures[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top[0]] = i - top[0]
		}

		stack = append(stack, []int{i, temperatures[i]})
	}
	
	return res
}
