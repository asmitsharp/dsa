1type MinHeap []string
2
3func (h MinHeap) Len() int { return len(h) }
4
5func (h MinHeap) Less(i, j int) bool {
6    if len(h[i]) == len(h[j]) {
7        return h[i] < h[j]
8    }
9    return len(h[i]) < len(h[j])
10}
11
12func (h MinHeap) Swap(i, j int) {
13    h[i], h[j] = h[j], h[i]
14}
15
16func (h *MinHeap) Push(x interface{}) {
17    *h = append(*h, x.(string))
18}
19
20func (h *MinHeap) Pop() interface{} {
21    old := *h
22    n := len(old)
23    x := old[n-1]
24    *h = old[:n-1]
25    return x
26}
27
28func kthLargestNumber(nums []string, k int) string {
29    h := &MinHeap{}
30    heap.Init(h)
31
32    for _, num := range nums {
33        heap.Push(h, num)
34        if h.Len() > k {
35            heap.Pop(h)
36        }
37    }
38
39    return (*h)[0]
40}