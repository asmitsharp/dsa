// func sortArray(nums []int) []int {
//     if len(nums) <= 1 {
//         return nums
//     }

//     mid := len(nums)/2
//     left := sortArray(nums[:mid])
//     right := sortArray(nums[mid:])

//     return merge(left, right)
// }

// func merge(left, right []int) []int {
//     result := make([]int, 0, len(left)+len(right))

//     i,j := 0,0

//     for i < len(left) && j < len(right) {
//         if left[i] <= right[j] {
//             result = append(result, left[i])
//             i++
//         } else {
//             result = append(result, right[j])
//             j++
//         }
//     }

//     for i < len(left) {
//         result = append(result, left[i])
//         i++
//     }

//     for j < len(right) {
//         result = append(result, right[j])
//         j++
//     }

//     return result
// }

// Quick Sort

// func sortArray(nums []int) []int {
//     if len(nums) <= 1 {
//         return nums
//     }

//     helper(nums, 0, len(nums) - 1)
//     return nums
// }


// func helper(nums []int, low, high int) {
//  if low < high {
//     pivotIndex := partition(nums, low, high)

//     helper(nums, low, pivotIndex - 1)
//     helper(nums, pivotIndex, high)
//  }
// }

// func partition(nums []int, low, high int) int {
//     pivot := nums[high]

//     i := low - 1

//     for j := low; j < high; j++ {
//         if nums[j] <= pivot {
//             i++
//             nums[i], nums[j] = nums[j], nums[i]
//         }
//     }

//     nums[i+1], nums[high] = nums[high], nums[i+1]
//     return i + 1
// }

// Heap Sort 

func sortArray(nums []int) []int {
    heapSort(nums)
    return nums
}

func heapSort(arr []int) {
    n := len(arr)
    
    // Build max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    
    // Extract elements from heap
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
}

func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2
    
    if left < n && arr[left] > arr[largest] {
        largest = left
    }
    if right < n && arr[right] > arr[largest] {
        largest = right
    }
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}


