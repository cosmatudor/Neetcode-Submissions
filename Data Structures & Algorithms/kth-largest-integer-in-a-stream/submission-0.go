

// Definim un tip nou care va implementa interfața heap.Interface
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type KthLargest struct {
    minHeap *IntHeap
    k       int
}

func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    heap.Init(h)
    
    kl := KthLargest{
        minHeap: h,
        k:       k,
    }
    
    for _, num := range nums {
        kl.Add(num)
    }
    
    return kl
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this.minHeap, val)
    
    // Menținem dimensiunea heap-ului la exact k
    if this.minHeap.Len() > this.k {
        heap.Pop(this.minHeap)
    }
    
    // Vârful min-heap-ului este al k-lea cel mai mare element
    return (*this.minHeap)[0]
}