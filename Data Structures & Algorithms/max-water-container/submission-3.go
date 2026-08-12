func maxArea(heights []int) int {
    i := 0
    j := len(heights) - 1

    maxArea := 0

    for i < j {
        min := heights[j]
        pivot := i

        if heights[i] < heights[j] {
            min = heights[i]
            pivot = j
        }

        area := min * (j - i)
        if area > maxArea {
            maxArea = area
        }

        if pivot == i{
            j--
        } else if pivot == j {
            i++
        }
    }

    return maxArea
}
