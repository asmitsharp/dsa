import "math"

func judgeSquareSum(c int) bool {
    a := 0
    b := int(math.Floor(math.Sqrt(float64(c))))

    for a <= b {
        sum := a * a + b * b
        if sum == c {
            return true
        } else if sum < c {
            a++
        } else {
            b--
        }
    }

    return false
}