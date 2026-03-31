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

    for i := 0; i <= n - 2; i++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }

        target := (-1) * nums[i]

        j := i + 1
        k := n
        for j < k {
            sum := nums[j] + nums[k]
            if sum == target {
                res = append(res, []int{nums[i],nums[j],nums[k]})

                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                
                for j < k && nums[k] == nums[k-1] {
                    k--
                }

                j++
                k--
            } else if sum < target {
                j++
            } else {
                k--
            }
        }
    }

    return res
}
