func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]struct{})
    max := 1
    cnt := 0
    for i := 0; i < len(s); i++ {
        _, exists := set[s[i]]
        if exists {
            if cnt > max {
                max = cnt
            }
            cnt = 1
        } else {
            cnt++
            set[s[i]] = struct{}{}
        }
    }

    return max
}
