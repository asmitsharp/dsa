import (
    "math"
	"strconv"
)

func nextGreaterElement(n int) int {
    // convert to digits
    s := strconv.Itoa(n)
    digits := []rune(s)

    len := len(digits)

    i := len - 2
    // find pivot
    for i >= 0 && digits[i] >= digits[i+1] {
        i--
    }

    if i < 0 {
        return -1
    }

    j := len - 1
    for digits[j] <= digits[i] {
        j--
    }

    // swap
    digits[i], digits[j] = digits[j], digits[i]

    left, right := i + 1, len - 1
    for left < right {
        digits[left], digits[right] = digits[right], digits[left]
        left++
        right--
    }

    resultStr := string(digits)
    result, err := strconv.ParseInt(resultStr, 10, 64)

    if err != nil || result > math.MaxInt32 {
		return -1
	}

    return int(result)
}