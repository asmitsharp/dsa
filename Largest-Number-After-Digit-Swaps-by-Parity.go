1type MaxHeap []int
2
3func (h MaxHeap) Len() int           { return len(h) }
4func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // FLIPPED → Max-Heap
5func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
6
7func (h *MaxHeap) Push(x interface{}) {
8    *h = append(*h, x.(int))
9}
10
11func (h *MaxHeap) Pop() interface{} {
12    old := *h
13    n := len(old)
14    x := old[n-1]
15    *h = old[:n-1]
16    return x
17}
18
19func largestInteger(num int) int {
20    digits := []int{}
21
22    for num > 0 {
23        digits = append(digits, num%10)
24        num /= 10
25    }
26
27    for i, j := 0, len(digits) - 1; i < j; i , j = i+1, j-1 {
28        digits[i] , digits[j] = digits[j], digits[i]
29    }
30
31    odd := &MaxHeap{}
32    even := &MaxHeap{}
33
34    heap.Init(odd)
35    heap.Init(even)
36
37    for _, d := range digits {
38        if d % 2 == 0 {
39            heap.Push(even, d)
40        } else {
41            heap.Push(odd, d)
42        }
43    }
44
45    res := 0
46    for _, d := range digits {
47        res *= 10 
48        if d%2 == 0 {
49            res += heap.Pop(even).(int)
50        } else {
51            res += heap.Pop(odd).(int)
52        }
53    }
54
55    return res
56}