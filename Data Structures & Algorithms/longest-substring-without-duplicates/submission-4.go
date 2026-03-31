func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]struct{})

    i := 0
    j := 1
    // pivot := s[i]
    max := 0

    for j < len(s) {
        set[s[j]] = struct{}{}
        if len(set) < j - i {
            if j - i > max {
                max = j - i
            }

            i++
            for len(set) < j - i {
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

    return max - 1
}

// p w x w



// z x x z x y z