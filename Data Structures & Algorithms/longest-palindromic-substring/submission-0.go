func longestPalindrome(s string) string {
    res := ""
	resLen := 0

	for idx := range s {
		l, r := idx, idx 
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > resLen {
				resLen = r - l + 1
				res = s[l:r+1]
			}
			l--
			r++
		} 

		l, r = idx, idx+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > resLen {
				resLen = r - l + 1
				res = s[l:r+1]
			}
			l--
			r++
		} 
	}

	return res
}
