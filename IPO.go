type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } 
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
\t*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
\told := *h
\tn := len(old)
\tval := old[n-1]
\t*h = old[0 : n-1]
\treturn val
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
\tn := len(profits)

\tprojects := make([][2]int, n)
\tfor i := 0; i < n; i++ {
\t\tprojects[i] = [2]int{capital[i], profits[i]}
\t}

\tsort.Slice(projects, func(i, j int) bool {
\t\treturn projects[i][0] < projects[j][0]
\t})

\tmaxHeap := &MaxHeap{}
\theap.Init(maxHeap)

\ti := 0

\tfor j := 0; j < k; j++ {
\t\tfor i < n && projects[i][0] <= w {
\t\t\theap.Push(maxHeap, projects[i][1])
\t\t\ti++
\t\t}

\t\tif maxHeap.Len() == 0 {
\t\t\tbreak
\t\t}

\t\tw += heap.Pop(maxHeap).(int)
\t}

\treturn w
}