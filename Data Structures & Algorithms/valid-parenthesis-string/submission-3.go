func checkValidString(s string) bool {
	// *)(()*()
	// loop
		// if closed > open
			// if enough *, continue
			// otherwise, stop -> false
	// if open - closed > remaining len(*) -> false
	// otherwise -> true

	// ((((()(((

	minOpen , maxOpen := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
			case '(':
			minOpen++
			maxOpen++

			case ')':
			minOpen--
			maxOpen--

			case '*':
			minOpen--
			maxOpen++
		}

		if maxOpen < 0 {
			return false
		}

		if minOpen < 0 {
			minOpen = 0
		}
	}

	return minOpen == 0
}
