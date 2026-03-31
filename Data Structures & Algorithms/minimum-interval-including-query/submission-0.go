type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1]}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *MinHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	qc := make([]int, len(queries))
	copy(qc, queries)
	sort.Slice(qc, func(i, j int) bool {
		return qc[i] < qc[j]
	})

	bookKeeping := make(map[int]int)
	minHeap := MinHeap{}
	i := 0
	for _, q := range qc {
		for i < len(intervals) && intervals[i][0] <= q {
			heap.Push(&minHeap, [2]int{intervals[i][1]-intervals[i][0]+1, intervals[i][1]})
			i++
		}

		for minHeap.Len() != 0 && minHeap[0][1] < q {
			heap.Pop(&minHeap)
		}

		var top int
		if minHeap.Len() != 0 {
			top = minHeap[0][0]
		} else {
			top = -1
		}

		bookKeeping[q] = top
	}

	res := []int{}
	for _, q := range queries {
		res = append(res, bookKeeping[q])
	}

	return res
}
