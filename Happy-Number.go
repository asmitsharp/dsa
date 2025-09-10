func isHappy(n int) bool {
    seen := make(map[int]bool)
    
    for n != 1 && !seen[n] {
        seen[n] = true
        n = squaresum(n)
    }
    return n == 1
}

func squaresum(num int) int {
    sum := 0
    for num > 0 {
        d := num % 10
        sum += d * d
        num /= 10
    }
    return sum
}

