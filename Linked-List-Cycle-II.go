/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    hasCycle := false
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            hasCycle = true
            break
        }
    }

    if !hasCycle {
        return nil
    }

    ptr := head
    for ptr != slow {
        ptr = ptr.Next
        slow = slow.Next
    }
    return ptr
}
