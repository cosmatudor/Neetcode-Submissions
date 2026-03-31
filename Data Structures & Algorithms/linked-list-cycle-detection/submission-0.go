/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    visited := make(map[int]bool)

    for head != nil {
        if visited[head.Val] {
            return true
        }
        visited[head.Val] = true

        head = head.Next
    }

    return false
}
