func maxArea(heights []int) int {
    i := 0
    j := len(heights) - 1

    maxArea := 0
    pivot := -1

    for i < j {
        min := heights[j]
        if heights[i] < heights[j]{
            min = heights[i]
            pivot = j
        } else {
            min = heights[j]
            pivot = i
        }

        area := min * (j - i)
        if area > maxArea {
            maxArea = area
        }

        if pivot == i {
            for heights[pivot] >= heights[j] && i < j {
                j--
            }
        } else if pivot == j {
            for heights[i] < heights[pivot] && i < j {
                i++
            }
        }
    }

    return maxArea
}
