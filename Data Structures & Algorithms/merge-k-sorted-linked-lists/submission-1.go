/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MaxHeap []*ListNode
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(*ListNode))
}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := MaxHeap{}
    heap.Init(&h)

    for _, list := range lists {
        heap.Push(&h, list)
    }

    aux := &ListNode{}
    res := &ListNode{Next: aux}
    for h.Len() != 0 {
        elem := heap.Pop(&h).(*ListNode)
        aux.Next = elem
        aux = aux.Next
        if elem.Next != nil {
            heap.Push(&h, elem.Next)
        }
    }

    return res.Next.Next
}
