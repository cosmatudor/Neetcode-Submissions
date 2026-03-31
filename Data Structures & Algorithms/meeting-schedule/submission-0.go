/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	lastInterval := intervals[0]
	for _, currInterval := range intervals[1:] {
		if lastInterval.end > currInterval.start {
			return false
		}
		lastInterval = currInterval
	}

	return true
}
