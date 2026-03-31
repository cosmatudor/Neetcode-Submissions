func characterReplacement(s string, k int) int {
    freq := make(map[byte]int)

    i := 0
    j := 0
    maxFreq := 0
    res := 0
    for j < len(s) {
        freq[s[j]]++ 
        if freq[s[j]] > maxFreq {
            maxFreq = freq[s[j]]
        } 


        for (j - i + 1) - maxFreq > k {
            freq[s[i]]--
            i++
        }

        if (j - i + 1) > res {
            res = j - i + 1
        }        

        j++
    }

    return res
}
