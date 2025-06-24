func clearDigits(s string) string {
    var stack []rune

    for _, char := range s {
        if '0' <= char  && char <= '9' {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
        } else {
            stack = append(stack, char)
        }
    }

    return string(stack)
}