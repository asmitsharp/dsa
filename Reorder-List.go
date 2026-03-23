1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func reorderList(head *ListNode)  {
9    if head == nil || head.Next == nil {
10        return
11    }
12    slow, fast := head, head
13
14    for fast != nil && fast.Next != nil {
15        slow = slow.Next
16        fast = fast.Next.Next
17    }
18
19    second := slow.Next
20    slow.Next = nil
21
22    var prev *ListNode
23    for second != nil {
24        next := second.Next
25        second.Next = prev
26        prev = second
27        second = next
28    }
29
30    first := head
31    second = prev
32
33    for second != nil {
34        tmp1 := first.Next
35        tmp2 := second.Next
36
37        first.Next = second
38        second.Next = tmp1
39
40        first = tmp1
41        second = tmp2
42    }
43
44}