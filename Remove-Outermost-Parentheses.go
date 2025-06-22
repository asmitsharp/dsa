func removeOuterParentheses(s string) string {
    var result []rune
    count := 0

    for _, char := range s {

        if char == '(' {
            
            if count > 0 {
                result = append(result, char)
            }
            count++
        } else {
            count--
            if count > 0 {
                result = append(result, char)
            }
        }
    }

    return string(result)
}