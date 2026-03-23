1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func oddEvenList(head *ListNode) *ListNode {
9    if head == nil {
10        return head
11    }
12
13    odd := head
14    even := head.Next
15
16    evenHead := even
17
18    for even != nil && even.Next != nil {
19
20        odd.Next = even.Next
21        odd = odd.Next
22
23        even.Next = odd.Next
24        even = even.Next
25    }
26
27    odd.Next = evenHead
28    return head
29}