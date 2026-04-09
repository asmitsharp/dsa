1import "container/heap"
2type MinHeap []int
3
4func (h MinHeap) Len() int {return len(h)}
5func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
6func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
7
8func (h *MinHeap) Push(x interface{}) {
9    *h = append(*h, x.(int))
10}
11
12func (h *MinHeap) Pop() interface{} {
13    old := *h
14    n := len(old)
15    x := old[n-1]
16    *h = old[:n-1]
17
18    return x
19}
20
21func findKthLargest(nums []int, k int) int {
22    h := &MinHeap{}
23    heap.Init(h)
24    
25    for _, num := range nums {
26        heap.Push(h, num)
27        if h.Len() > k {
28            heap.Pop(h)
29        }
30    }
31
32    return (*h)[0]
33}