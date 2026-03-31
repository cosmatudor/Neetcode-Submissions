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

    cnt := 0
    target := length - n

    dummy := &ListNode{Val: 0, Next: head}
    prev := dummy
    curr := head
    for curr != nil {
        if cnt == target {
            prev.Next = curr.Next
            break
        }
        prev = curr
        curr = curr.Next
        cnt++
    }

    return dummy.Next
}
