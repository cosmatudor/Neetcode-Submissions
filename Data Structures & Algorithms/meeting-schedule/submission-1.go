/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	prevEnd := 0
	for _, interval := range intervals {
		if interval.start < prevEnd {
			return false
		}
		prevEnd = interval.end
	}

	return true
}
