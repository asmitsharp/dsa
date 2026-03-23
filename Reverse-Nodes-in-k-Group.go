1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func reverseKGroup(head *ListNode, k int) *ListNode {
9    dummy := &ListNode{Next : head}
10    prev := dummy
11
12    for {
13        kth := prev
14        for i := 0; i < k && kth != nil ; i++ {
15            kth = kth.Next
16        }
17
18        if kth == nil {
19            break
20        }
21
22        curr := prev.Next
23        nextGroup := kth.Next
24
25        var prevNode = nextGroup
26        for curr != nextGroup {
27            tmp := curr.Next
28            curr.Next = prevNode
29            prevNode = curr
30            curr = tmp
31        }
32
33        tmp := prev.Next
34        prev.Next = kth
35        prev = tmp
36    }
37
38    return dummy.Next
39}