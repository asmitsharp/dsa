// import (
// \t\container/heap\
// )

// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int))}
// func (h *IntHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0 : n-1]
//     return x
// }

// func findKthLargest(nums []int, k int) int {
//     h := &IntHeap{}
//     heap.Init(h)

//     for _, num := range nums {
//         if h.Len() < k {
//             heap.Push(h, num)
//         } else if num > (*h)[0] {
//             heap.Pop(h)
//             heap.Push(h, num)
//         }
//     }

//     return (*h)[0]
// }


// func findKthLargest(nums []int, k int) int {

//     if len(nums) == 0 || k <= 0 || k > len(nums) {
//         return -1
//     }
   
//     for i := 0; i < k; i++ {
//         maxIndex := i
//         for j := i + 1 ; j < len(nums); j++ {
//             if nums[maxIndex] < nums[j] {
//                 maxIndex = j
//             }
//         }

//         nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
//     }

//     return nums[k-1]
// }


// Max Heap

func findKthLargest(nums []int, k int) int {
    n := len(nums)

    for i := n/2 - 1; i >= 0; i-- {
        heapify(nums, n, i)
    }

    for i := 0; i < k - 1; i++ {
        nums[0], nums[n-1-i] = nums[n-1-i], nums[0]
        heapify(nums, n-1-i, 0)
    }

    return nums[0]
}

func heapify(nums []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && nums[left] > nums[largest] {
        largest = left
    }
    if right < n && nums[right] > nums[largest] {
        largest = right
    }

    if largest != i {
        nums[i], nums[largest] = nums[largest], nums[i]
        heapify(nums, n, largest)
    }
}