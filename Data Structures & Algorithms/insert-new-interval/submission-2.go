func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
    i := 0
    n := len(intervals)

    // 1. Add all intervals that come BEFORE the newInterval
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    }

    // 2. Merge all overlapping intervals into newInterval
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    // Add the final merged version of newInterval
    res = append(res, newInterval)

    // 3. Add all remaining intervals that come AFTER
    for i < n {
        res = append(res, intervals[i])
        i++
    }

    return res
}