
func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0, len(nums)-k+1)
    // Deque jisme indices store honge
    deque := make([]int, 0)

    for i := 0; i < len(nums); i++ {
        // 1. Remove indices jo window ke bahar hain
        if len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }

        // 2. Remove chhote elements ke indices from back
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
            deque = deque[:len(deque)-1]
        }

        deque = append(deque, i)

        // 4. Agar window size k ho gayi, to max (front of deque) ko result mein add karo
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }

    return result
}