/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    prev := dummy
    curr := head

    for curr != nil {
        val := curr.Val
        ctn := 0
        
        for runner := curr; runner != nil && runner.Val == val; runner = runner.Next {
            ctn++
        }

        if ctn == 1 {
            prev.Next = curr
            prev = curr
        }

        for i := 0; i < ctn; i++ {
            curr = curr.Next
        }
    }

    prev.Next = nil
    return dummy.Next
}