
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

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: MinHeap{},
		maxHeap: MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 {
		heap.Push(&this.minHeap, num)
		return
	}

	if num > this.minHeap[0] {
		heap.Push(&this.minHeap, num)
	} else {
		heap.Push(&this.maxHeap, num)
	}

	if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		x := heap.Pop(&this.minHeap).(int)
		heap.Push(&this.maxHeap, x)
	}

	if this.maxHeap.Len()-this.minHeap.Len() > 0 {
		x := heap.Pop(&this.maxHeap).(int)
		heap.Push(&this.minHeap, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len()-this.minHeap.Len() == 0 {
		fmt.Println(this.maxHeap[0], this.minHeap[0])
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
	}

	return float64(this.minHeap[0])
}
