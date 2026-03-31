func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)

    for i, x := range nums {
        diff := target - x
        y, exists := hashmap[diff]
        if exists {
            return []int{y, i}
        }

        hashmap[x] = i
    }

    return []int{0,0}
}
