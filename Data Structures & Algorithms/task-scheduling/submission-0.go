type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _, task := range tasks {
		freq[task]++
	}
	
	time := 0
	queue := [][]int{}
	maxH := MaxHeap{}
	heap.Init(&maxH)

	for c := 'A'; c <= 'Z'; c++ {
		if freq[byte(c)] != 0 {
			heap.Push(&maxH, freq[byte(c)])
		}
	}

	for maxH.Len() > 0 || len(queue) > 0 {		
		if maxH.Len() > 0 {
			newFreqOfChar := heap.Pop(&maxH).(int) - 1
			if newFreqOfChar > 0 {
				queue = append(queue, []int{newFreqOfChar, time + n})
			}
		}

		for len(queue) > 0 && queue[0][1] <= time {
			heap.Push(&maxH, queue[0][0])
			queue = queue[1:]
	
		}

		time++
	}

	return time
}
