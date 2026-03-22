1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func reverseList(head *ListNode) *ListNode {
9    var prev *ListNode
10    curr := head 
11
12    for curr != nil {
13        next := curr.Next
14        curr.Next = prev
15        prev = curr
16        curr = next
17    }
18
19    return prev
20}