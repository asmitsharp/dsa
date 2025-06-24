func minLength(s string) int {
    var stack []rune

    for _,char := range s {
        if len(stack) > 0 && 
           ((char == 'B' && stack[len(stack)-1] == 'A') || 
            (char == 'D' && stack[len(stack)-1] == 'C')) {
            
                stack = stack[:len(stack) - 1]

        } else {
            stack = append(stack, char)
        }
    }

    return len(stack)
}