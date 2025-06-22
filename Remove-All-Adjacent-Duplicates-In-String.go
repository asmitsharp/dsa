func removeDuplicates(s string) string {
    stack := []byte{}
    i := 0
    for i < len(s) {
        if len(stack) > 0 && stack[len(stack) - 1] == s[i] {
            stack = stack[:len(stack) - 1] // pop
        } else {
            stack = append(stack, s[i])
        }
        i+=1
    }

    return string(stack)
}