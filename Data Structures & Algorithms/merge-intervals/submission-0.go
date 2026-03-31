func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
    	return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		nextInterval := intervals[i]
		currInterval := res[len(res)-1]

		if currInterval[1] >= nextInterval[0] {
			currInterval[1] = max(currInterval[1], nextInterval[1])
		} else {
			res = append(res, nextInterval)
		}
	}

	return res
}
