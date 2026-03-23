1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func swapPairs(head *ListNode) *ListNode {
9
10    dummy := &ListNode{Next: head}
11    prev := dummy
12
13    for prev.Next != nil && prev.Next.Next != nil {
14        a := prev.Next
15        b := a.Next
16
17        a.Next = b.Next
18        b.Next = a
19        prev.Next = b
20
21        prev = a
22    }
23    
24    return dummy.Next
25}