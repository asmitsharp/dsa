1/**
2 * Definition for singly-linked list.
3 * type ListNode struct {
4 *     Val int
5 *     Next *ListNode
6 * }
7 */
8 type NodeHeap []*ListNode
9
10func (h NodeHeap) Len() int            { return len(h) }
11func (h NodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
12func (h NodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
13func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
14func (h *NodeHeap) Pop() interface{} {
15    old := *h
16    n := len(old)
17    node := old[n-1]
18    *h = old[:n-1]
19    return node
20}
21func mergeKLists(lists []*ListNode) *ListNode {
22    h := &NodeHeap{}
23    heap.Init(h)
24
25    for _, list := range lists {
26        if list != nil {
27            heap.Push(h, list)
28        }
29    }
30
31    dummy := &ListNode{}
32    curr := dummy
33
34    for h.Len() > 0 {
35        node := heap.Pop(h).(*ListNode)
36        curr.Next = node
37        curr = curr.Next
38
39        if node.Next != nil {
40            heap.Push(h, node.Next)
41        }
42    }
43
44    return dummy.Next
45}