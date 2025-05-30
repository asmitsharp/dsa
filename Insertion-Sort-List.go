/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := &ListNode{Val : 0, Next: head}
    curr := head.Next
    lastSorted := head

    for curr != nil {
        if curr.Val >= lastSorted.Val {
            lastSorted = curr
            curr = curr.Next
        } else {
            lastSorted.Next = curr.Next

            prev := dummy
            for prev.Next != nil && prev.Next.Val < curr.Val {
                prev = prev.Next
            }

            curr.Next = prev.Next
            prev.Next = curr

            curr = lastSorted.Next
        }
    }
    return dummy.Next
}