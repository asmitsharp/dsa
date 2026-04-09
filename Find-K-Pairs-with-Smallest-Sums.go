1
2type Node struct {
3    i, j int
4    sum  int
5}
6
7type MinHeap []Node
8
9func (h MinHeap) Len() int { return len(h) }
10func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
11func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
12
13func (h *MinHeap) Push(x interface{}) {
14    *h = append(*h, x.(Node))
15}
16
17func (h *MinHeap) Pop() interface{} {
18    old := *h
19    n := len(old)
20    x := old[n-1]
21    *h = old[:n-1]
22    return x
23}
24
25func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
26    h := &MinHeap{}
27    heap.Init(h)
28
29    for i := 0; i < len(nums1) && i < k; i++ {
30        heap.Push(h, Node{i, 0, nums1[i] + nums2[0]})
31    }
32
33    res := [][]int{}
34
35    for k > 0 && h.Len() > 0 {
36         node := heap.Pop(h).(Node)
37        i, j := node.i, node.j
38
39        res = append(res, []int{nums1[i], nums2[j]})
40        k--
41        if j+1 < len(nums2) {
42            heap.Push(h, Node{i, j+1, nums1[i] + nums2[j+1]})
43        }
44    }
45    return res
46}