/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	 sort.Slice(intervals, func(i,j int) bool {
		return intervals[i].end < intervals[j].end
	})

	meetingsPerDay := []int{intervals[0].end}
	for _, currInterval := range intervals[1:] {
		i := 0
		for i < len(meetingsPerDay) && currInterval.start < meetingsPerDay[i] {
			i++
		}
		if i == len(meetingsPerDay) {
			meetingsPerDay = append(meetingsPerDay, currInterval.end)
		} else {
			meetingsPerDay[i] = currInterval.end
		}
	}

	return len(meetingsPerDay)
}
