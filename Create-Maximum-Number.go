func maxNumber(nums1 []int, nums2 []int, k int) []int {
    result := make([]int, 0, k)

    for i := max(0, k-len(nums2)); i <= min(k, len(nums1)); i++ {
        seq1 := maxSubs(nums1, i)
        seq2 := maxSubs(nums2, k - i)

        merged := merge(seq1, seq2)

        if isGreater(merged, result) {
            result = merged
        }
    }

    return result
}

func maxSubs(num []int , k int) []int {
    stack := make([]int, 0, k)

    for i := 0; i < len(num); i++ {
        for len(stack) > 0 && stack[len(stack) - 1] < num[i] && len(stack) + (len(num) - i) > k {
            stack = stack[:len(stack) - 1]
        }

        if len(stack) < k {
            stack = append(stack, num[i])
        }
    }

    return stack
}

func merge(seq1, seq2 []int) []int {
    result := make([]int, 0, len(seq1) + len(seq2))
    i, j := 0, 0

    for i < len(seq1) && j < len(seq2) {

        if seq1[i] > seq2[j] || (seq1[i] == seq2[j] && isGreater(seq1[i:], seq2[j:])) {
            result = append(result, seq1[i])
            i++
        } else {
            result = append(result, seq2[j])
            j++
        }
    }

    result = append(result, seq1[i:]...)
    result = append(result, seq2[j:]...)

    return result
}

func isGreater(a, b []int) bool {
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] > b[i] {
            return true
        }
        if a[i] < b[i] {
            return false
        }
    }
    return len(a) > len(b)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}