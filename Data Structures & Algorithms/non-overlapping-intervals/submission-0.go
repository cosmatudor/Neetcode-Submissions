func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int) bool {
		return (intervals[i][1] < intervals[j][1]) ||
			((intervals[i][1] == intervals[j][1]) && (intervals[i][0] > intervals[j][0]))
	})

	fmt.Println(intervals)
	maxEnd := intervals[0][1]
	cnt := 1
	for i := 1; i <= len(intervals) - 1; i++ {
		if maxEnd <= intervals[i][0] {
			maxEnd = intervals[i][1]
			cnt++
		}
	}

	return len(intervals) - cnt
}
