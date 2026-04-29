type Pair struct {
	freq 	int
	val		int
}

type MinHeap []Pair
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() any {
    n := len(*h)
    pair := (*h)[n-1]
    *h = (*h)[:n-1]
    return pair
}

func topKFrequent(nums []int, k int) []int {
	h := MinHeap{}
    heap.Init(&h)

    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }

    for num, freq := range count {
        pair := Pair{freq, num}
        heap.Push(&h, pair)
        if h.Len() > k {
            heap.Pop(&h)
        }
    }

    res := make([]int, len(h))
    for i := k-1 ; i >= 0 && h.Len() > 0; i-- {
        pair := heap.Pop(&h).(Pair)
        res[i] = pair.val
    }

    return res
}
