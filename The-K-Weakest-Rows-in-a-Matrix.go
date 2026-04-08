1type Pair struct {
2    ones int
3    idx  int
4}
5
6type MaxHeap []Pair
7
8func (h MaxHeap) Len() int { return len(h) }
9
10func (h MaxHeap) Less(i, j int) bool {
11    if h[i].ones == h[j].ones {
12        return h[i].idx > h[j].idx   // bigger index stronger
13    }
14    return h[i].ones > h[j].ones     // more 1s stronger
15}
16
17func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
18
19func (h *MaxHeap) Push(x interface{}) {
20    *h = append(*h, x.(Pair))
21}
22
23func (h *MaxHeap) Pop() interface{} {
24    old := *h
25    n := len(old)
26    val := old[n-1]
27    *h = old[:n-1]
28    return val
29}
30
31func kWeakestRows(mat [][]int, k int) []int {
32    h := &MaxHeap{}
33    heap.Init(h)
34
35    for i, row := range mat {
36        count := 0
37        for _, v := range row {
38            if v == 1 {
39                count++
40            } else {
41                break
42            }
43        }
44
45        heap.Push(h, Pair{count, i})
46
47        if h.Len() > k {
48            heap.Pop(h)
49        }
50    }
51
52    res := make([]int, k)
53    for i := k - 1; i >= 0; i-- {
54        res[i] = heap.Pop(h).(Pair).idx
55    }
56
57    return res
58}