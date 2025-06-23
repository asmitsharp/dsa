func removeDuplicateLetters(s string) string {
    lastO := make(map[rune]int)
    for i, char := range s {
        lastO[char] = i
    }

    var stack []rune
    seen := make(map[rune]bool)

    for i , char := range s {
        if !seen[char] {
            for len(stack) > 0 && char < stack[len(stack) - 1] && lastO[stack[len(stack) - 1]] > i {
                delete(seen, stack[len(stack) - 1])
                stack = stack[:len(stack) - 1]
            }

            stack = append(stack, char)
            seen[char] = true
        }
    }

    return string(stack)
}