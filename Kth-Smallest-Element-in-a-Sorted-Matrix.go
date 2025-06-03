
type Element struct {
    Val int
    Row int
    Col int
}


type MinHeap []*Element

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Element)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func kthSmallest(matrix [][]int, k int) int {
    minHeap := &MinHeap{}

    for col := 0; col < len(matrix[0]); col++ {
        heap.Push(minHeap, &Element{Val: matrix[0][col], Row: 0, Col: col})
    }

    for i := 0; i < k - 1; i++ {
        elem := heap.Pop(minHeap).(*Element)

        if elem.Row < len(matrix)-1 {
            heap.Push(minHeap, &Element{Val : matrix[elem.Row+1][elem.Col], Row : elem.Row + 1, Col : elem.Col})
        }
    }

    return heap.Pop(minHeap).(*Element).Val
}