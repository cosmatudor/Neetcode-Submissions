func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var freq [26]int

    for i := range s {
        freq[s[i] - 'a'] += 1
        freq[t[i] - 'a'] -= 1
    }

    for _, x := range freq {
        if x != 0 {
            return false
        }
    }
    return true
}
