1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func sortList(head *ListNode) *ListNode {
9    if head == nil || head.Next == nil {
10        return head
11    }
12
13    slow, fast := head, head.Next
14    for fast != nil && fast.Next != nil {
15        slow = slow.Next
16        fast = fast.Next.Next
17    }
18
19    mid := slow.Next
20    slow.Next = nil
21
22    left := sortList(head)
23    right := sortList(mid)
24
25    return merge(left, right)
26}
27
28func merge(l1, l2 *ListNode) *ListNode {
29    dummy := &ListNode{}
30    curr := dummy
31
32    for l1 != nil && l2 != nil {
33        if l1.Val < l2.Val {
34            curr.Next = l1
35            l1 = l1.Next
36        } else {
37            curr.Next = l2
38            l2 = l2.Next
39        }
40        curr = curr.Next
41    }
42
43    if l1 != nil {
44        curr.Next = l1
45    } else {
46        curr.Next = l2
47    }
48
49    return dummy.Next
50}