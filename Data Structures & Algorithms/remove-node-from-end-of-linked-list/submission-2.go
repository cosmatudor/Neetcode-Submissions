/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0

    temp := head
    for temp != nil {
        length++
        temp = temp.Next
    }

    target := length - n

    dummy := &ListNode{Next: head}
    curr := dummy
    for cnt := 0; cnt < target; cnt++ {
        curr = curr.Next
    }

    curr.Next = curr.Next.Next
    return dummy.Next
}
