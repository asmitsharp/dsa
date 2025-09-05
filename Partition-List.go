/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    beforeDummy := &ListNode{}
    afterDummy := &ListNode{}

    before := beforeDummy
    after := afterDummy

    curr := head
    
    for curr != nil {
        if curr.Val < x {
            before.Next = curr
            before = before.Next
        } else {
            after.Next = curr
            after = after.Next
        }
        curr = curr.Next
    }

    after.Next = nil
    before.Next = afterDummy.Next
    return beforeDummy.Next
    
}