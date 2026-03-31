func search(nums []int, t int) int {
    // t < mid
    //      left < t => right = mid - 1
    //      left > t => left = mid + 1
    // mid < t
    //      t < right => left = mid + 1
    //      t > right => right = mid - 1

    left, right := 0, len(nums)-1

    for left < right {
        mid := (left + right) / 2

        if t == nums[mid] {
            return mid
        }

        if t < nums[mid] {
            if nums[left] < t {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if t < nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}
