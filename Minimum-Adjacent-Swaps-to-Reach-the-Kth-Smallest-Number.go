func getMinSwaps(num string, k int) int {
    originalNum := num

    digits := []rune(num)

    for i := 0; i < k; i++ {
        if !next(digits) {
            return -1
        }
    }

    targetNum := string(digits)

    return minAdjacentSwaps(originalNum, targetNum)
}

func minAdjacentSwaps(original, target string) int {
    orig := []byte(original)
    tar := []byte(target)

    n := len(orig)
    swaps := 0
    for i := 0; i < n; i++ {
        j := i
        for orig[j] != tar[i]{
            j++
        }

        swaps += j - i

        for j > i {
            orig[j], orig[j-1] = orig[j-1], orig[j]
            j--
        }
    }
    return swaps
}

func next(digits []rune) bool {
    n := len(digits)

    i := n - 2
    for i >= 0 && digits[i] >= digits[i + 1] {
        i--
    }

    if i < 0 {
        return false
    }

    j := n - 1
    for digits[j] <= digits[i] {
        j--
    }

    digits[i], digits[j] = digits[j], digits[i]

    left, right := i + 1, n - 1
    for left < right {
        digits[left], digits[right] = digits[right], digits[left]
        left++
        right--
    }

    return true
}