func trap(height []int) int {
    n := len(height)
    leftMax := make([]int, n)
    rightMax := make([]int, n)

    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i-1])
    }
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i+1])
    }

    res := 0
    for i, h := range height {
        res += max(min(leftMax[i], rightMax[i]) - h, 0)
    }

    return res
}
