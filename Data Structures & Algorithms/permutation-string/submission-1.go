func checkInclusion(s1 string, s2 string) bool {
	windowSize := len(s1)
	count := make(map[byte]int)
	count2 := make(map[byte]int)

	for _, ch := range s1 {
		count[byte(ch)]++
	}

	i, j := 0, 0
	for j < len(s2) {
		count2[s2[j]]++

		if j - i + 1 < windowSize {
			j++
			continue
		}

		ok := true
		for ch := range count {
			if count[ch] != count2[ch] {
				ok = false
			} 
		}

		if ok {
			return true
		}

		count2[s2[i]]--
		i++
		j++
	}

	return false
}