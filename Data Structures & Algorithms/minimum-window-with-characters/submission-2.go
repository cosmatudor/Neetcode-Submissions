func minWindow(s string, t string) string {
    freq := make(map[byte]int)
    for _, ch := range t {
        freq[byte(ch)]++
    }
	
	count := len(t)
	i , j := 0, 0
	resI, resJ := -1, -1
	res := 100001
	for j < len(s) {
		if freq[s[j]] > 0 {
			count--
		}

		freq[s[j]]--
		
		for count == 0 {
			if j - i + 1 < res {
				res = j - i + 1
				resI = i
				resJ = j
			}

			freq[s[i]]++
			if freq[s[i]] > 0 {
				count++
			}
			i++
		}

		j++
	}

	if resI == -1 && resJ == -1 {
		return ""
	}
	return s[resI:resJ+1]
}