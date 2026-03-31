/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getMidNode(head *ListNode) *ListNode {
    slow := head
    fast := head

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        next := curr.Next

        curr.Next = prev

        prev = curr
        curr = next
    }

    return prev
}

func reorderList(head *ListNode) {
    midNode := getMidNode(head)
    fmt.Println("midNode", midNode.Val)

    // break the list in 2 from the middle
    secondHead := midNode.Next
    midNode.Next = nil

    // reverse the second list
    second := reverseList(secondHead)
    fmt.Println("second", second.Val)

    // merge the lists
    first := head

    for second != nil {
        fmt.Println("first value", first.Val)
        firstNext := first.Next
        secondNext := second.Next

        first.Next = second
        second.Next = firstNext

        first = firstNext
        second = secondNext
    }
}
