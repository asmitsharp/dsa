1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8func isPalindrome(head *ListNode) bool {
9    slow, fast := head, head
10
11    for fast != nil && fast.Next != nil {
12        slow = slow.Next
13        fast = fast.Next.Next
14    }
15
16    curr := slow
17    var prev *ListNode
18    for curr != nil {
19        next := curr.Next
20        curr.Next = prev
21        prev = curr
22        curr = next
23    }
24
25    left , right := head, prev
26    for right != nil {
27        if left.Val != right.Val {
28            return false
29        }
30        left = left.Next
31        right = right.Next
32    }
33    return true
34}