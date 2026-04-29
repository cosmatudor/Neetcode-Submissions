
type MaxHeap [][]int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	x1, y1 := h[i][0], h[i][1]
	x2, y2 := h[j][0], h[j][1]

	dist1 := math.Sqrt(float64(x1*x1 + y1*y1))
	dist2 := math.Sqrt(float64(x2*x2 + y2*y2))

	return dist1 > dist2 
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}

func kClosest(points [][]int, k int) [][]int {
	h := MaxHeap{}
	heap.Init(&h)

	for _, point := range points {
		heap.Push(&h, point)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	res := [][]int{}
	for len(h) != 0 {
		res = append(res, heap.Pop(&h).([]int))
	}

	return res
}
