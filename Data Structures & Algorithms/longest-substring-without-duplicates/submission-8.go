func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    set := make(map[byte]struct{})

    i := 0
    j := 1
    // pivot := s[i]
    max := 0
    set[s[0]] = struct{}{}
    for j < len(s) {
        set[s[j]] = struct{}{}

        if len(set) < j - i + 1 {
            if j - i > max {
                max = j - i
            }

            delete(set, s[i])
            i++
            for len(set) < j - i + 1{
                delete(set, s[i])
                i++
            }
        } else {
            j++
        }

        // if s[j] == s[i] {
        //     if j - i + 1 > max {
        //         max = j - i
        //     }
        //     freq[pivot]--

        //     i++
        //     for freq[s[i]] > 1 {
        //         i++
        //     }
        //     pivot = s[i]
        //     j = i + 1
        // } else {
        //     freq[s[j]]++
        //     j++
        // }
    }

    if j - i > max {
        max = j - i
    }

    return max
}

// p w x w



// z x x z x y z