1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func removeNthFromEnd(head *ListNode, n int) *ListNode {
9    dummy := &ListNode{Next : head}
10
11    p1 := dummy
12    p2 := dummy
13
14    for i := 0; i <= n; i++ {
15        p2 = p2.Next
16    }
17
18    for p2 != nil {
19        p1 = p1.Next
20        p2 = p2.Next
21    }
22
23    p1.Next = p1.Next.Next
24
25    return dummy.Next
26}