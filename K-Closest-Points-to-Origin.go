1
2type Point struct {
3    x int
4    y int
5    dist int
6}
7
8type MaxHeap []Point
9
10func (h MaxHeap) Len() int {return len(h)}
11func (h MaxHeap) Less(i, j int) bool {return h[i].dist > h[j].dist}
12func (h MaxHeap) Swap(i, j int)  {h[i], h[j] = h[j], h[i]}
13
14func (h *MaxHeap) Push(x interface{}) {
15    *h = append(*h, x.(Point))
16}
17
18func (h *MaxHeap) Pop() interface{} {
19    old := *h
20    n := len(old)
21    x := old[n-1]
22    *h = old[:n-1]
23    return x
24}
25
26func kClosest(points [][]int, k int) [][]int {
27    h := &MaxHeap{}
28    heap.Init(h)
29
30    for _, p := range points {
31        dist := p[0]*p[0] + p[1]*p[1]
32
33        heap.Push(h, Point{
34            x : p[0],
35            y : p[1],
36            dist : dist,
37        })
38
39        if h.Len() > k {
40            heap.Pop(h)
41        }
42    }
43
44    res := make([][]int, k)
45    for i := 0; i < k; i++ {
46        p := heap.Pop(h).(Point)
47        res[i] = []int{p.x, p.y}
48    }
49
50    return res
51}