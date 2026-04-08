1
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
17    return x
18}
19
20type KthLargest struct {
21   k int 
22   heap MinHeap
23}
24
25
26func Constructor(k int, nums []int) KthLargest {
27    kl := KthLargest {
28        k : k,
29    }
30
31    for _ , num := range nums {
32        kl.Add(num)
33    }
34
35    return kl
36}
37
38
39func (this *KthLargest) Add(val int) int {
40    if len(this.heap) < this.k {
41        heap.Push(&this.heap, val)
42    } else if val > this.heap[0] {
43       heap.Pop(&this.heap)
44       heap.Push(&this.heap, val)
45    }
46
47    return this.heap[0]
48}
49
50
51/**
52 * Your KthLargest object will be instantiated and called as such:
53 * obj := Constructor(k, nums);
54 * param_1 := obj.Add(val);
55 */