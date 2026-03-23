1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func reverseEvenLengthGroups(head *ListNode) *ListNode {
9    dummy := &ListNode{Next :  head}
10    prev := dummy
11
12    groupSize := 1
13
14    for prev.Next != nil {
15        curr := prev.Next
16        count := 0
17
18        for count < groupSize && curr != nil {
19            curr = curr.Next
20            count++
21        }
22
23        if count % 2 == 0 {
24            prev = reverse(prev, count)
25        } else {
26            for i := 0; i < count; i++ {
27                prev = prev.Next
28            }
29        }
30
31        groupSize++
32    }
33
34    return dummy.Next
35}
36
37func reverse(prev *ListNode, k int) *ListNode {
38    curr := prev.Next
39
40    for i := 1; i < k; i++ {
41        next := curr.Next
42        curr.Next = next.Next
43        next.Next = prev.Next
44        prev.Next = next
45    }
46
47    return curr
48}