func characterReplacement(s string, k int) int {
    maxFreq := 0
    res := 0
    count := make(map[byte]int)

    i, j := 0, 0
    for j < len(s) {
        count[s[j]]++
        maxFreq = max(maxFreq, count[s[j]])

        for maxFreq + k < j - i + 1 {
            count[s[i]]--
            i++
        }

        res = max(res, j - i + 1)
        j++
    }

    return res
}

// max(freq) + k >= j - i + 1