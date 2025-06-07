type Heap struct {
\tValues   []int
\tLessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() int          { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
\th.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
\treturn v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

func NewHeap(less func(int, int) bool) *Heap {
\treturn &Heap{LessFunc: less}
}

type MedianFinder struct {
\tsmallHeap *Heap
\tlargeHeap *Heap
}

func Constructor() MedianFinder {
\treturn MedianFinder{
\t\tsmallHeap: NewHeap(func(a, b int) bool {
\t\t\treturn a > b
\t\t}),
\t\tlargeHeap: NewHeap(func(a, b int) bool {
\t\t\treturn a < b
\t\t}),
\t}
}

func (mf *MedianFinder) AddNum(num int) {
\tif (mf.smallHeap.Len()+mf.largeHeap.Len())%2 == 0 {
\t\theap.Push(mf.largeHeap, num)
\t\theap.Push(mf.smallHeap, heap.Pop(mf.largeHeap))
\t} else {
\t\theap.Push(mf.smallHeap, num)
\t\theap.Push(mf.largeHeap, heap.Pop(mf.smallHeap))
\t}
}

func (mf *MedianFinder) FindMedian() float64 {
\tif (mf.smallHeap.Len()+mf.largeHeap.Len())%2 == 0 {
\t\treturn (float64(mf.smallHeap.Peek()) + float64(mf.largeHeap.Peek())) / 2
\t}
\treturn float64(mf.smallHeap.Peek())
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */