func isValid(s string) bool {
    stack := make([]rune, 0)

    brackets := map[rune]rune{
        ')': '(',
		']': '[',
		'}': '{',
    }

    for _, char := range s {
        if char == '(' || char == '[' || char == '{' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }

            last := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if last != brackets[char] {
                return false
            }
        }
    }

    return len(stack) == 0
}