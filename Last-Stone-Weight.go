1type MaxHeap []int
2
3func (h MaxHeap) Len() int { return len(h) }
4func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max heap
5func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
6
7func (h *MaxHeap) Push(x interface{}) {
8    *h = append(*h, x.(int))
9}
10
11func (h *MaxHeap) Pop() interface{} {
12    old := *h
13    n := len(old)
14    val := old[n-1]
15    *h = old[:n-1]
16    return val
17}
18
19func lastStoneWeight(stones []int) int {
20    h := MaxHeap(stones)
21    heap.Init(&h)
22
23    for h.Len() > 1 {
24        x := heap.Pop(&h).(int)
25        y := heap.Pop(&h).(int)
26
27        if x != y {
28           heap.Push(&h, x-y)
29        }
30    }
31
32    if h.Len() == 0 {
33        return 0
34    }
35
36    return heap.Pop(&h).(int)
37}