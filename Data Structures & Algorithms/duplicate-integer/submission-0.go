func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    for _, x := range nums {
        set[x] = struct{}{} 
    }

    return len(set) != len(nums)
}
