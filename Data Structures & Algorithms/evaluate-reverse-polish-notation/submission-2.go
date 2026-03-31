func evalRPN(tokens []string) int {
	stack := []int{}
	
	for _, token := range tokens {
		switch token {
			case "+":
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1 + op2)
			case "-":
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1 - op2)
			case "*":
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1 * op2)
			case "/":
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1 / op2)
			default:
				n, _ := strconv.Atoi(token)
				stack = append(stack, n)
		}
	}

	return stack[len(stack)-1]
}
