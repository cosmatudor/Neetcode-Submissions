func lengthOfLongestSubstring(s string) int {
    maxLen := 0
    set := make(map[byte]struct{})

    i, j := 0, 0
    for j < len(s) {
        for {
            if _, ok := set[s[j]]; !ok {
                break
            }
            
            delete(set, s[i])
            i++
        }

        set[s[j]] = struct{}{}

        if j - i + 1 > maxLen {
            maxLen = j - i + 1
        }

        j++
    }

    return maxLen
}
