func calculate(s string) int {
    stack := []int{}
    num := 0
    op := '+'

    for i := 0; i < len(s); i++ {
        ch := s[i]

        if ch >= '0' && ch <= '9' {
            num = num*10 + int(ch-'0')
        }

        if (ch < '0' || ch > '9') && ch != ' ' || i == len(s)-1 {
            switch op {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                last := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, last*num)
            case '/':
                last := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, last/num)
            }

            op = rune(ch)
            num = 0
        }
    }

    result := 0
    for _, val := range stack {
        result += val
    }

    return result
}
