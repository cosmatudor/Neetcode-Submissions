type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	fmt.Println(maxHeap)
	heap.Init(&maxHeap)

	for maxHeap.Len() >= 2 {
		stone1 := heap.Pop(&maxHeap).(int)
		stone2 := heap.Pop(&maxHeap).(int)

		if stone1-stone2 > 0 {
			heap.Push(&maxHeap, stone1-stone2)
		}
	}

	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(&maxHeap).(int)
}

