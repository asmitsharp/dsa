1type Node struct {
2    val int
3    r   int
4    c   int
5}
6
7type MinHeap []Node
8
9func (h MinHeap) Len() int { return len(h) }
10func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
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
25func kthSmallest(matrix [][]int, k int) int {
26    n := len(matrix)
27    h := &MinHeap{}
28    heap.Init(h)
29
30    for r := 0; r < n; r++ {
31        heap.Push(h, Node{matrix[r][0], r, 0})
32    }
33
34    for i := 0; i < k-1; i++ {
35        node := heap.Pop(h).(Node)
36
37        if node.c+1 < n {
38            heap.Push(h, Node{matrix[node.r][node.c+1], node.r, node.c+1})
39        }
40    }
41    return heap.Pop(h).(Node).val
42}