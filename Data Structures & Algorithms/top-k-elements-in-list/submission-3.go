func topKFrequent(nums []int, k int) []int {
    // 1. compute freq
    // 2. add in a backet coresponding to a certai freq

    freq := make(map[int]int)
    maxFreq := 0
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        if freq[nums[i]] > maxFreq {
            maxFreq = freq[nums[i]]
        }
    }
    
    buckets := make(map[int][]int)
    for val, count := range freq {
        if _, ok := buckets[count]; !ok {
            buckets[count] = []int{}
        }
        buckets[count] = append(buckets[count], val)
    }

    res := []int{}
    for i := maxFreq; i > 0; i-- {
        if len(buckets[i]) == 0 {
            continue
        }

        for len(buckets[i]) != 0  && k > 0 {
            res = append(res, buckets[i][len(buckets[i])-1])
            buckets[i] = buckets[i][:len(buckets[i])-1]
            k--
        }
    }

    return res
}
