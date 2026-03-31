func maxProduct(nums []int) int {
    res := nums[0]
    curMin, curMax := 1, 1
    for _, num := range nums {
        tmp := curMax * num
        curMax = max(num*curMax, max(num*curMin, num))
        curMin = min(tmp, min(num*curMin, num))
        res = max(res, curMax)
    }
    return res
}