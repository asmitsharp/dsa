/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
     if head == nil || head.Next == nil || k == 0 {
        return head
    }

    l := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        l++
    }

    tail.Next = head

    k = k % l
    if k == 0 {
        tail.Next = nil
        return head
    }

    newHead := l - k
    newTail := head
    for i := 1; i < newHead; i++ {
        newTail = newTail.Next
    }
    nHead := newTail.Next

    newTail.Next = nil

    return nHead

}