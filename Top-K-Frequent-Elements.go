1
2type Pair struct {
3    num int
4    freq int
5}
6
7type MinHeap []Pair
8
9func (h MinHeap) Len() int {return len(h)}
10func (h MinHeap) Less(i, j int) bool {return h[i].freq < h[j].freq}
11func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
12
13func (h *MinHeap) Push(x interface{}) {
14  *h = append(*h, x.(Pair))
15}
16
17func (h *MinHeap) Pop() interface{} {
18    old := *h
19    n := len(old)
20    x := old[n-1]
21    *h = old[:n-1]
22
23    return x
24}
25
26func topKFrequent(nums []int, k int) []int {
27    freq := map[int]int{}
28    for _, num := range nums {
29        freq[num]++
30    }
31
32     h := &MinHeap{}
33     heap.Init(h)
34
35     for num, f := range freq {
36        heap.Push(h, Pair{num, f})
37        if h.Len() > k {
38            heap.Pop(h)
39        }
40     }
41
42     res := []int{}
43     for h.Len() > 0 {
44        res = append(res, heap.Pop(h).(Pair).num)
45     }
46     return res
47}