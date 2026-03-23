1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func rotateRight(head *ListNode, k int) *ListNode {
9    if head == nil || head.Next == nil || k == 0 {
10        return head
11    }
12
13    length := 1
14    tail := head
15    for tail.Next != nil {
16        tail = tail.Next
17        length++
18    }
19
20    k = k % length
21    if k == 0 {
22        return head
23    }
24
25    tail.Next = head //circular
26
27    newTail := head
28    steps := length - k
29    for i := 1; i < steps; i++ {
30        newTail = newTail.Next
31    } 
32
33    newHead := newTail.Next
34    newTail.Next = nil
35
36    return newHead
37}