1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func middleNode(head *ListNode) *ListNode {
9    if head == nil || head.Next == nil {
10        return head
11    }
12
13    slow, fast := head, head
14    for fast != nil && fast.Next != nil {
15        slow = slow.Next
16        fast = fast.Next.Next
17    }
18
19    return slow
20}