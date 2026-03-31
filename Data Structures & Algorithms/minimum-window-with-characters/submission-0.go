func minWindow(s string, t string) string {
    freq := make(map[byte]int)
    for _, ch := range t {
        freq[byte(ch)]++
    }

    i := 0
    j := 0
    count := len(t)
    res_i, res_j := 0, 0
    min := 1001

    for j < len(s) {
        if freq[s[j]] > 0 {
            count--
        }

        freq[s[j]]--

        for count == 0 {
            if j - i + 1 < min {
                min = j - i + 1
                res_i, res_j = i, j
            }

            freq[s[i]]++

            if freq[s[i]] > 0 {
                count++
            }
            i++
        }

        j++
    }

    if min == 1001 {
        return ""
    }
    return s[res_i:res_j+1]

}
