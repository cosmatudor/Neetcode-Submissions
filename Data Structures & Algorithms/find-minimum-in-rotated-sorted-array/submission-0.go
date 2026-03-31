func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    mid := len(nums) / 2

    minLeft := findMin(nums[:mid])
    minRight := findMin(nums[mid:])

    return min(minLeft, minRight)
}
