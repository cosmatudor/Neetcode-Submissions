/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
 
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}


func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	minHeap := MinHeap{}

	heap.Push(&minHeap, intervals[0].end)
	for _, currInterval := range intervals[1:] {
		top := minHeap[0]
		if currInterval.start < top {
			heap.Push(&minHeap, currInterval.end)
		}
	} 

	return minHeap.Len()
}
