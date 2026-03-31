import (
    "slices"
)

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return nil
    }

    n := len(nums) - 1

    slices.Sort(nums)

    var res [][]int

    lastI := nums[0] + 1

    for i := 0; i <= n - 2; i++ {
        if nums[i] == lastI {
            continue
        }

        target := (-1) * nums[i]

        j := i + 1
        k := n
        for j < k {
            sum := nums[j] + nums[k]
            if sum == target {
                res = append(res, []int{nums[i],nums[j],nums[k]})
                j++
                k--
            } else if sum < target {
                j++
            } else {
                k--
            }
        }

        lastI = nums[i]
    }

    return res
}
