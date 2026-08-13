func checkInclusion(s1 string, s2 string) bool {
	windowSize := len(s1)
	count := make(map[byte]int)

	for _, ch := range s1 {
		count[byte(ch)]++
	}

	i, j := 0, 0
	for j < len(s2) {
		// count[s2[j]]++

		if j - i + 1 < windowSize {
			j++
			continue
		}

		freqs := computeFreq(i, j, s2)
		
		ok := true
		for ch := range count {
			if count[ch] != freqs[ch] {
				ok = false
			} 
		}

		if ok {
			return true
		}

		i++
		j++
	}

	return false
}

func computeFreq(i, j int, s string) map[byte]int {
	freqs := make(map[byte]int)
	for k := i; k <= j; k++ {
		freqs[s[k]]++
	}

	return freqs
}
