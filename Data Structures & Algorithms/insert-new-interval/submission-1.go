func areOverlapping(interval1, interval2 []int) bool {
	return (interval1[1] >= interval2[0] && interval1[0] <= interval2[0])  || (interval2[0] <= interval1[1] && interval2[0] <= interval1[0])
}

func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
	var index int
	for index = 0; index < len(intervals); index++ {
		if intervals[index][1] >= newInterval[0] {
			break
		}

		res = append(res, intervals[index])
	}

	fmt.Println("intervals", res)

	for index < len(intervals) && areOverlapping(newInterval, intervals[index]){
		newInterval = []int{min(newInterval[0], intervals[index][0]), max(newInterval[1], intervals[index][1])}
		index++
	}

	fmt.Println(newInterval)
	fmt.Println(index)


	res = append(res, newInterval)
	if index <= len(intervals) - 1 {
		res = append(res, intervals[index:]...)
	}

	return res
}
