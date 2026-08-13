func minWindow(s string, t string) string {
    freq := make(map[byte]int)
	freq2 := make(map[byte]int)
    for _, ch := range t {
        freq[byte(ch)]++
    }
	

	i , j := 0, 0
	resI, resJ := -1, -1
	res := 100001
	for j < len(s) {
		freq2[s[j]]++
		
		for sameFreq(freq, freq2) {
			if j - i + 1 < res {
				res = j - i + 1
				resI = i
				resJ = j
			}

			freq2[s[i]]--
			i++
		}

		j++
	}

	if resI == -1 && resJ == -1 {
		return ""
	}
	return s[resI:resJ+1]
}

func sameFreq(freq1, freq2 map[byte]int) bool {
	for ch := range freq1 {
		if freq1[ch] > freq2[ch] {
			return false
		}
	}

	return true
}
