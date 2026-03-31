func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]struct{})
    max := 0
    i := 0

    // j is the 'right' boundary of our window
    for j := 0; j < len(s); j++ {
        // While s[j] is already in the window, shrink from the left
        for {
            _, exists := set[s[j]]
            if !exists {
                break
            }
            delete(set, s[i])
            i++
        }

        // Now we are sure s[j] is unique in our window
        set[s[j]] = struct{}{}
        
        // Calculate current window size
        windowSize := j - i + 1
        if windowSize > max {
            max = windowSize
        }
    }

    return max
}