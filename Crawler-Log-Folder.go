func minOperations(logs []string) int {
    stack := []string{}

    for _, op := range logs {
        if op == "../" {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        }  else if op == "./" {
            continue
        } else {
            stack = append(stack, op)
        }
    }

    return len(stack)
}