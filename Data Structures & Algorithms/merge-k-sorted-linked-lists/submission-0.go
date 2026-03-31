/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }

    if list1 != nil {
        curr.Next = list1
    }
    if list2 != nil {
        curr.Next = list2
    }

    return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    resList := lists[0]
    for i := 1; i < len(lists); i++ {
        resList = mergeLists(resList, lists[i])
    }

    return resList
}
