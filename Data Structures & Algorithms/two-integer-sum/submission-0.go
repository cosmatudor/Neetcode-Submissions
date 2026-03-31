func twoSum(nums []int, target int) []int {
    i := 0
    j := len(nums) - 1

    for i < j && nums[i] + nums[j] != target {
        if nums[i] + nums[j] > target {
            j -= 1
        }
        if nums[i] + nums[j] < target {
            i += 1
        }
    } 

    return []int{i,j}
}
