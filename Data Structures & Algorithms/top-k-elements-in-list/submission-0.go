func topKFrequent(nums []int, k int) []int {
    buckets := make([][]int, len(nums) + 1)
    freq := make(map[int]int)

    for _, x := range nums {
        freq[x] += 1
    }

    for i, _ := range freq {
        buckets[freq[i]] = append(buckets[freq[i]], i)
    }

    var res []int
    n := len(nums)
    for k > 0 && n > 0{
        for len(buckets[n]) == 0 {
            n -= 1
        }

        top := buckets[n][len(buckets[n])-1]
        buckets[n] = buckets[n][:len(buckets[n])-1]

        k -= 1
        res = append(res, top)
    }

    return res
}
