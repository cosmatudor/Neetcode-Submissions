/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    prevNode := nil
    currNode := head

    for currNode != nil {
        nextNode := currNode.Next
        currNode.Next = prevNode

        prevNode = currNode
        currNode = nextNode
    }

    head = prevNode
    return head
}
